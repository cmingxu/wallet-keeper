package omnijson

type ListAccountsResult map[string]float64

type ListAccountsCommand struct {
	MinConf int64 `json:"minConf"`
}

func (ListAccountsCommand) Method() string {
	return "listaccounts"
}

func (ListAccountsCommand) ID() string {
	return "1"
}

func (cmd ListAccountsCommand) Params() []interface{} {
	return []interface{}{cmd.MinConf, 9999999}
}
