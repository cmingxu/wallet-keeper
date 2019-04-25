package omnijson

/*
Result
{"result":["mkVW5QARPvvAY328y9LMHep5ZhuiTrgdd4"],"error":null,"id":"curltext"}
*/

//GetAddressesByAccountResult ...
type GetAddressesByAccountResult = []string

//GetAddressesByAccountCommand ...
type GetAddressesByAccountCommand struct {
	Account string
}

//Method ...
func (GetAddressesByAccountCommand) Method() string {
	return "getaddressesbyaccount"
}

//ID ...
func (GetAddressesByAccountCommand) ID() string {
	return "1"
}

//Params ...
func (cmd GetAddressesByAccountCommand) Params() []interface{} {
	return []interface{}{cmd.Account}
}
