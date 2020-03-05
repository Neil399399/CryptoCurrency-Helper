package ethereum

import (
	"context"
	"fmt"
	"testing"
)

const nowBlock = 5680385

var addressBook = []string{
	"0x2035a145Fa186C408B0aF174E31F2D4C27054219",
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

func TestContractListener(t *testing.T) {
	ethRepo := NewEthClient()
	ethRepo.blockRange = 10
	_, err := ethRepo.GetContractRecord(testContractAddress, addressBook, nowBlock)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetGasPrice(t *testing.T) {
	ethRepo := NewEthClient()
	ctx := context.Background()
	gasprice, err := ethRepo.client.SuggestGasPrice(ctx)
	fmt.Println(gasprice)
	if err != nil {
		t.Fatal(err)
	}
}
