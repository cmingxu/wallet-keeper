package omnijson

/*
Result
[                       (array of string)
  "hash",                 (string) the hash of the transaction
  ...
]
*/

type OmniListBlockTransactionsResult = []string

type OmniListBlockTransactionsCommand struct {
	Block int64
}

func (OmniListBlockTransactionsCommand) Method() string {
	return "omni_listblocktransactions"
}

func (OmniListBlockTransactionsCommand) ID() string {
	return "1"
}

func (cmd OmniListBlockTransactionsCommand) Params() []interface{} {
	return []interface{}{cmd.Block}
}
