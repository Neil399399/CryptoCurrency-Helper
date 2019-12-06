package btc

import (
	"fmt"
	"testing"

	_ "github.com/btcsuite/btcwallet/walletdb/bdb"
)

const toAddr = "munZ5L7fE8Hmiqvc2f3ze9Wo1TjBXfACxR"
const fromAddr = "mvN8gFRPEwwt8XBpwr7gkFgDCPNMhtNyXA"
const publicKey = "o1qS7m37GJp8JB19nbTgWtoWxr9X7QYMc8Zjs5jYVYsA"
const txid = "c66e89bbab9675622e1156dbe98d8b9a3718f9b6b9bbeef6a04ce96dad1d1afd"
const txFee = 10000

func TestGetBalance(t *testing.T) {
	// publicKeyB := base58.Decode(publicKey)

	// dbPath := filepath.Join(os.TempDir(), "example.db")
	// fmt.Println("dbPath", dbPath)
	// db, err := walletdb.Create("bdb", dbPath)
	// if err != nil {
	// 	fmt.Println("Create")
	// 	panic(err)
	// }
	// defer os.Remove(dbPath)
	// defer db.Close()

	// w, err := wallet.Open(db, publicKeyB, nil, &chaincfg.TestNet3Params, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(w.ChainParams())
	btcTx := NewBtcClient()
	resp, err := btcTx.ListenBitcoinChain(toAddr, 1611501)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func TestSendTransaction(t *testing.T) {
	btcTx := NewBtcClient()
	vault := NewVaultClient("http://localhost:8200", "root")

	btcTx.txFee = txFee
	btcTx.vaultClient = *vault

	msgTx, unspentTx, err := btcTx.CreateTransaction(txid, fromAddr, toAddr, int64(10000))
	if err != nil {
		panic(err)
	}

	signedTx, err := btcTx.SignTransaction(msgTx, unspentTx)
	if err != nil {
		panic(err)
	}

	err = btcTx.ValidateTranscation(signedTx, unspentTx)
	if err != nil {
		panic(err)
	}

	// txhash, err := btcTx.SendTransaction(signedTx)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("TXHASH", txhash)
}
