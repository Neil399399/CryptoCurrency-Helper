package bitcoin

import (
	"github.com/Neil399399/bitcoin-helper/vault"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
)

type Client struct {
	httpClient  *rpcclient.Client
	vaultClient vault.Vault
}

// NewBtcClient new the rpc client
func NewBtcClient(host, user, password string) *Client {
	connCfg := &rpcclient.ConnConfig{
		Host:         host,
		User:         user,
		Pass:         password,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		panic("create new client failed, so sad")
	}
	return &Client{
		httpClient: client,
	}
}

func (c *Client) GetBlock(currentBlockHeight int64) (*wire.MsgBlock, error) {
	currBlockHash, err := c.httpClient.GetBlockHash(currentBlockHeight)
	if err != nil {
		return nil, err
	}

	currBlock, err := c.httpClient.GetBlock(currBlockHash)
	if err != nil {
		return nil, err
	}
	return currBlock, nil
}

func (c *Client) GetBalance(address string) (string, error) {
	accBalance, err := c.httpClient.GetBalance(address)
	if err != nil {
		return "", err
	}
	return accBalance.String(), nil
}

func (c *Client) GetTransaction(txhash string) (*btcjson.GetTransactionResult, error) {
	hash, err := chainhash.NewHashFromStr(txhash)
	if err != nil {
		return nil, err
	}
	transacation, err := c.httpClient.GetTransaction(hash)
	if err != nil {
		return nil, err
	}
	return transacation, nil
}
