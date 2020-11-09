package omni

import (
	"testing"

	"github.com/Neil399399/bitcoin-helper/bitcoin"
)

// config
const HOST = "35.234.25.4:30333"
const LOGIN_ACC = "user"
const LOGIN_PWD = "123456"
const testAddress_1 = "mnLPFNZbfNRJaFyZ71MEZ3L8xMaBaCP5di" // create by vault, childIdx=9000
const testAddress_2 = "mtwX2rrT7kQRKdeDWs9U8aR7qRg4oR71Ap" // create by vault, childIdx=9001
const omniToken_address = "moneyqMan7uh8FqdCA2BV5yZ8qVrc9ikLP"
const OmniTokenId = 1
const testOmniTokenId = 2 // test omni token property id
const testPayload = "00000000000000020000000000989680"

func TestGetInfo(t *testing.T) {
	config := &ConnConfig{
		Host: HOST,
		User: LOGIN_ACC,
		Pass: LOGIN_PWD,
	}
	newClient := New(config)
	resp, err := newClient.GetInfo()
	if err != nil {
		t.Fatal("Get omni info failed, so sad", err)
	}

	t.Log("VersionInt:", resp.VersionInt)
	t.Log("Version:", resp.Version)
	t.Log("BitcoinCoreVersion:", resp.BitcoinCoreVersion)
	t.Log("CommitInfo:", resp.CommitInfo)
	t.Log("Block:", resp.Block)
	t.Log("BlockTimestamp:", resp.BlockTimestamp)
	t.Log("BlockTransaction:", resp.BlockTransaction)
	t.Log("TotalTransaction:", resp.TotalTransaction)
}

func TestGetBalance(t *testing.T) {
	config := &ConnConfig{
		Host: HOST,
		User: LOGIN_ACC,
		Pass: LOGIN_PWD,
	}
	newClient := New(config)
	resp, err := newClient.GetBalance(testAddress_1, OmniTokenId)
	if err != nil {
		t.Fatal("Get balance failed, so sad", err)
	}
	t.Log("Balance:", resp.Balance)
	t.Log("Reserved:", resp.Reserved)
	t.Log("Frozen:", resp.Frozen)
}

func TestGetAllBalancesForAddress(t *testing.T) {
	config := &ConnConfig{
		Host: HOST,
		User: LOGIN_ACC,
		Pass: LOGIN_PWD,
	}
	newClient := New(config)
	resp, err := newClient.GetAllBalancesForAddress(testAddress_2)
	if err != nil {
		t.Fatal("Get balances by address failed, so sad", err)
	}

	for i := 0; i < len(resp); i++ {
		eachBalance := resp[i]
		t.Log("PropertyId:", eachBalance.PropertyId)
		t.Log("Name:", eachBalance.Name)
		t.Log("Balance:", eachBalance.Balance)
		t.Log("Reserved:", eachBalance.Reserved)
		t.Log("Frozen:", eachBalance.Frozen)
		t.Log("----------------------------------------")
	}
}

// func TestGetAddressTxids(t *testing.T) {
// 	config := &ConnConfig{
// 		Host: HOST,
// 		User: LOGIN_ACC,
// 		Pass: LOGIN_PWD,
// 	}
// 	newClient := New(config)
// 	resp, err := newClient.GetAddressTxids([]string{testAddress_1})
// 	if err != nil {
// 		t.Fatal("Get address txids failed, so sad", err)
// 	}
// 	t.Log("Txids:", resp)
// }

// func TestGetListOfUnspent(t *testing.T) {
// 	config := &ConnConfig{
// 		Host: HOST,
// 		User: LOGIN_ACC,
// 		Pass: LOGIN_PWD,
// 	}
// 	newClient := New(config)
// 	resp, err := newClient.ListUnSpent([]string{testAddress_1})
// 	if err != nil {
// 		t.Fatal("Get list of unspent failed, so sad", err)
// 	}

