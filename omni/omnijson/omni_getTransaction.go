package omnijson

/*
{
  "txid" : "hash",                 // (string) the hex-encoded hash of the transaction
  "sendingaddress" : "address",    // (string) the Bitcoin address of the sender
  "referenceaddress" : "address",  // (string) a Bitcoin address used as reference (if any)
  "ismine" : true|false,           // (boolean) whether the transaction involes an address in the wallet
  "confirmations" : nnnnnnnnnn,    // (number) the number of transaction confirmations
  "fee" : "n.nnnnnnnn",            // (string) the transaction fee in bitcoins
  "blocktime" : nnnnnnnnnn,        // (number) the timestamp of the block that contains the transaction
  "valid" : true|false,            // (boolean) whether the transaction is valid
  "positioninblock" : n,           // (number) the position (index) of the transaction within the block
  "version" : n,                   // (number) the transaction version
  "type_int" : n,                  // (number) the transaction type as number
  "type" : "type",                 // (string) the transaction type as string
  [...]                            // (mixed) other transaction type specific properties
}

// real response
{
"txid":"c0083a2213527f205fb77819220d58395d593cfb1d25dd0d9a264bc6e68b18e2",
"fee":"0.00100000",
"sendingaddress":"mp5mRGJtfSrSbu3ngLccFJhCgagRRzcFMZ",
"ismine":false,
"version":0,
"type_int":54,
"type":"Create Property - Manual",
"propertytype":"unknown",
"ecosystem":"main",
"category":"",
"subcategory":"",
"propertyname":"vatomic.prototyping::ImageCard_BC::Test01",
"data":"",
"url":"",
"amount":"0",
"valid":false,
"invalidreason":"Invalid property type",
"blockhash":"00000000000003e121dd0ce9ad1951185973ed5918f9a01fdb720fb74608fba4",
"blocktime":1470927286,
"positioninblock":1,
"block":921866,
"confirmations":417
}

*/

type OmniGetTransactionResult = struct {
	TxID             string `json:"txid"`
	SendingAddress   string `json:"sendingaddress"`
	ReferenceAddress string `json:"referenceaddress"`
	Ismine           bool   `json:"ismine"`
	Confirmations    int64  `json:"confirmations"`
	Amount           string `json:"amount"`
	Fee              string `json:"fee"`
	Block            int64  `json:"block"`
	BlockHash        string `json:"blockhash"`
	BlockTime        int64  `json:"blocktime"`
	Vaild            bool   `json:"valid"`
	Positioninblock  int64  `json:"positioninblock"`
	Version          int32  `json:"version"`
	TypeInt          int32  `json:"type_int"`
	Type             string `json:"type"`
}

type OmniGetTransactionCommand struct {
	TxID string
}

func (OmniGetTransactionCommand) Method() string {
	return "omni_gettransaction"
}

func (OmniGetTransactionCommand) ID() string {
	return "1"
}

func (cmd OmniGetTransactionCommand) Params() []interface{} {
	return []interface{}{cmd.TxID}
}
