package omnijson

type SendRawTransactionResult = string

type SendRawTransactionCommand struct {
	Hex           string
	AllowHighFees bool
}

func (SendRawTransactionCommand) Method() string {
	return "sendrawtransaction"
}

func (SendRawTransactionCommand) ID() string {
	return "1"
}

func (cmd SendRawTransactionCommand) Params() []interface{} {
	return []interface{}{cmd.Hex, cmd.AllowHighFees}
}
