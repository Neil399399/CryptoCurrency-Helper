package omni

import (
	"math/big"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
)

const (
	OMNI_OP_CODE   = "6f6d6e69"
	VAULT_KEYID    = "aetheras_btc_4"
	VAULT_COINTYPE = "btc"
	VAULT_NETWORK  = "testnet"
	VAULT_CHILDIDX = "9000"
)

type Utxo struct {
	Address     string
	TxID        string
	OutputIndex uint32
	Script      []byte
	Satoshis    int64
}

func (c *Client) CreateTransaction(from, to string, propertyID int64, amount int64, fee int64, feeAddress string) (*wire.MsgTx, []byte, error) {
	// get address unspent txid
	unspentTx, err := c.ListUnSpent([]string{from})
	if err != nil {
		return nil, nil, err
	}
	// create tx input (omni_createrawtx_input)
	var redemTx *wire.MsgTx
	redemTx = wire.NewMsgTx(wire.TxVersion)
	for _, utxo := range unspentTx.Utxos {
		hash, err := chainhash.NewHashFromStr(utxo.Tx)
		if err != nil {
			return nil, nil, err
		}
		outPoint := wire.NewOutPoint(hash, uint32(utxo.Vout))
		txIn := wire.NewTxIn(outPoint, nil, nil)
		redemTx.AddTxIn(txIn)
	}

	// create tx opreturn (omni_createrawtx_opreturn)
	payload := createPayload(propertyID, amount)
	pkScript, err := txscript.NullDataScript(payload)
	if err != nil {
		return nil, nil, err
	}
	outputs := wire.NewTxOut(int64(0), pkScript)
	redemTx.AddTxOut(outputs)

	// add reference (omni_createrawtx_reference)
	rcvScript := getPayToAddrScript(to)
	txOut := wire.NewTxOut(0, rcvScript)
	redemTx.AddTxOut(txOut)

	// create tx change (omni_createrawtx_change)
	changeScript := getPayToAddrScript(feeAddress)
	changeTxOut := wire.NewTxOut(fee, changeScript)
	redemTx.AddTxOut(changeTxOut)

	return redemTx, getPayToAddrScript(from), nil
}

func (c *Client) SignTransaction(redemTx *wire.MsgTx, scriptPubKey []byte) (*wire.MsgTx, error) {
	// SignTransaction sign the transdaction with vault
	for index, txIn := range redemTx.TxIn {
		// hash transaction
		hash, err := txscript.CalcSignatureHash(scriptPubKey, txscript.SigHashAll, redemTx, index)
		if err != nil {
			return nil, err
		}
		// call vault to sign tx
		respSig, respPK, err := c.vaultClient.Sign(VAULT_KEYID, VAULT_COINTYPE, VAULT_NETWORK, VAULT_CHILDIDX, base58.Encode(hash))
		if err != nil {
			return nil, err
		}

		// get response and convert to bytes
		sigB := base58.Decode(respSig)
		pkB := base58.Decode(respPK)

		// add sign hash
		sigB = append(sigB, byte(txscript.SigHashAll))

		// combine sigtx and publickey to signature
		signature, err := txscript.NewScriptBuilder().AddData(sigB).AddData(pkB).Script()
		if err != nil {
			return nil, err
		}
		txIn.SignatureScript = signature
	}
	return redemTx, nil
}

func (c *Client) SendRawTransaction(from, to, tx string) (string, error) {
	hash, err := c.SendRawTx(from, to, tx, from, "")
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func createPayload(propertyID, amount int64) []byte {
	propertyIDHex := intToHex(propertyID)
	amountHex := intToHex(amount)
	return []byte(OMNI_OP_CODE + propertyIDHex + amountHex)
}

func getPayToAddrScript(address string) []byte {
	rcvAddress, _ := btcutil.DecodeAddress(address, &chaincfg.TestNet3Params)
	rcvScript, _ := txscript.PayToAddrScript(rcvAddress)
	return rcvScript
}

func intToHex(value int64) string {
	realValueB := common.LeftPadBytes(big.NewInt(value).Bytes(), 8)
	return common.Bytes2Hex(realValueB)
}