// 	for i := 0; i < len(resp.Utxos); i++ {
// 		each := resp.Utxos[i]
// 		t.Log("Address:", each.Address)
// 		t.Log("Tx:", each.Tx)
// 		t.Log("OutputIndex:", each.OutputIndex)
// 		t.Log("Script:", each.Script)
// 		t.Log("Satoshis:", each.Satoshis)
// 		t.Log("Height:", each.Height)
// 		t.Log("Vout:", each.Vout)
// 		t.Log("Coinbase:", each.Coinbase)
// 		t.Log("----------------------------------------")
// 	}
// 	t.Log("Hash:", resp.Hash)
// 	t.Log("Height:", resp.Height)
// }

func TestGetTransaction(t *testing.T) {
	config := &ConnConfig{
		Host: HOST,
		User: LOGIN_ACC,
		Pass: LOGIN_PWD,
	}
	newClient := New(config)
	const testTxID = "4c26bdaabfe737348f9b06ad21fc27abe4a92e34812ca5a080297353ad75968d"
	resp, err := newClient.GetTransaction(testTxID)
	if err != nil {
		t.Fatal("Get Transaction Failed, so sad", err.Error())
	}
	t.Log("----------------------------------------")
	t.Log("TxID:", resp.TxID)
	t.Log("SendingAddress:", resp.SendingAddress)
	t.Log("ReferenceAddress:", resp.ReferenceAddress)
	t.Log("Positioninblock:", resp.Positioninblock)
	t.Log("Block:", resp.Block)
	t.Log("BlockHash:", resp.BlockHash)
	t.Log("BlockTime:", resp.BlockTime)
	t.Log("Ismine:", resp.Ismine)
	t.Log("Vaild:", resp.Vaild)
	t.Log("Version:", resp.Version)
	t.Log("Amount:", resp.Amount)
	t.Log("Fee:", resp.Fee)
	t.Log("Type:", resp.Type)
	t.Log("TypeInt:", resp.TypeInt)
	t.Log("Confirmations:", resp.Confirmations)
	t.Log("----------------------------------------")
}

func TestCreateOpReturn(t *testing.T) {
	config := &ConnConfig{
		Host: HOST,
		User: LOGIN_ACC,
		Pass: LOGIN_PWD,
	}
	newClient := New(config)
	respB, err := newClient.CreateOpReturn("", testPayload)
	if err != nil {
		t.Fatal("Create OpReturn Failed, so sad", err.Error())
	}
	t.Log("OP_RETURN:", respB)
}

func TestSendTransaction(t *testing.T) {
	config := &ConnConfig{
		Host:          HOST,
		User:          LOGIN_ACC,
		Pass:          LOGIN_PWD,
		BitcoinNetFee: 10000,
	}
	newClient := New(config)
	amount := int64(20000000)
	SignKey := bitcoin.SignInfo{
		CoinType: "btc",
		Network:  "testnet",
		KeyID:    "aetheras_btc_4",
		ChildIdx: "9001",
	}

	txIds := []TxID{
		{
			TxID:        "ff3a48d43c95794f70510c99e35acf9618d81a024d58145d9acd761eea9a764a",
			OutputIndex: 1,
			Balance:     5000000,
		},
	}

	// create transaction
	unSignTx, publicKeyScript, err := newClient.CreateOmniTransaction(txIds, testAddress_2, testAddress_1, testOmniTokenId, amount)
	if err != nil {
		t.Fatal("Create Transaction Failed, so sad", err)
	}
	t.Log("unSignTx", unSignTx)

	// sign transaction
	signedTx, err := newClient.SignTransaction(unSignTx, publicKeyScript, SignKey)
	if err != nil {
		t.Fatal("Sign Transaction Failed, so sad", err)
	}
	t.Log("signed hash", signedTx.TxHash().String())

	// validate transaction
	err = newClient.ValidateTranscation(signedTx, publicKeyScript)
	if err != nil {
		t.Fatal("Validate Transaction Failed, so sad", err)
	}

	// send transacrion
	rawTxHash, err := newClient.SendOmniTransaction(signedTx)
	if err != nil {
		t.Fatal("Send Transaction Failed, so sad", err)
	}
	t.Log("rawTxHash", rawTxHash)
}
