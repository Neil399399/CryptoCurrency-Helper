package omnijson

// [                          // (array of JSON objects)
//   {
//     "propertyid" : n,          // (number) the property identifier
//     "name" : "name",           // (string) the name of the property
//     "balance" : "n.nnnnnnnn",  // (string) the available balance of the address
//     "reserved" : "n.nnnnnnnn", // (string) the amount reserved by sell offers and accepts
//     "frozen" : "n.nnnnnnnn"    // (string) the amount frozen by the issuer (applies to managed properties only)
//   },
//   ...
// ]

type OmniGetAllBalancesForAddressResult = []struct {
	PropertyId int64
	Name       string
	Balance    string
	Reserved   string
	Frozen     string
}

type OmniGetAllBalancesForAddressCommand struct {
	Address string
}

func (OmniGetAllBalancesForAddressCommand) Method() string {
	return "omni_getallbalancesforaddress"
}

func (OmniGetAllBalancesForAddressCommand) ID() string {
	return "1"
}

func (cmd OmniGetAllBalancesForAddressCommand) Params() []interface{} {
	return []interface{}{cmd.Address}
}
