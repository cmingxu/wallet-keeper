package omnijson

/*
Result
{
	"result":"mkVW5QARPvvAY328y9LMHep5ZhuiTrgdd4",
	"error":null,
	"id":"curltext"
}
*/

//GetNewAddressResult ...
type GetNewAddressResult = string

//GetNewAddressCommand ...
type GetNewAddressCommand struct {
	Account string
}

//Method ...
func (GetNewAddressCommand) Method() string {
	return "getnewaddress"
}

//ID ...
func (GetNewAddressCommand) ID() string {
	return "1"
}

//Params ...
func (cmd GetNewAddressCommand) Params() []interface{} {
	return []interface{}{cmd.Account}
}
