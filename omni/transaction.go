package omni

import (
	"math/big"

	"github.com/Neil399399/bitcoin-helper/bitcoin"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
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
}

func (c *Client) CreateOmniTransaction(txids []TxID, fromAddr string, toAddr string, propertyID, amount int64) (*wire.MsgTx, []byte, error) {
	var redemTx *wire.MsgTx
	redemTx = wire.NewMsgTx(wire.TxVersion)
	var unspentTx Utxo
	var total int64
	for _, txid := range txids {
		total = total + txid.Balance
		unspentTx = Utxo{
			Address:     fromAddr,
			TxID:        txid.TxID,
			OutputIndex: txid.OutputIndex,
			Script:      getPayToAddrScript(fromAddr),
		}

		hash, err := chainhash.NewHashFromStr(unspentTx.TxID)
		if err != nil {
			return nil, nil, err
		}

		// Create raw tx
		outPoint := wire.NewOutPoint(hash, unspentTx.OutputIndex)
		txIn := wire.NewTxIn(outPoint, nil, nil)
		redemTx.AddTxIn(txIn)
	}
	// Create TxOut
	rcvScript := getPayToAddrScript(toAddr)
	txOut := wire.NewTxOut(0, rcvScript)
	redemTx.AddTxOut(txOut)

	// Create tx opreturn (omni_createrawtx_opreturn)
	payload := createPayload(propertyID, amount)
	pkScript, err := txscript.NullDataScript(payload)
	if err != nil {
		return nil, nil, err
	}
	outputs := wire.NewTxOut(0, pkScript)
	redemTx.AddTxOut(outputs)

	// If the above TxOut leads to change, let the change flow back to sneder
	change := total - c.config.BitcoinNetFee
	if change > 0 {
		changeScript := getPayToAddrScript(fromAddr)
		changeTxOut := wire.NewTxOut(change, changeScript)
		redemTx.AddTxOut(changeTxOut)
	}
	return redemTx, unspentTx.Script, nil
}

func (c *Client) SignTransaction(redemTx *wire.MsgTx, scriptPubKey []byte, signInfo bitcoin.SignInfo) (*wire.MsgTx, error) {
	return c.btcTxClient.SignTransaction(redemTx, scriptPubKey, signInfo)
}

// ValidateTranscation check the transaction is valid
func (c *Client) ValidateTranscation(tx *wire.MsgTx, script []byte) error {
	return c.btcTxClient.ValidateTranscation(tx, script, 0)
}

// SendTransaction send the transaction by bitcoin client
func (c *Client) SendOmniTransaction(tx *wire.MsgTx) (*chainhash.Hash, error) {
	return c.btcTxClient.SendTransaction(tx)
}

func createPayload(propertyID, amount int64) []byte {
	propertyIDHex := intToHex(propertyID)
	amountHex := intToHex(amount)
	return common.Hex2Bytes(OMNI_OP_CODE + propertyIDHex + amountHex)
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
