package omnijson

/*
Result
{
	"result":"mkVW5QARPvvAY328y9LMHep5ZhuiTrgdd4",
	"error":null,
	"id":"curltext"
}
*/

//GetAccountAddressResult ...
type GetAccountAddressResult = string

//GetAccountAddressCommand ...
type GetAccountAddressCommand struct {
	Account string
}

//Method ...
func (GetAccountAddressCommand) Method() string {
	return "getaccountaddress"
}

//ID ...
func (GetAccountAddressCommand) ID() string {
	return "1"
}

//Params ...
func (cmd GetAccountAddressCommand) Params() []interface{} {
	return []interface{}{cmd.Account}
}
