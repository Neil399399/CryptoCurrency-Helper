package ethereum

import (
	"context"
	"fmt"
	"testing"

	"github.com/Neil399399/bitcoin-helper/vault"
	"github.com/ethereum/go-ethereum/core/types"
)

const testAddressSender = "0x66C479b428c44A50Ecb85e85d30138F08D454906"
const testAddressTester = "0xF0d65479732eedc406C00FFB29BC9dD426780eE4"
const testContractAddress = "0xd6Aa859b99A546Cf91f7ea39AEF04CDE5a0F6f23"

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

func TestTokenTransactionWithMemo(t *testing.T) {
	ethRepo := NewEthClient()
	vault := vault.NewVaultClient("http://localhost:8200", "root")
	ethRepo.vaultClient = *vault
	ctx := context.Background()
	ethRepo.contractAddr = testContractAddress

	tx, err := ethRepo.CreateTransactionWithMemo(ctx, true, testAddressSender, testAddressTester, 1000, "hello world") // 1 = 1ETH
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
