package keeper

type Keeper interface {
	GetBlockCount() (int64, error)

	// check if the coin core service avaliable or not,
	// the error might caused by misconfiguration or
	// runtime error. Error happend indicates fatal and
	// could not recover, suicide might be the best choice.
	Ping() error

	// Returns address under accont, use default account if
	// not provided
	GetAddress(account string) (string, error)

	// Return new address under account
	GetNewAddress(account string) (string, error)

	GetAddresses(account string) ([]string, error)

	// List all accounts/labels together with how much satoshi remains.
	ListAccounts() (map[string]float64, error)
}
