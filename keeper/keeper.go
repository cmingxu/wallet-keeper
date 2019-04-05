package keeper

type Keeper interface {
	GetBlockCount() (int64, error)

	// check if the coin core service avaliable or not,
	// the error might caused by misconfiguration or
	// runtime error. Error happend indicates fatal and
	// could not recover, suicide might be the best choice.
	Ping() error

	// Returns a valid coin address for both receive from other.
	GetAddress() (string, error)

	// Returns a valid coin address for both receive from other.
	GetAddresses() (map[string]string, error)

	// List all accounts/labels together with how much satoshi remains.
	ListAccounts() (map[string]float64, error)
}
