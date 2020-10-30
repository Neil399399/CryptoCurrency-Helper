package bitcoin

import "testing"

const (
	HOST          = "35.234.25.4:30333"
	LOGIN_ACC     = "user"
	LOGIN_PWD     = "123456"
	testAddress_1 = "mnLPFNZbfNRJaFyZ71MEZ3L8xMaBaCP5di"
	testTxID      = "7b28aab5b6c3dc5ea7003568879392a18034141a52f855ba2d57568850e1fc99"
)

func TestGetBlock(t *testing.T) {
	newClient := NewBtcClient(HOST, LOGIN_ACC, LOGIN_PWD)
	blockInfo, err := newClient.GetBlock(nowBlockNumber)
	if err != nil {
		t.Fatal("Get Block Failed, so sad", err)
	}

	t.Log("BlockHash:", blockInfo.BlockHash().String())
	t.Log("PrevBlock:", blockInfo.Header.PrevBlock.String())
	t.Log("Timestamp:", blockInfo.Header.Timestamp)
	t.Log("Version:", blockInfo.Header.Version)
	t.Log("Nonce:", blockInfo.Header.Nonce)
}

func TestGetBlance(t *testing.T) {
	newClient := NewBtcClient(HOST, LOGIN_ACC, LOGIN_PWD)
	accBalance, err := newClient.GetBalance(testAddress_1)
	if err != nil {
		t.Fatal("Get Account Balance Failed", err)
	}
	t.Log("Account Balance:", accBalance)
}

func TestGetTransaction(t *testing.T) {
	newClient := NewBtcClient(HOST, LOGIN_ACC, LOGIN_PWD)
	txInfo, err := newClient.GetTransaction(testTxID)
	if err != nil {
		t.Fatal("Get Account Balance Failed", err)
	}

	t.Log("BlockIndex:", txInfo.BlockIndex)
	t.Log("BlockHash:", txInfo.BlockHash)
	t.Log("BlockTime:", txInfo.BlockTime)
	t.Log("TxID:", txInfo.TxID)
	t.Log("Confirmations:", txInfo.Confirmations)
	t.Log("Hex:", txInfo.Hex)
	t.Log("Time:", txInfo.Time)
	t.Log("TimeReceived:", txInfo.TimeReceived)
	t.Log("Amount:", txInfo.Amount)
	t.Log("Fee:", txInfo.Fee)
	for _, detail := range txInfo.Details {
		t.Log("[Detail] Account:", detail.Account)
		t.Log("[Detail] Address:", detail.Address)
		t.Log("[Detail] Amount:", detail.Amount)
		t.Log("[Detail] Category:", detail.Category)
		t.Log("[Detail] Fee:", detail.Fee)
		t.Log("[Detail] InvolvesWatchOnly:", detail.InvolvesWatchOnly)
		t.Log("[Detail] Vout:", detail.Vout)
	}
}
