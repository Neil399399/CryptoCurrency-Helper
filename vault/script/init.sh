#!/bin/bash
# This script is used to cleanly provision a new vault server
export VAULT_ADDR=http://127.0.0.1:8200
export PLUGIN_SHASUM=fbdbeb8c3698bec46ca9dcf1d707ed3633ca24e97686112f1647f5359872875a

set -e
echo "Vault Network: ${VAULT_ADDR}"
echo "Make Sure the Vault CLI is logged in with the proper permissions"

vault audit enable file file_path=stdout
vault secrets enable -version=1 -path=secrets/db kv
vault secrets enable -version=1 -path=secrets/services kv
vault secrets enable -version=1 -path=secrets/authorization kv
vault auth enable approle

echo "Vault Setup Complete! If this is a production deployment, now is the time to: 
  1.) add some kv credentials 
  2.) block the cloud loadbalancer IP from public access 
  3.) check the stored values, access rights and roles through the CLI or the vault UI 
  4.) revoke the root token"

# Deploy plugin in vault
vault write sys/plugins/catalog/vault.plugin.linux.amd64 sha_256=${PLUGIN_SHASUM} command=vault.plugin.linux.amd64
vault secrets enable -path=aetheras-plugin -plugin-name=vault.plugin.linux.amd64 plugin