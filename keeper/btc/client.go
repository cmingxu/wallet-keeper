package btc

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	log "github.com/sirupsen/logrus"
)

var DEFAULT_ACCOUNT = "duckduck"

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

func (client *Client) GetAddress(account string) (string, error) {
	if len(account) == 0 {
		account = DEFAULT_ACCOUNT
	}

	address, err := client.rpcClient.GetAccountAddress(DEFAULT_ACCOUNT)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

// TODO
// GetNewAddress does map to `getnewaddress` rpc call now
// rpcclient doesn't have such golang wrapper func.
func (client *Client) GetNewAddress(account string) (string, error) {
	if len(account) == 0 {
		account = DEFAULT_ACCOUNT
	}

	address, err := client.rpcClient.GetNewAddress(DEFAULT_ACCOUNT)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

func (client *Client) GetAddressesByAccount(account string) ([]string, error) {
	if len(account) == 0 {
		account = DEFAULT_ACCOUNT
	}
	log.Println(account)
	addresses, err := client.rpcClient.GetAddressesByAccount(account)
	if err != nil {
		return []string{}, err
	}

	addrs := make([]string, 0)
	for _, addr := range addresses {
		addrs = append(addrs, addr.String())
	}

	return addrs, nil
}

func (client *Client) ListAccounts() (map[string]float64, error) {
	accounts := make(map[string]float64)
	accountsWithAmount, err := client.rpcClient.ListAccounts()
	if err != nil {
		return accounts, err
	}

	for account, amount := range accountsWithAmount {
		accounts[account] = amount.ToBTC()
	}

	return accounts, nil
}

func (client *Client) SendToAddress(address string, amount float64) error {
	decoded, err := decodeAddress(address, chaincfg.TestNet3Params)
	if err != nil {
		return err
	}

	btcAmount, err := convertToBtcAmount(amount)
	if err != nil {
		return err
	}

	hash, err := client.rpcClient.SendToAddressComment(decoded, btcAmount, "comment", "commentto")
	if err != nil {
		return err
	}
	log.Println(hash)

	return nil
}

func decodeAddress(address string, cfg chaincfg.Params) (btcutil.Address, error) {
	decodedAddress, err := btcutil.DecodeAddress(address, &cfg)
	if err != nil {
		return nil, err
	}

	return decodedAddress, nil
}

func convertToBtcAmount(amount float64) (btcutil.Amount, error) {
	return btcutil.NewAmount(amount)
}
