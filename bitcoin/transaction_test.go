package bitcoin

import (
	"fmt"
	"testing"

	"github.com/Neil399399/bitcoin-helper/vault"
)

const toAddr = "munZ5L7fE8Hmiqvc2f3ze9Wo1TjBXfACxR"
const fromAddr = "mvN8gFRPEwwt8XBpwr7gkFgDCPNMhtNyXA"
const publicKey = "o1qS7m37GJp8JB19nbTgWtoWxr9X7QYMc8Zjs5jYVYsA"
const txFee = 10000

var txids = []string{
	"9bcdc1c03e4a723ccdf69ffb674325a966bb28ba6035f6bb4a817ea558e21357",
}

func TestSendTransaction(t *testing.T) {
	btcTx := NewBtcClient()
	vault := vault.NewVaultClient("http://localhost:8200", "root")

	btcTx.txFee = txFee
	btcTx.vaultClient = *vault

	msgTx, unspentTx, err := btcTx.CreateTransaction(txids, fromAddr, toAddr, int64(10000))
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
	fmt.Println("TXHASH", txhash)
}
