package ethereum

import (
	"context"
	"fmt"
	"testing"

	"github.com/Neil399399/bitcoin-helper/vault"
)

const testAddressSender = "0x1Ed2001e00Da365b0b589f5f15507982235B30D5"
const testAddressTester = "0x27bbe78C9FE77A0959b0Cf219cfADFEdB311462e"

func TestTransaction(t *testing.T) {
	ethRepo := NewEthClient()
	vault := vault.NewVaultClient("http://localhost:8200", "root")
	ethRepo.vaultClient = *vault
	ctx := context.Background()

	tx, err := ethRepo.CreateTransaction(ctx, false, testAddressSender, testAddressTester, 3)
	if err != nil {
		t.Fatal(err)
	}

	signedTx, err := ethRepo.SignTransaction(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	txHash, err := ethRepo.SendTransaction(ctx, signedTx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("txHash", txHash)
}
