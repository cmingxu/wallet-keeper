package cryptocurrency

type Wallet interface {
	NewClient(config ClientConfig) Client

	GetBlockCount()
}
