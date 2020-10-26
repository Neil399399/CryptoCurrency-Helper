set -e
echo "Vault Network: ${VAULT_ADDR}"
echo "Init dev key in vault"

BTC_KEY_1="walnut front vocal present level asset evoke poem two sentence trend enforce menu include supply"
BTC_KEY_2="essence pulp flag slice strike wheel upon category turn mushroom diet review knife emotion brown"
BTC_KEY_3="agree pottery copper thunder ski park shrug time farm caught powder shrimp gym photo kind"
ETH_KEY_1="midnight slender oil ranch brain legend brother cube arrive catalog october bronze spatial napkin involve"
ETH_KEY_2="tortoise tank scout spray trim total area defy life trend relax produce sudden tenant indoor"
ETH_KEY_3="dance fan cannon manual segment liberty stomach abstract return exhaust make payment style exercise survey"

# for admin
BTC_KEY_4="agree pottery copper thunder ski park shrug time farm caught powder shrimp gym photo kind"
ETH_KEY_4="dance fan cannon manual segment liberty stomach abstract return exhaust make payment style exercise survey"

# Deploy Key in vault
vault write aetheras-plugin/btc/hot/kv/aetheras_btc_1 value="${BTC_KEY_1}"
vault write aetheras-plugin/btc/hot/kv/aetheras_btc_2 value="${BTC_KEY_2}"
vault write aetheras-plugin/btc/hot/kv/aetheras_btc_3 value="${BTC_KEY_3}"
vault write aetheras-plugin/eth/hot/kv/aetheras_eth_1 value="${ETH_KEY_1}"
vault write aetheras-plugin/eth/hot/kv/aetheras_eth_2 value="${ETH_KEY_2}"
vault write aetheras-plugin/eth/hot/kv/aetheras_eth_3 value="${ETH_KEY_3}"

# Deploy admin key in vault
vault write aetheras-plugin/btc/hot/kv/aetheras_btc_4 value="${BTC_KEY_4}"
vault write aetheras-plugin/eth/hot/kv/aetheras_eth_4 value="${ETH_KEY_4}"
vault write aetheras-plugin/generate/btc/aetheras_btc_4 network="testnet" childIdx="89"
vault write aetheras-plugin/generate/eth/aetheras_eth_4 childIdx="89"
vault write aetheras-plugin/generate/btc/aetheras_btc_4 network="testnet" childIdx="90"
vault write aetheras-plugin/generate/eth/aetheras_eth_4 childIdx="90"
vault write aetheras-plugin/generate/btc/aetheras_btc_4 network="testnet" childIdx="31415"
vault write aetheras-plugin/generate/eth/aetheras_eth_4 childIdx="31415"
vault write aetheras-plugin/generate/btc/aetheras_btc_4 network="testnet" childIdx="314159"
vault write aetheras-plugin/generate/eth/aetheras_eth_4 childIdx="314159"
vault write aetheras-plugin/generate/btc/aetheras_btc_4 network="testnet" childIdx="3141592"
vault write aetheras-plugin/generate/eth/aetheras_eth_4 childIdx="3141592"
vault write aetheras-plugin/generate/btc/aetheras_btc_4 network="testnet" childIdx="31415926"
vault write aetheras-plugin/generate/eth/aetheras_eth_4 childIdx="31415926"
vault write aetheras-plugin/generate/btc/aetheras_btc_4 network="testnet" childIdx="314159265"
vault write aetheras-plugin/generate/eth/aetheras_eth_4 childIdx="314159265"