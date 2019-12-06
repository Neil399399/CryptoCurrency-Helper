package bitcoin

// VoutDetail is bitcoin transaction detail
type VoutDetail struct {
	BlockHash     string
	Txid          string
	Address       []string
	Category      string
	Amount        float64
	Vout          uint32
	Confirmations uint64
	Time          int64
	LockTime      uint32
	Blocktime     int64
}

// Utxo is bitcoin Unspent Transaction Output
type Utxo struct {
	Address     string
	TxID        string
	OutputIndex uint32
	Script      []byte
	Satoshis    int64
}
