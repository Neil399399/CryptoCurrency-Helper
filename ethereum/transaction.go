package ethereum

import (
	"context"
	"math/big"

	"github.com/Neil399399/bitcoin-helper/vault"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

type EthTx struct {
	client       *ethclient.Client
	blockRange   int64
	gasLimit     int64
	contractAddr string
	vaultClient  vault.Vault
}

func NewEthClient() *EthTx {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/022fa52e0c6f4c50977e2ff86faaa7cb")
	if err != nil {
		panic(err)
	}
	return &EthTx{
		client:     client,
		blockRange: 15,
		gasLimit:   6200000,
	}
}

func (e *EthTx) CreateTransaction(ctx context.Context, ERC20 bool, fromAddress, toAddress string, value int64) (*types.Transaction, error) {
	var newTxn *types.Transaction

	nonce, err := e.client.PendingNonceAt(ctx, common.HexToAddress(fromAddress))
	if err != nil {
		return nil, err
	}

	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	// set target address.
	targetAddr := common.HexToAddress(toAddress)
	valueBN := big.NewInt(value)
	wei := big.NewInt(10)
	demon := wei.Exp(wei, big.NewInt(18), nil)
	valueBN.Mul(valueBN, demon)

	gasLimit := e.gasLimit
	fee := big.NewInt(0)
	fee.Mul(gasPrice, big.NewInt(gasLimit))

	if ERC20 {
		amount := big.NewInt(value)
		tokenAddress := common.HexToAddress(e.contractAddr)

		transferFnSignature := []byte("transfer(address,uint256)")
		hash := sha3.NewLegacyKeccak256()
		hash.Write(transferFnSignature)
		methodID := hash.Sum(nil)[:4]
		paddedAddress := common.LeftPadBytes(targetAddr.Bytes(), 32)
		paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
		// save to byte array.
		var data []byte
		data = append(data, methodID...)
		data = append(data, paddedAddress...)
		data = append(data, paddedAmount...)

		newTxn = types.NewTransaction(nonce, tokenAddress, big.NewInt(0), uint64(gasLimit), gasPrice, data)
	} else {
		newTxn = types.NewTransaction(72, targetAddr, valueBN, uint64(gasLimit), gasPrice, nil)
	}
	return newTxn, nil
}

func (e *EthTx) SignTransaction(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	chainID, err := e.client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	hash := types.NewEIP155Signer(chainID).Hash(tx)
	respSig, _, err := e.vaultClient.Sign("aetheras_eth_4", "eth", "", "70", base58.Encode(hash.Bytes()))
	if err != nil {
		return nil, err
	}
	// getresponse and convert to bytes
	sigB := base58.Decode(respSig)
	return tx.WithSignature(types.NewEIP155Signer(chainID), sigB)
}

func (e *EthTx) SendTransaction(ctx context.Context, signedTx *types.Transaction) (string, error) {
	err := e.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", err
	}
	transactionHash := signedTx.Hash()
	return transactionHash.Hex(), nil
}
