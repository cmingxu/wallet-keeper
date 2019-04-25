package omnijson

/*
Result:
"transaction"            (string) hex string of the transaction
*/

type CreateRawTransactionResult = string

type CreateRawTransactionCommand struct {
	Parameters []CreateRawTransactionParameter
}

type CreateRawTransactionParameter struct {
	Tx   string `json:"txid"`
	Vout uint32 `json:"vout"`
}

type createrawtransactionOutput struct {
	Address string `json:"address,omitempty"`
	Data    string `json:"data,omitempty"`
}

func (CreateRawTransactionCommand) Method() string {
	return "createrawtransaction"
}

func (CreateRawTransactionCommand) ID() string {
	return "1"
}

func (cmd CreateRawTransactionCommand) Params() []interface{} {
	return []interface{}{cmd.Parameters, createrawtransactionOutput{}}
}
