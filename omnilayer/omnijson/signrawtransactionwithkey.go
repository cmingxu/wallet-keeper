package omnijson

type SignRawTransactionWithKeyResult = string

type SignRawTransactionWithKeyCommand struct {
	Hex      string
	Previous []Previous
	Keys     []string
	Type     string
}

func (SignRawTransactionWithKeyCommand) Method() string {
	return "signrawtransactionwithkey"
}

func (SignRawTransactionWithKeyCommand) ID() string {
	return "1"
}

func (cmd SignRawTransactionWithKeyCommand) Params() []interface{} {
	return []interface{}{cmd.Hex, cmd.Keys, cmd.Previous, cmd.Type}
}
