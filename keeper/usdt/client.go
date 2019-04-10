package usdt


import (
	"github.com/cmingxu/wallet-keeper/keeper"
	log "github.com/sirupsen/logrus"
	omnilayer "github.com/xiaods/omnilayer-go"
	"github.com/xiaods/omnilayer-go/omnijson"
	"strconv"
	"github.com/btcsuite/btcd/btcjson"
)


// default account for reserved usage, which represent
// account belongs to enterpise default
var DEFAULT_ACCOUNT = "duckduck"

// default confirmation
var DEFAULT_CONFIRMATION = 1

// default property id
var USDT_PROPERTY_ID = 2

type Client struct {
	rpcClient *omnilayer.Client
}


// GetBlockCount
func (client *Client) GetBlockCount() (int64, error) {
	var res omnijson.GetBlockChainInfoResult
	if res, err := client.rpcClient.GetBlockChainInfo(); err == nil {
		return res.Blocks,nil
	}
	return res.Blocks,nil
}

// connect to omnicore with HTTP RPC transport
func NewClient(host, user, pass string) (*Client, error) {
	connCfg := &omnilayer.ConnConfig{
		Host:         host,
		User:         user,
		Pass:         pass,
	}

	client := &Client{}
	client.rpcClient = omnilayer.New(connCfg)

	return client, nil
}

// Ping
func (client *Client) Ping() error {
	_ , err := client.rpcClient.GetInfo()
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
func (client *Client) GetAccountInfo(account string) (keeper.Account, error) {
	addresses, err := client.GetAddressesByAccount(account)
	if err != nil {
		return keeper.Account{}, err
	}

	var balance float64
	balance = 0
	for _, addr := range addresses {
		cmd := omnijson.OmniGetBalanceCommand{
			Address: addr,
			PropertyID: int32(USDT_PROPERTY_ID),
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
	if len(account) == 0 {
		account = DEFAULT_ACCOUNT
	}

	address, err := client.rpcClient.GetNewAddress(account)
	if err != nil {
		return "", err
	}

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
	return nil
}

//SendFrom ...omni_funded_send
func (client *Client) SendFrom(account, address string, amount float64) error {
	addresses, err := client.rpcClient.GetAddressesByAccount(account)
	if err != nil {
		return err
	}
	for _, addr := range addresses {
		hash, _ := client.rpcClient.OmniFoundedSend(addr, address,int64(USDT_PROPERTY_ID), int64(amount), addr)
		log.Infof("SendFrom USDT, from: %v, to: %v, amount: %v, got hash:%v",addr, address, int64(amount), hash)
	}
	return nil
}


// Move - omni_funded_send
func (client *Client) Move(from, to string, amount float64) (bool, error) {
	hash, err := client.rpcClient.OmniFoundedSend(from, to,int64(USDT_PROPERTY_ID), int64(amount), to)
	if err != nil {
		log.Errorf("Move USDT, from: %v, to: %v, amount: %v, got hash:%v, error: %v", from, to, int64(amount), hash, err)
		return false, err
	}

	log.Infof("Move USDT, from: %v, to: %v, amount: %v, got hash:%v",from, to, int64(amount), hash)
	return true ,nil
}


// ListUnspentMin
func (client *Client) ListUnspentMin(minConf int) ([]btcjson.ListUnspentResult, error) {
	return nil,nil
}
