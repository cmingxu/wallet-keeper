package omnijson

/*
{
  "balance" : "n.nnnnnnnn",   (string) the available balance of the address
  "reserved" : "n.nnnnnnnn"   (string) the amount reserved by sell offers and accepts
  "frozen" : "n.nnnnnnnn"     (string) the amount frozen by the issuer (applies to managed properties only)
}
*/

type OmniGetBalanceResult = struct {
	Balance  string
	Reserved string
	Frozen   string
}

type OmniGetBalanceCommand struct {
	Address    string
	PropertyID int32
}

func (OmniGetBalanceCommand) Method() string {
	return "omni_getbalance"
}

func (OmniGetBalanceCommand) ID() string {
	return "1"
}

func (cmd OmniGetBalanceCommand) Params() []interface{} {
	return []interface{}{cmd.Address, cmd.PropertyID}
}
