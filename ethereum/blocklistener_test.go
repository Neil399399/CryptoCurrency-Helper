package ethereum

import (
	"context"
	"fmt"
	"testing"
)

const nowBlock = 6185540

// target address
var addressBook = []string{
	"0xF0d65479732eedc406C00FFB29BC9dD426780eE4",
}

// listen target
func TestBlockListener(t *testing.T) {
	ethRepo := NewEthClient()
	ethRepo.blockRange = 1
	txs, err := ethRepo.BlockListener(addressBook, nowBlock)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("address 1 Txs", txs[addressBook[0]])
}

func TestContractListener(t *testing.T) {
	ethRepo := NewEthClient()
	ethRepo.blockRange = 10
	txs, err := ethRepo.ContractListenr(testContractAddress, addressBook, nowBlock)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("address 1 Txs", txs[addressBook[0]])
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
