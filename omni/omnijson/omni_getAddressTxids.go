package omnijson

/*
Result
{
  [
 "hash",
  ]
}
*/

type OmniAddressTxidsResult struct {
	Hash string `json:"hash"`
}

type OmniAddressTxidsCommand struct {
	Addresses []string `json:"addresses"`
}

func (OmniAddressTxidsCommand) Method() string {
	return "getaddresstxids"
}

func (OmniAddressTxidsCommand) ID() string {
	return "1"
}

func (cmd OmniAddressTxidsCommand) Params() []interface{} {
	return []interface{}{cmd.Addresses}
}
