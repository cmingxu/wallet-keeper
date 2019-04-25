package omnijson

/*
Result
[                   (array of json object)
  {
    "txid" : "txid",          (string) the transaction id
    "vout" : n,               (numeric) the vout value
    "address" : "address",    (string) the bitcoin address
    "account" : "account",    (string) DEPRECATED. The associated account, or "" for the default account
    "scriptPubKey" : "key",   (string) the script key
    "amount" : x.xxx,         (numeric) the transaction amount in BTC
    "confirmations" : n,      (numeric) The number of confirmations
    "redeemScript" : n        (string) The redeemScript if scriptPubKey is P2SH
    "spendable" : xxx,        (bool) Whether we have the private keys to spend this output
    "solvable" : xxx          (bool) Whether we know how to spend this output, ignoring the lack of keys
  }
  ,...
]
*/

type ListUnspentResult = []struct {
	Tx            string  `json:"txid"`
	Address       string  `json:"address"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	RedeemScript  string  `json:"redeemScript"`
	Amount        float64 `json:"amount"`
	Confirmations int64   `json:"confirmations"`
	Vout          uint32  `json:"vout"`
	Spendable     bool    `json:"spendable"`
	Solvable      bool    `json:"solvable"`
}

type ListUnspentCommand struct {
	Min       int
	Addresses []string
}

func (ListUnspentCommand) Method() string {
	return "listunspent"
}

func (ListUnspentCommand) ID() string {
	return "1"
}

func (cmd ListUnspentCommand) Params() []interface{} {
	return []interface{}{cmd.Min, 9999999, cmd.Addresses}
}
