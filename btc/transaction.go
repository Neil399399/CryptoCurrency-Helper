package btc

import (
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
)

type BTCTx struct {
	client      *rpcclient.Client
	txFee       int64
	vaultClient Vault
}

func NewBtcClient() *BTCTx {
	connCfg := &rpcclient.ConnConfig{
		Host:         "epona:18332",
		User:         "bitcoinrpc",
		Pass:         "l5MgAmLQLrJyafDs2QnQX6sQYzJSFrBrC42d60H34YcH",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		panic("panic so sad")
	}
	return &BTCTx{
		client: client,
	}
}

func convertPublicToAddress(btcPublicKey string) (string, error) {
	addrPub, err := btcutil.NewAddressPubKey(base58.Decode(btcPublicKey), &chaincfg.TestNet3Params)
	if err != nil {
		return "", err
	}
	return addrPub.EncodeAddress(), nil
}

func (t *BTCTx) getOutputIndex(txid string, addr string) (int64, uint32, error) {
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

func getPayToAddrScript(address string) []byte {
	rcvAddress, _ := btcutil.DecodeAddress(address, &chaincfg.TestNet3Params)
	rcvScript, _ := txscript.PayToAddrScript(rcvAddress)
	return rcvScript
}

func (t *BTCTx) CreateTransaction(txid string, fromAddr string, toAddr string, amount int64) (*wire.MsgTx, Utxo, error) {
	unspentAmount, outputIndex, err := t.getOutputIndex(txid, fromAddr)
	if err != nil {
		return nil, Utxo{}, err
	}
	unspentTx := Utxo{
		Address:     fromAddr,
		TxID:        txid,
		OutputIndex: outputIndex,
		Script:      getPayToAddrScript(fromAddr),
		Satoshis:    amount, // the amount we want to send
	}

	hash, err := chainhash.NewHashFromStr(unspentTx.TxID)
	if err != nil {
		return nil, Utxo{}, err
	}

	// Creste raw tx
	redemTx := wire.NewMsgTx(wire.TxVersion)
	outPoint := wire.NewOutPoint(hash, unspentTx.OutputIndex)
	txIn := wire.NewTxIn(outPoint, nil, nil)
	redemTx.AddTxIn(txIn)

	// Create TxOut
	rcvScript := getPayToAddrScript(toAddr)
	outCoin := unspentTx.Satoshis
	txOut := wire.NewTxOut(outCoin, rcvScript)
	redemTx.AddTxOut(txOut)

	// If the above TxOut leads to change, let the change flow back to sneder
	change := unspentAmount - unspentTx.Satoshis - t.txFee
	if change > 0 {
		changeScript := getPayToAddrScript(fromAddr)
		changeTxOut := wire.NewTxOut(change, changeScript)
		redemTx.AddTxOut(changeTxOut)
	}
	return redemTx, unspentTx, nil
}

func (t *BTCTx) SignTransaction(tx *wire.MsgTx, unspentTx Utxo) (*wire.MsgTx, error) {
	// hash transaction
	hash, err := txscript.CalcSignatureHash(unspentTx.Script, txscript.SigHashAll, tx, 0)
	if err != nil {
		return nil, err
	}

	// call vault to sign tx
	respSig, respPK, err := t.vaultClient.Sign("aetheras_btc_3", "btc", "testnet", "1", base58.Encode(hash))
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
	tx.TxIn[0].SignatureScript = signature
	return tx, nil
}

// validate signature
func (t *BTCTx) ValidateTranscation(tx *wire.MsgTx, unspentTx Utxo) error {
	flags := txscript.StandardVerifyFlags
	vm, err := txscript.NewEngine(unspentTx.Script, tx, 0, flags, nil, nil, unspentTx.Satoshis) // Set to 0 because we only have one input
	if err != nil {
		return err
	}

	if err := vm.Execute(); err != nil {
		return err
	}
	return nil
}

// send transaction
func (t *BTCTx) SendTransaction(tx *wire.MsgTx) (*chainhash.Hash, error) {
	txhash, err := t.client.SendRawTransaction(tx, true)
	if err != nil {
		return nil, err
	}
	return txhash, nil
}

type VoutDetail2 struct {
	BlockHash     string
	Txid          string
	Address       string
	Category      string
	Amount        float64
	Vout          uint32
	Confirmations uint64
	Time          int64
	LockTime      uint32
	Blocktime     int64
}

func (t *BTCTx) ListenBitcoinChain(addr string, endBlockHeight int64) (*[]VoutDetail2, error) {
	result := []VoutDetail2{}
	currentBlockHeight, err := t.client.GetBlockCount()
	if err != nil {
		return nil, err
	}
	currentBlockHeight = 1611503
	fmt.Println("currentBlockHeight", currentBlockHeight)
	fmt.Println("endBlockHeight", endBlockHeight)

	// The genesis block coinbase is not considered an ordinary transaction
	// and cannot be retrieved. So we should stop at 1st block.
	if endBlockHeight < 1 {
		endBlockHeight = 1
	}

	if endBlockHeight > currentBlockHeight {
		return nil, errors.New("Block doesn't exist")
	}

	currBlockHash, err := t.client.GetBlockHash(currentBlockHeight)
	if err != nil {
		return nil, err
	}
	fmt.Println("currBlockHash", currBlockHash)

	currBlock, err := t.client.GetBlock(currBlockHash)
	if err != nil {
		return nil, err
	}
	fmt.Println("currBlock", currBlock)

	endBlockHash, err := t.client.GetBlockHash(endBlockHeight)
	if err != nil {
		return nil, err
	}
	fmt.Println("endBlockHash", endBlockHash)

	for {
		for _, tx := range currBlock.Transactions {
			txHash := tx.TxHash()
			txRawResult, err := t.client.GetRawTransactionVerbose(&txHash)
			if err != nil {
				return nil, err
			}

			for _, out := range txRawResult.Vout {
				if len(out.ScriptPubKey.Addresses) > 0 && addr == out.ScriptPubKey.Addresses[0] {
					voutDetail := VoutDetail2{
						BlockHash:     txRawResult.BlockHash,
						Txid:          txRawResult.Txid,
						Address:       addr,
						Category:      "receive",
						Amount:        out.Value,
						Vout:          out.N,
						Confirmations: txRawResult.Confirmations,
						Time:          txRawResult.Time,
						LockTime:      txRawResult.LockTime,
						Blocktime:     txRawResult.Blocktime,
					}
					result = append(result, voutDetail)
				}
			}
		}

		if currBlockHash.String() == endBlockHash.String() {
			break
		}

		currBlockHash = &currBlock.Header.PrevBlock
		currBlock, err = t.client.GetBlock(currBlockHash)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}
