package ethereum

type Event struct {
	Network         string
	IsERC20         bool
	ContractAddress string
	AddressBook     []string
	ReferenceID     []string
	LastBlock       uint64
}

type TxDetail struct {
	network     string
	from        string
	to          string
	txnHash     string
	value       uint64
	symbol      string
	blockNumber uint64
	timeStamp   string
	referenceID string
	erc         bool
}
