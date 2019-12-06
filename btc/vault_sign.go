package btc

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

type Vault struct {
	client *api.Client
}

func NewVaultClient(vaultURL, vaultToken string) *Vault {
	//create the client
	vaultConfig := api.DefaultConfig()
	vaultConfig.Address = vaultURL
	// new client
	Client, err := api.NewClient(vaultConfig)
	if err != nil {
		return nil
	}

	// set token
	Client.SetToken(vaultToken)
	return &Vault{
		client: Client,
	}
}

func (v *Vault) Sign(keyID, cointype, network, childIdx, message string) (string, string, error) {
	// vault path and request
	vaultSignPath := fmt.Sprintf("aetheras-plugin/signer/%s/%s", cointype, keyID)
	btcData := map[string]interface{}{
		"cointype": cointype,
		"keyID":    keyID,
		"network":  network,
		"childIdx": childIdx,
		"value":    message,
	}

	// response
	resp, err := v.client.Logical().Write(vaultSignPath, btcData)
	if err != nil {
		return "", "", err
	}
	respSig := fmt.Sprintf("%v", resp.Data["value"])
	respPk := fmt.Sprintf("%v", resp.Data["publicKey"])

	return respSig, respPk, nil
}
