package omnijson

/*
Result
[                   (array of json object)
 "utxos": [
    {
      "address": "base58address", // (string) The address
      "txid": "hash"              // (string) The output txid
      "outputIndex": n,           // (number) The block height
      "script": "hex",            // (string) The script hex encoded
      "satoshis": n,              // (number) The number of satoshis of the output
      "height": 301,              // (number) The block height
      "coinbase": true            // (boolean) Whether it's a coinbase transaction
    },
    ...
  ],
  "hash": "hash",                 // (string) The current block hash
  "height": n                     // (string) The current block height
  ,...
]
*/

type utxos = []struct {
	Address     string  `json:"address"`
	Tx          string  `json:"txid"`
	OutputIndex int64   `json:"outputIndex"`
	Script      string  `json:"script"`
	Satoshis    float64 `json:"satoshis"`
	Height      int64   `json:"height"`
	Vout        uint32  `json:"vout"`
	Coinbase    bool    `json:"coinbase"`
}

type OmniListUnSpentResult struct {
	Utxos  utxos  `json:"utxos"`
	Hash   string `json:"hash"`
	Height string `json:"height"`
}

type OmniListUnspentCommand struct {
	Addresses []string
}

func (OmniListUnspentCommand) Method() string {
	return "getaddressutxos"
}

func (OmniListUnspentCommand) ID() string {
	return "1"
}

func (cmd OmniListUnspentCommand) Params() []interface{} {
	return []interface{}{cmd.Addresses}
}
