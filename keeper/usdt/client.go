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

// GetBlockCount
func (client *Client) GetBlockCount() (int64, error) {
	client.l.Infof("[GetBlockCount]")

	var res omnijson.GetBlockChainInfoResult
	if res, err := client.rpcClient.GetBlockChainInfo(); err == nil {
		return res.Blocks, nil
	}
	return res.Blocks, nil
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
	client.l.Infof("[GetAddress] for account %s", account)

	address, err := client.rpcClient.GetAccountAddress(account)
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

	client.l.Infof("[GetAccountInfo] account %s, with minConf %d", account, minConf)
	var balance float64 = 0
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

	client.l.Infof("[GetAddressesByAccount] for account %s", account)
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
func (client *Client) ListAccountsMinConf(minConf int) (map[string]float64, error) {
	accounts := make(map[string]float64)

	client.l.Infof("[ListAccountsMinConf] with minConf %d", minConf)
	accountsWithAmount, err := client.rpcClient.ListAccounts(int64(minConf))
	if err != nil {
		return accounts, err
	}

	for account, _ := range accountsWithAmount {
		var accountInfo keeper.Account
		accountInfo, err = client.GetAccountInfo(account, minConf)
		if err != nil {
			accounts[account] = -1
		} else {
			accounts[account] = accountInfo.Balance
		}
	}

	return accounts, nil
}

// SendToAddress ...
// USDT RPC don't need this Stub func
func (client *Client) SendToAddress(address string, amount float64) (string, error) {
	return "", keeper.ErrNotSupport
}

//SendFrom ...omni_funded_send
func (client *Client) SendFrom(account, address string, amount float64) (string, error) {
	client.l.Infof("[SendFrom] from account %s to address %s with amount %f ", account, address, amount)
	fromAddr, err := client.rpcClient.GetAccountAddress(account)
	if err != nil {
		client.l.Errorf("[SendFrom] go error: %s", err)
		return "", err
	}

	hash, _ := client.rpcClient.OmniFoundedSend(fromAddr, address, client.propertyId, floatToString(amount), fromAddr)
	client.l.Infof("[SendFrom] from account %s to address %s with amount %f, result hash %s ", fromAddr, address, amount, hash)

	return hash, nil
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
