package vault

import (
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
)

const btcMasterKeyRef = "aetheras_btc_4"
const ethMasterKeyRef = "aetheras_eth_4"

func TestGeneratePublicKey(t *testing.T) {
	var publicKeys = map[string]string{}
	v := NewVaultClient("http://localhost:8200", "root")

	// vault path
	vaultBtcPath := fmt.Sprintf("aetheras-plugin/generate/btc/%s", btcMasterKeyRef)
	vaultEthPath := fmt.Sprintf("aetheras-plugin/generate/eth/%s", ethMasterKeyRef)

	btcData := map[string]interface{}{
		"cointype": "btc",
		"network":  "testnet",
		"keyID":    btcMasterKeyRef,
		"childIdx": "90",
	}

	ethData := map[string]interface{}{
		"cointype": "eth",
		"network":  "",
		"keyID":    ethMasterKeyRef,
		"childIdx": "90",
	}

	// generate address
	btcResp, err := v.client.Logical().Write(vaultBtcPath, btcData)
	if err != nil || btcResp == nil {
		log.Error().Err(err).Msg("error to create btc address")
	}

	ethResp, err := v.client.Logical().Write(vaultEthPath, ethData)
	if err != nil || ethResp == nil {
		log.Error().Err(err).Msg("error to create btc address")
	}
	btcPub := fmt.Sprintf("%v", btcResp.Data["publicKey"])
	ethPub := fmt.Sprintf("%v", ethResp.Data["publicKey"])

	publicKeys["btc"] = btcPub
	publicKeys["eth"] = ethPub
	btcAddr, ethAddr, err := convertPublicToAddress(publicKeys["btc"], publicKeys["eth"], "testnet")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Bitcoin address", btcAddr)
	fmt.Println("Ethereum address", ethAddr)
}

func convertPublicToAddress(btcPublicKey, ethPublicKey, btcNetwork string) (string, string, error) {
	var addrPub *btcutil.AddressPubKey
	var err error
	if btcNetwork == "testnet" {
		addrPub, err = btcutil.NewAddressPubKey(base58.Decode(btcPublicKey), &chaincfg.TestNet3Params)
		if err != nil {
			return "", "", err
		}
	} else {
		addrPub, err = btcutil.NewAddressPubKey(base58.Decode(btcPublicKey), &chaincfg.MainNetParams)
		if err != nil {
			return "", "", err
		}
	}

	ecdsaPub, err := crypto.DecompressPubkey(base58.Decode(ethPublicKey))
	if err != nil {
		return "", "", err
	}
	return addrPub.EncodeAddress(), crypto.PubkeyToAddress(*ecdsaPub).Hex(), nil
}

func TestGetEmailPWD(t *testing.T) {
	v := NewVaultClient("http://localhost:8200", "root")
	resp, err := v.client.Logical().Read("secret/data/nex-sms")
	if err != nil {
		fmt.Println(err)
	}

	// respData := resp.Data["data"]

	m, ok := resp.Data["data"].(map[string]interface{})
	if !ok {
		return
	}

	var account, pwd string
	for k, v := range m {
		account = fmt.Sprintf("%v", k)
		pwd = fmt.Sprintf("%v", v)
	}
	fmt.Println(account, pwd)

	value, ok := m["sssaas"]
	if !ok {
		fmt.Println("NOT FOUND")
	}
	fmt.Println(value)
}
