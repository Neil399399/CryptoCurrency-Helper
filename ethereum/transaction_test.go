package ethereum

import (
	"context"
	"fmt"
	"testing"

	"github.com/Neil399399/bitcoin-helper/vault"
	"github.com/ethereum/go-ethereum/core/types"
)

const testAddressSender = "0x97c71401adD1eb50659820df1148857504228565"
const testAddressTester = "0x7b1feCA9E2d3a43eeC4c59Af5576d831394C8f69"
const testContractAddress = "0xd6aa859b99a546cf91f7ea39aef04cde5a0f6f23"

func TestTransaction(t *testing.T) {
	ethRepo := NewEthClient()
	vault := vault.NewVaultClient("http://localhost:8200", "root")
	ethRepo.vaultClient = *vault
	ctx := context.Background()
	ethRepo.contractAddr = testContractAddress

	tx, err := ethRepo.CreateTransaction(ctx, false, testAddressSender, testAddressTester, 0)
	if err != nil {
		t.Fatal(err)
	}

	signedTx, err := ethRepo.SignTransaction(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	B, err := signedTx.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	var test2 types.Transaction
	test2.UnmarshalJSON(B)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(test2)

	txHash, err := ethRepo.SendTransaction(ctx, &test2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("txHash", txHash)
}
