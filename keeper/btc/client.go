package btc

import (
	"os"
	"path/filepath"

	"github.com/cmingxu/wallet-keeper/keeper"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	log "github.com/sirupsen/logrus"
)

// default account for reserved usage, which represent
// account belongs to enterpise default
var DEFAULT_ACCOUNT = "duckduck"

// default confirmation
var DEFAULT_CONFIRMATION = 6

type Client struct {
	rpcClient *rpcclient.Client
	l         *log.Logger
}

// connect to bitcoind with HTTP RPC transport
func NewClient(host, user, pass, logDir string) (*Client, error) {
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

	logPath := filepath.Join(logDir, "btc.log")
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return nil, err
	}
	client.l = &log.Logger{
		Out:       logFile,
		Level:     log.DebugLevel,
		Formatter: new(log.JSONFormatter),
	}

	return client, nil
}

// Ping
func (client *Client) Ping() error {
	return client.rpcClient.Ping()
}

// GetBlockCount
func (client *Client) GetBlockCount() (int64, error) {
	client.l.Infof("[GetBlockCount]")
	return client.rpcClient.GetBlockCount()
}

// GetAddress - default address
func (client *Client) GetAddress(account string) (string, error) {
	if len(account) == 0 {
		account = DEFAULT_ACCOUNT
	}
	client.l.Infof("[GetAddress] for account %s", account)

	address, err := client.rpcClient.GetAccountAddress(account)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

// Create Account
// Returns customized account info
func (client *Client) CreateAccount(account string) (keeper.Account, error) {
	client.l.Infof("[CreateAccount] for account %s", account)
	// GetAddress will create account if not exists
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

// GetAccountInfo
func (client *Client) GetAccountInfo(account string, minConf int) (keeper.Account, error) {
	var accountsMap map[string]float64
	var err error

	client.l.Infof("[GetAccountInfo] account %s, with minConf %d", account, minConf)
	if accountsMap, err = client.ListAccountsMinConf(minConf); err != nil {
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

	address, err := client.rpcClient.GetNewAddress(account)
	if err != nil {
		return "", err
	}

	client.l.Infof("[GetNewAddress] for account %s, address is %s", account, address.String())
	return address.String(), nil
}

// GetAddressesByAccount
func (client *Client) GetAddressesByAccount(account string) ([]string, error) {
	if len(account) == 0 {
		account = DEFAULT_ACCOUNT
	}

	client.l.Infof("[GetAddressesByAccount] for account %s", account)
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

// ListAccountsMinConf
func (client *Client) ListAccountsMinConf(minConf int) (map[string]float64, error) {
	accounts := make(map[string]float64)

	client.l.Infof("[ListAccountsMinConf] with minConf %d", minConf)
	accountsWithAmount, err := client.rpcClient.ListAccountsMinConf(minConf)
	if err != nil {
		return accounts, err
	}

	for account, amount := range accountsWithAmount {
		accounts[account] = amount.ToBTC()
	}

	return accounts, nil
}

// SendToAddress
func (client *Client) SendToAddress(address string, amount float64) (string, error) {
	client.l.Infof("[SendToAddress] to address %s: %f", address, amount)
	decoded, err := decodeAddress(address, chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	btcAmount, err := convertToBtcAmount(amount)
	if err != nil {
		return "", err
	}

	hash, err := client.rpcClient.SendToAddressComment(decoded, btcAmount, "", "")
	if err != nil {
		return "", err
	}
	client.l.Infof("[SendToAddress] to address %s: %f, hash is %s", address, amount, hash)

	return hash.String(), nil
}

// TODO check validity of account and have sufficent balance
func (client *Client) SendFrom(account, address string, amount float64) (string, error) {
	client.l.Infof("[SendFrom] from account %s to address %s with amount %f ", account, address, amount)
	decoded, err := decodeAddress(address, chaincfg.TestNet3Params)
	if err != nil {
		return "", err
	}

	btcAmount, err := convertToBtcAmount(amount)
	if err != nil {
		return "", err
	}

	hash, err := client.rpcClient.SendFrom(account, decoded, btcAmount)
	if err != nil {
		return "", err
	}
	client.l.Infof("[SendFrom] from account %s to address %s with amount %f, result hash %s ", account, address, amount, hash)

	return hash.String(), nil
}

// Move
func (client *Client) Move(from, to string, amount float64) (bool, error) {
	client.l.Infof("[Move] from %s to %s with amount %f ", from, to, amount)
	btcAmount, err := convertToBtcAmount(amount)
	if err != nil {
		return false, err
	}

	client.l.Infof("[Move] success from %s to %s with amount %f ", from, to, amount)
	return client.rpcClient.Move(from, to, btcAmount)
}

// ListUnspentMin
func (client *Client) ListUnspentMin(minConf int) ([]btcjson.ListUnspentResult, error) {
	return client.rpcClient.ListUnspentMin(minConf)
}

// decodeAddress from string to decodedAddress
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
