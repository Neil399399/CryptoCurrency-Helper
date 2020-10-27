package omni

import (
	"testing"
)

const testAddress = "mnLPFNZbfNRJaFyZ71MEZ3L8xMaBaCP5di"
const OmniTokenId = 1 // omni token property id

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
		t.Error("get omni info failed", err)
	}
	t.Log("Balance:", resp.Balance)
	t.Log("Reserved:", resp.Reserved)
	t.Log("Frozen:", resp.Frozen)
}
