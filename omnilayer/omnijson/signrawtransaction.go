package omnijson

/*
{
  "hex" : "value",           (string) The hex-encoded raw transaction with signature(s)
  "complete" : true|false,   (boolean) If the transaction has a complete set of signatures
  "errors" : [                 (json array of objects) Script verification errors (if there are any)
    {
      "txid" : "hash",           (string) The hash of the referenced, previous transaction
      "vout" : n,                (numeric) The index of the output to spent and used as input
      "scriptSig" : "hex",       (string) The hex-encoded signature script
      "sequence" : n,            (numeric) Script sequence number
      "error" : "text"           (string) Verification or signing error related to the input
    }
    ,...
  ]
}
*/

type SignRawTransactionResult = struct {
	Hex      string                    `json:"hex"`
	Complete bool                      `json:"complete"`
	Errors   []signRawTransactionError `json:"errors"`
}

type signRawTransactionError struct {
	TxID      string `json:"txid"`
	ScriptSig string `json:"scriptSig"`
	Error     string `json:"error"`
	Vout      uint32 `json:"vout"`
	Sequence  uint32 `json:"sequence"`
}

type SignRawTransactionCommand struct {
	Hex      string
	Previous []Previous
	Keys     []string
	Type     string
}

type Previous struct {
	TxID         string  `json:"txid"`
	Vout         uint32  `json:"vout"`
	ScriptPubKey string  `json:"scriptPubKey"`
	RedeemScript string  `json:"redeemScript"`
	Value        float64 `json:"value"`
}

func (SignRawTransactionCommand) Method() string {
	return "signrawtransaction"
}

func (SignRawTransactionCommand) ID() string {
	return "1"
}

func (cmd SignRawTransactionCommand) Params() []interface{} {
	return []interface{}{cmd.Hex, cmd.Previous, cmd.Keys, cmd.Type}
}
