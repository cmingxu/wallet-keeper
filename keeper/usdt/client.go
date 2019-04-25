package usdt

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/cmingxu/wallet-keeper/keeper"
	"github.com/cmingxu/wallet-keeper/omnilayer"
	"github.com/cmingxu/wallet-keeper/omnilayer/omnijson"

	"github.com/btcsuite/btcd/btcjson"
	log "github.com/sirupsen/logrus"
)

// default account for reserved usage, which represent
// account belongs to enterpise default
var DEFAULT_ACCOUNT = "duckduck"

// default confirmation
var DEFAULT_CONFIRMATION = 1

type Client struct {
	rpcClient *omnilayer.Client
	l         *log.Logger

	propertyId int64
}

// GetBlockCount
func (client *Client) GetBlockCount() (int64, error) {
	var res omnijson.GetBlockChainInfoResult
	if res, err := client.rpcClient.GetBlockChainInfo(); err == nil {
		return res.Blocks, nil
	}
	return res.Blocks, nil
}

// connect to omnicore with HTTP RPC transport
func NewClient(host, user, pass, logDir string, propertyId int64) (*Client, error) {
	connCfg := &omnilayer.ConnConfig{
		Host: host,
		User: user,
		Pass: pass,
	}

	client := &Client{propertyId: propertyId}
	client.rpcClient = omnilayer.New(connCfg)

	logPath := filepath.Join(logDir, "usdt.log")
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
	_, err := client.rpcClient.GetInfo()
	return err
}

// GetAddress - default address
func (client *Client) GetAddress(account string) (string, error) {
	if len(account) == 0 {
		account = DEFAULT_ACCOUNT
	}

	address, err := client.rpcClient.GetNewAddress(account)
	if err != nil {
		return "", err
	}

	return address, nil
}

// Create Account
// Returns customized account info
func (client *Client) CreateAccount(account string) (keeper.Account, error) {
	// GetAddress will create account if not exists
	client.l.Infof("[CreateAccount] for account %s", account)
	address, err := client.GetAddress(account)
	if err != nil {
		return keeper.Account{}, err
	}

	client.l.Infof("[CreateAccount] success for account %s", account)
	return keeper.Account{
		Account:   account,
		Balance:   0.0,
		Addresses: []string{address},
	}, nil
}

// GetAccountInfo
func (client *Client) GetAccountInfo(account string, minConf int) (keeper.Account, error) {
	addresses, err := client.GetAddressesByAccount(account)
	if err != nil {
		return keeper.Account{}, err
	}

	var balance float64
	balance = 0
	for _, addr := range addresses {
		cmd := omnijson.OmniGetBalanceCommand{
			Address:    addr,
			PropertyID: int32(client.propertyId),
		}
		if curBalance, err := client.rpcClient.OmniGetBalance(cmd); err == nil {
			if b, err := strconv.ParseFloat(curBalance.Balance, 64); err == nil {
				balance += b
			}
		}
	}

	return keeper.Account{
		Account:   account,
		Balance:   balance,
		Addresses: addresses,
	}, nil
}

// GetAddressesByAccount
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
		addrs = append(addrs, addr)
	}

	return addrs, nil
}

// GetNewAddress ...
func (client *Client) GetNewAddress(account string) (string, error) {
	client.l.Infof("[GetNewAddress] for account %s", account)
	if len(account) == 0 {
		account = DEFAULT_ACCOUNT
	}

	address, err := client.rpcClient.GetNewAddress(account)
	if err != nil {
		return "", err
	}

	client.l.Infof("[GetNewAddress] for account %s, address is %s", account, address)
	return address, nil
}

// ListAccountsMinConf
// USDT RPC don't need this Stub func
func (client *Client) ListAccountsMinConf(conf int) (map[string]float64, error) {
	accounts := make(map[string]float64)

	return accounts, nil
}

// SendToAddress ...
// USDT RPC don't need this Stub func
func (client *Client) SendToAddress(address string, amount float64) error {
	client.l.Infof("[SendToAddress] to address %s: %f", address, amount)
	return nil
}

//SendFrom ...omni_funded_send
func (client *Client) SendFrom(account, address string, amount float64) error {
	client.l.Infof("[SendFrom] from account %s to address %s with amount %f ", account, address, amount)
	addresses, err := client.rpcClient.GetAddressesByAccount(account)
	if err != nil {
		return err
	}
	for _, addr := range addresses {
		hash, _ := client.rpcClient.OmniFoundedSend(addr, address, client.propertyId, floatToString(amount), addr)
		client.l.Infof("[SendFrom] from account %s to address %s with amount %f, result hash %s ", account, addr, amount, hash)
	}
	return nil
}

// Move - omni_funded_send
func (client *Client) Move(from, to string, amount float64) (bool, error) {
	client.l.Infof("[Move] from %s to %s with amount %f ", from, to, amount)
	hash, err := client.rpcClient.OmniFoundedSend(from, to, client.propertyId, floatToString(amount), to)
	if err != nil {
		return false, err
	}

	client.l.Infof("[Move] success from %s to %s with amount %f hash is %s ", from, to, amount, hash)
	return true, nil
}

// ListUnspentMin
func (client *Client) ListUnspentMin(minConf int) ([]btcjson.ListUnspentResult, error) {
	return nil, nil
}

//util
func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
