package ethereum

import (
	"context"
	"fmt"
	"testing"

	"github.com/Neil399399/bitcoin-helper/vault"
	"github.com/ethereum/go-ethereum/core/types"
)

const testAddressSender = "0x791e05d274d6e0da6CE5E6433A7B9310765894E3"
const testAddressTester = "0xF0d65479732eedc406C00FFB29BC9dD426780eE4"
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

	signedTx, err := ethRepo.SignTransaction(ctx, tx, "90")
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

func TestTransactionWithMemo(t *testing.T) {
	ethRepo := NewEthClient()
	vault := vault.NewVaultClient("http://localhost:8200", "root")
	ethRepo.vaultClient = *vault
	ctx := context.Background()
	ethRepo.contractAddr = testContractAddress

	tx, err := ethRepo.CreateTransactionWithMemo(ctx, false, testAddressSender, testAddressTester, 1, "hello world") // 1 = 1ETH
	if err != nil {
		t.Fatal(err)
	}

	signedTx, err := ethRepo.SignTransaction(ctx, tx, "90")
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
