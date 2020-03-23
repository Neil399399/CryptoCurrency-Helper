package bitcoin

import (
	"fmt"
	"testing"

	"github.com/Neil399399/bitcoin-helper/vault"
)

const toAddr = "munZ5L7fE8Hmiqvc2f3ze9Wo1TjBXfACxR"
const fromAddr = "mvN8gFRPEwwt8XBpwr7gkFgDCPNMhtNyXA"
const txFee = 10000

var txids = []string{
	"0e5fc224f87e90e03d011c4f9ef74593d99315dc4a8c7ef65294160a796aff7b",
	"78e0a2a0b37d606968b229f46293cfbcda13526c71713c80f2682baf14b10193",
	"3dc50ea185d8f7f8e857ec0a3811b815c6974c49baadefd2303073db8161ed0a",
	"7d6f5014a38ff1e7e889430cc2ed87be4335a0a1089c19ffbc5e96f1e68e3bda",
}

func TestSendTransaction(t *testing.T) {
	btcTx := NewBtcClient()
	vault := vault.NewVaultClient("http://localhost:8200", "root")

	btcTx.txFee = txFee
	btcTx.vaultClient = *vault

	msgTx, unspentTx, err := btcTx.CreateTransaction(txids, fromAddr, toAddr, int64(239700000))
	if err != nil {
		t.Fatal(err)
	}

	signedTx, err := btcTx.SignTransaction(msgTx, unspentTx)
	if err != nil {
		t.Fatal(err)
	}

	err = btcTx.ValidateTranscation(signedTx, unspentTx)
	if err != nil {
		t.Fatal(err)
	}

	txhash, err := btcTx.SendTransaction(signedTx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Transaction hash:", txhash.String())
}
