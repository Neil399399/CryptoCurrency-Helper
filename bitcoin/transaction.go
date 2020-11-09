package bitcoin

import (
	"github.com/Neil399399/bitcoin-helper/vault"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
)

const OMNI_OP_CODE = "6f6d6e69"

// BtcTx is bttcoin transaction repo
type BtcTx struct {
	client      *rpcclient.Client
	txFee       int64
	vaultClient *vault.Vault
}

func BtcTxClient(btcClient *rpcclient.Client, vaultClient *vault.Vault, txFee int64) *BtcTx {
	return &BtcTx{
		client:      btcClient,
		vaultClient: vaultClient,
		txFee:       txFee,
	}
}

func (t *BtcTx) getOutputIndex(txid string, addr string) (int64, uint32, error) {
	txHash, err := chainhash.NewHashFromStr(txid)
	if err != nil {
		return 0, 0, err
	}

	txRawResult, err := t.client.GetRawTransactionVerbose(txHash)
	if err != nil {
		return 0, 0, err
	}

	for idx, output := range txRawResult.Vout {
		if addr == output.ScriptPubKey.Addresses[0] {
			amount := int64(output.Value * 100000000)
			return amount, uint32(idx), nil
		}
	}
	return 0, 0, nil
}

// CreateTransaction create a new bitcoin transaction (testnet)
func (t *BtcTx) CreateTransaction(txids []string, fromAddr string, toAddr string, amount int64) (*wire.MsgTx, []byte, error) {
	var redemTx *wire.MsgTx
	redemTx = wire.NewMsgTx(wire.TxVersion)
	var unspentTx Utxo
	var total int64
	for _, txid := range txids {
		unspentAmount, outputIndex, err := t.getOutputIndex(txid, fromAddr)
		if err != nil {
			return nil, nil, err
		}
		total = total + unspentAmount
		unspentTx = Utxo{
			Address:     fromAddr,
			TxID:        txid,
			OutputIndex: outputIndex,
			Script:      getPayToAddrScript(fromAddr),
		}

		hash, err := chainhash.NewHashFromStr(unspentTx.TxID)
		if err != nil {
			return nil, nil, err
		}

		// Creste raw tx
		outPoint := wire.NewOutPoint(hash, unspentTx.OutputIndex)
		txIn := wire.NewTxIn(outPoint, nil, nil)
		redemTx.AddTxIn(txIn)
	}
	// Create TxOut
	rcvScript := getPayToAddrScript(toAddr)
	outCoin := amount
	txOut := wire.NewTxOut(outCoin, rcvScript)
	redemTx.AddTxOut(txOut)

	// If the above TxOut leads to change, let the change flow back to sneder
	change := total - amount - t.txFee
	if change > 0 {
		changeScript := getPayToAddrScript(fromAddr)
		changeTxOut := wire.NewTxOut(change, changeScript)
		redemTx.AddTxOut(changeTxOut)
	}

	return redemTx, nil, nil
}

// CreateTransaction create a new bitcoin transaction (testnet)
func (t *BtcTx) CreateTransactionWithMemo(txids []string, fromAddr string, toAddr string, amount int64, memo string) (*wire.MsgTx, []byte, error) {
	var redemTx *wire.MsgTx
	redemTx = wire.NewMsgTx(wire.TxVersion)
	var unspentTx Utxo
	var total int64
	for _, txid := range txids {
		unspentAmount, outputIndex, err := t.getOutputIndex(txid, fromAddr)
		if err != nil {
			return nil, nil, err
		}
		total = total + unspentAmount
		unspentTx = Utxo{
			Address:     fromAddr,
			TxID:        txid,
			OutputIndex: outputIndex,
			Script:      getPayToAddrScript(fromAddr),
		}

		hash, err := chainhash.NewHashFromStr(unspentTx.TxID)
		if err != nil {
			return nil, nil, err
		}

		// Creste raw tx
		outPoint := wire.NewOutPoint(hash, unspentTx.OutputIndex)
		txIn := wire.NewTxIn(outPoint, nil, nil)
		redemTx.AddTxIn(txIn)
	}
	// Create TxOut
	rcvScript := getPayToAddrScript(toAddr)
	outCoin := amount
	txOut := wire.NewTxOut(outCoin, rcvScript)
	redemTx.AddTxOut(txOut)

	// If the above TxOut leads to change, let the change flow back to sneder
	change := total - amount - t.txFee
	if change > 0 {
		changeScript := getPayToAddrScript(fromAddr)
		changeTxOut := wire.NewTxOut(change, changeScript)
		redemTx.AddTxOut(changeTxOut)
	}

	// add comment
	pkScript, _ := txscript.NullDataScript([]byte(memo))
	outputs := wire.NewTxOut(int64(0), pkScript)
	redemTx.AddTxOut(outputs)

	return redemTx, unspentTx.Script, nil
}

// SignTransaction sign the transdaction with vault
func (t *BtcTx) SignTransaction(tx *wire.MsgTx, script []byte, signInfo SignInfo) (*wire.MsgTx, error) {
	for i := 0; i < len(tx.TxIn); i++ {

		// hash transaction
		hash, err := txscript.CalcSignatureHash(script, txscript.SigHashAll, tx, i)
		if err != nil {
			return nil, err
		}
		// call vault to sign tx
		respSig, respPK, err := t.vaultClient.Sign(signInfo.KeyID, signInfo.CoinType, signInfo.Network, signInfo.ChildIdx, base58.Encode(hash))
		if err != nil {
			return nil, err
		}

		// getresponse and convert to bytes
		sigB := base58.Decode(respSig)
		pkB := base58.Decode(respPK)

		// add sign hash
		sigB = append(sigB, byte(txscript.SigHashAll))

		// combine sigtx and publickey to signature
		signature, err := txscript.NewScriptBuilder().AddData(sigB).AddData(pkB).Script()
		if err != nil {
			return nil, err
		}
		tx.TxIn[i].SignatureScript = signature
	}
	return tx, nil
}

// ValidateTranscation check the transaction is valid
func (t *BtcTx) ValidateTranscation(tx *wire.MsgTx, Script []byte, amount int64) error {
	flags := txscript.StandardVerifyFlags
	vm, err := txscript.NewEngine(Script, tx, 0, flags, nil, nil, amount) // Set to 0 because we only have one input
	if err != nil {
		return err
	}

	if err := vm.Execute(); err != nil {
		return err
	}
	return nil
}

// SendTransaction send the transaction to bitcoin chain
func (t *BtcTx) SendTransaction(tx *wire.MsgTx) (*chainhash.Hash, error) {
	txhash, err := t.client.SendRawTransaction(tx, true)
	if err != nil {
		return nil, err
	}
	return txhash, nil
}

func getPayToAddrScript(address string) []byte {
	rcvAddress, _ := btcutil.DecodeAddress(address, &chaincfg.TestNet3Params)
	rcvScript, _ := txscript.PayToAddrScript(rcvAddress)
	return rcvScript
}
