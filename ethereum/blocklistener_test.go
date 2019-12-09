package ethereum

import (
	"fmt"
	"testing"
)

const nowBlock = 5582755

var addressBook = []string{
	"0x1Ed2001e00Da365b0b589f5f15507982235B30D5",
	"0x27bbe78C9FE77A0959b0Cf219cfADFEdB311462e",
}

func TestBlockListener(t *testing.T) {
	ethRepo := NewEthClient()
	ethRepo.blockRange = 10
	txs, err := ethRepo.BlockListener(addressBook, nowBlock)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("address 1 Txs", txs[addressBook[0]])
	fmt.Println("address 2 Txs", txs[addressBook[1]])
}
