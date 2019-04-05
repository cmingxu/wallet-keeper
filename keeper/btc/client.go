package btc

import (
	"github.com/btcsuite/btcd/rpcclient"
)

var DEGAULT_ACCOUNT = "duckduck"

type Client struct {
	rpcClient *rpcclient.Client
}

// connect to bitcoind with HTTP RPC transport
func NewClient(host, user, pass string) (*Client, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         host,
		User:         user,
		Pass:         pass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client := &Client{}
	var err error
	client.rpcClient, err = rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}

	// check if bitcoind response
	err = client.rpcClient.Ping()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (client *Client) Ping() error {
	return client.rpcClient.Ping()
}

func (client *Client) GetBlockCount() (int64, error) {
	return client.rpcClient.GetBlockCount()
}

func (client *Client) GetAddress() (string, error) {
	address, err := client.rpcClient.GetAccountAddress(DEGAULT_ACCOUNT)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

func (client *Client) GetAddresses() (map[string]string, error) {
	return make(map[string]string), nil
}

func (client *Client) ListAccounts() (map[string]float64, error) {
	var accounts map[string]float64
	accountsWithAmount, err := client.rpcClient.ListAccounts()
	if err != nil {
		return accounts, err
	}

	for account, amount := range accountsWithAmount {
		accounts[account] = amount.ToBTC()
	}

	return accounts, nil
}
