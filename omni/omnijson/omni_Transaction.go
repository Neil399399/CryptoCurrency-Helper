package omnijson

/*
Result
{
"payload"  // (string) the hex-encoded payload
}
*/

type OmniCreatePayloadResult string

type OmniCreatePayloadCommand struct {
	PropertyId int32
	Amount     string
}

func (OmniCreatePayloadCommand) Method() string {
	return "omni_createpayload_simplesend"
}

func (OmniCreatePayloadCommand) ID() string {
	return "1"
}

func (cmd OmniCreatePayloadCommand) Params() []interface{} {
	return []interface{}{cmd.PropertyId, cmd.Amount}
}

/*
Result
{
"rawtx"  // (string) the hex-encoded modified raw transaction
}
*/

type OmniCreateOpReturnResult string

type OmniCreateOpReturnCommand struct {
	TxID    string
	Payload string
}

func (OmniCreateOpReturnCommand) Method() string {
	return "omni_createrawtx_opreturn"
}

func (OmniCreateOpReturnCommand) ID() string {
	return "1"
}

func (cmd OmniCreateOpReturnCommand) Params() []interface{} {
	return []interface{}{cmd.TxID, cmd.Payload}
}

/*
Result
{
"rawtx"  // (string) the hex-encoded modified raw transaction
}
*/

type OmniCreateInputResult string

type OmniCreateInputCommand struct {
	RawTx string
	TxID  string
	N     int32
}

func (OmniCreateInputCommand) Method() string {
	return "omni_createrawtx_input"
}

func (OmniCreateInputCommand) ID() string {
	return "1"
}

func (cmd OmniCreateInputCommand) Params() []interface{} {
	return []interface{}{cmd.RawTx, cmd.TxID, cmd.N}
}

/*
Result
{
"hash"  // (string) the hex-encoded transaction hash
}
*/

type OmniSendTransactionResult string

type OmniSendTransactionCommand struct {
	FromAddress     string
	ToAddress       string
	PropertyID      int32
	Amount          string
	RedeemAddress   string
	ReferenceAmount string
}

func (OmniSendTransactionCommand) Method() string {
	return "omni_createrawtx_input"
}

func (OmniSendTransactionCommand) ID() string {
	return "1"
}

func (cmd OmniSendTransactionCommand) Params() []interface{} {
	return []interface{}{cmd.FromAddress, cmd.ToAddress, cmd.PropertyID, cmd.Amount}
}

/*
Result
{
"hash"  // (string) the hex-encoded transaction hash
}
*/

type OmniSendRawTransactionResult string

type OmniSendRawTransactionCommand struct {
	FromAddress      string
	RawTx            string
	ReferenceAddress string
	RedeemAddress    string
	ReferenceAmount  string
}

func (OmniSendRawTransactionCommand) Method() string {
	return "omni_sendrawtx"
}

func (OmniSendRawTransactionCommand) ID() string {
	return "1"
}

func (cmd OmniSendRawTransactionCommand) Params() []interface{} {
	return []interface{}{cmd.FromAddress, cmd.RawTx, cmd.ReferenceAddress, cmd.RedeemAddress,cmd.ReferenceAmount}
}
