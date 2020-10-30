package bitcoin

import (
	"fmt"
	"testing"
)

var testAddressBook = []string{"munZ5L7fE8Hmiqvc2f3ze9Wo1TjBXfACxR", "mvN8gFRPEwwt8XBpwr7gkFgDCPNMhtNyXA"}

const nowBlockNumber = 1344887

func TestBlockListener(t *testing.T) {
	btcTx := NewBtcClient()
	events, err := btcTx.ListenBitcoinChain(testAddressBook, nowBlockNumber-5)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Events", events)

	//scan
	fmt.Println("address 1 transactions", events[testAddressBook[0]])
	fmt.Println("address 2 transactions", events[testAddressBook[1]])
}

func TestGetBlock(t *testing.T) {
	btcTx := NewBtcClient()
	btcTx.GetBlock(nowBlockNumber)
}
