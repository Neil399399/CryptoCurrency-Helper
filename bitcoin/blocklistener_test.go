package bitcoin

import (
	"log"
	"testing"

	"github.com/Neil399399/bitcoin-helper/vault"
)

var testAddressBook = []string{"munZ5L7fE8Hmiqvc2f3ze9Wo1TjBXfACxR", "mvN8gFRPEwwt8XBpwr7gkFgDCPNMhtNyXA"}

const nowBlockNumber = 1344887

func TestBlockListener(t *testing.T) {
	client := NewBtcClient(HOST, LOGIN_ACC, LOGIN_PWD)
	vault := vault.NewVaultClient(vaultHost, vaultToken)
	btcTx := BtcTxClient(client.HttpClient, vault, txFee)

	events, err := btcTx.ListenBitcoinChain(testAddressBook, nowBlockNumber-5)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("Events", events)

	//scan
	log.Println("address 1 transactions", events[testAddressBook[0]])
	log.Println("address 2 transactions", events[testAddressBook[1]])
}
