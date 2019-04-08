package btc

import (
	"github.com/cmingxu/wallet-keeper/keeper"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	log "github.com/sirupsen/logrus"
)

var DEFAULT_ACCOUNT = "duckduck"
var DEFAULT_CONFIRMATION = 6

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

func (client *Client) CreateAccount(account string) (keeper.Account, error) {
	address, err := client.GetAddress(account)
	if err != nil {
		return keeper.Account{}, err
	}

	return keeper.Account{
		Account:   account,
		Balance:   0.0,
		Addresses: []string{address},
	}, nil
}

func (client *Client) GetAccountInfo(account string) (keeper.Account, error) {
	var accountsMap map[string]float64
	var err error
	if accountsMap, err = client.ListAccountsMinConf(0); err != nil {
		return keeper.Account{}, err
	}

	balance, found := accountsMap[account]
	if !found {
		return keeper.Account{}, keeper.ErrAccountNotFound
	}

	addresses, err := client.GetAddressesByAccount(account)
	if err != nil {
		return keeper.Account{}, err
	}

	return keeper.Account{
		Account:   account,
		Balance:   balance,
		Addresses: addresses,
	}, nil
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

func (client *Client) ListAccountsMinConf(conf int) (map[string]float64, error) {
	accounts := make(map[string]float64)
	accountsWithAmount, err := client.rpcClient.ListAccountsMinConf(conf)
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
	log.Info("SendToAddressComment got hash", hash)

	return nil
}

// TODO check validity of account and have sufficent balance
func (client *Client) SendFrom(account, address string, amount float64) error {
	decoded, err := decodeAddress(address, chaincfg.TestNet3Params)
	if err != nil {
		return err
	}

	btcAmount, err := convertToBtcAmount(amount)
	if err != nil {
		return err
	}

	hash, err := client.rpcClient.SendFrom(account, decoded, btcAmount)
	if err != nil {
		return err
	}
	log.Info("SendFrom got hash", hash)

	return nil
}

func (client *Client) Move(from, to string, amount float64) (bool, error) {
	btcAmount, err := convertToBtcAmount(amount)
	if err != nil {
		return false, err
	}

	return client.rpcClient.Move(from, to, btcAmount)
}

func (client *Client) ListUnspentMin(minConf int) ([]btcjson.ListUnspentResult, error) {
	return client.rpcClient.ListUnspentMin(minConf)
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
