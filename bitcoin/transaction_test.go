package bitcoin

import (
	"testing"

	"github.com/Neil399399/bitcoin-helper/vault"
)

const (
	toAddr     = "mtwX2rrT7kQRKdeDWs9U8aR7qRg4oR71Ap"
	fromAddr   = "mnLPFNZbfNRJaFyZ71MEZ3L8xMaBaCP5di"
	vaultHost  = "http://localhost:8200"
	vaultToken = "root"
	txFee      = 10000
)

var txids = []string{
	"99ee131629ce75749d557238e08f2e709fac58b7a119adcff234e58d47b9a4f6",
	// "78e0a2a0b37d606968b229f46293cfbcda13526c71713c80f2682baf14b10193",
	// "3dc50ea185d8f7f8e857ec0a3811b815c6974c49baadefd2303073db8161ed0a",
	// "7d6f5014a38ff1e7e889430cc2ed87be4335a0a1089c19ffbc5e96f1e68e3bda",
}

func TestSendTransaction(t *testing.T) {
	client := NewBtcClient(HOST, LOGIN_ACC, LOGIN_PWD)
	vault := vault.NewVaultClient(vaultHost, vaultToken)
	btcTx := BtcTxClient(client.HttpClient, vault, txFee)
	const sendAmount = 50000 // satoshis
	SignKey := SignInfo{
		CoinType: "btc",
		Network:  "testnet",
		KeyID:    "aetheras_btc_4",
		ChildIdx: "9001",
	}

	msgTx, script, err := btcTx.CreateTransaction(txids, fromAddr, toAddr, sendAmount)
	if err != nil {
		t.Fatal(err)
	}

	signedTx, err := btcTx.SignTransaction(msgTx, script, SignKey)
	if err != nil {
		t.Fatal(err)
	}

	err = btcTx.ValidateTranscation(signedTx, script, sendAmount)
	if err != nil {
		t.Fatal(err)
	}

	txhash, err := btcTx.SendTransaction(signedTx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Transaction hash:", txhash.String())
}
