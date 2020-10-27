package omni

import (
	"testing"
)

const testAddress = "mnLPFNZbfNRJaFyZ71MEZ3L8xMaBaCP5di"
const OmniTokenId = 1     // omni token property id
const TestOmniTokenId = 2 // test omni token property id

func TestNewClient(t *testing.T) {
	config := &ConnConfig{
		Host: "localhost:18332",
		User: "user",
		Pass: "123456",
	}
	newClient := New(config)
	resp, err := newClient.GetInfo()
	if err != nil {
		t.Error("get omni info failed", err)
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
		Host: "localhost:18332",
		User: "user",
		Pass: "123456",
	}
	newClient := New(config)
	resp, err := newClient.GetBalance(testAddress, OmniTokenId)
	if err != nil {
		t.Error("get balance failed", err)
	}
	t.Log("Balance:", resp.Balance)
	t.Log("Reserved:", resp.Reserved)
	t.Log("Frozen:", resp.Frozen)
}

func TestGetAllBalancesForAddress(t *testing.T) {
	config := &ConnConfig{
		Host: "localhost:18332",
		User: "user",
		Pass: "123456",
	}
	newClient := New(config)
	resp, err := newClient.GetAllBalancesForAddress("moneyqMan7uh8FqdCA2BV5yZ8qVrc9ikLP")
	if err != nil {
		t.Error("get balances by address failed", err)
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
