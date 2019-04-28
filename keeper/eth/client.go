package eth

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"os"
	"path/filepath"
	"sync"

	"github.com/cmingxu/wallet-keeper/keeper"
	"github.com/cmingxu/wallet-keeper/notifier"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
)

const PASSWORD = "password"

var ErrNotValidAccountFile = errors.New("not valid account file")
var ErrNotDirectory = errors.New("not valid directory")

type Client struct {
	l *log.Logger

	// Checkout https://github.com/ethereum/go-ethereum/blob/master/rpc/client.go
	// for more details.
	ethRpcClient *rpc.Client

	// fs directory where to store wallet
	walletDir string

	// keystore
	store *keystore.KeyStore

	// account/address map lock, since ethereum doesn't support account
	// we should have our own account/address map internally.

	// only with this map we can provide services for the upstream services.
	accountFilePath    string
	accountAddressMap  map[string]string
	accountAddressLock sync.Mutex

	noti *notifier.Notifier
}

func NewClient(host, walletDir, accountFilePath, logDir string) (*Client, error) {
	client := &Client{
		walletDir:          walletDir,
		accountFilePath:    accountFilePath,
		accountAddressMap:  make(map[string]string),
		accountAddressLock: sync.Mutex{},
		noti:               notifier.New(),
	}

	// accountAddressMap initialization
	stat, err := os.Stat(client.accountFilePath)
	if err != nil {
		return nil, err
	}

	if !stat.Mode().IsRegular() {
		return nil, ErrNotValidAccountFile
	}

	err = client.loadAccountMap()
	if err != nil {
		return nil, err
	}

	// keystore initialization
	stat, err = os.Stat(walletDir)
	if err != nil {
		return nil, ErrNotDirectory
	}

	if !stat.IsDir() {
		return nil, ErrNotDirectory
	}
	client.store = keystore.NewKeyStore(walletDir, keystore.StandardScryptN, keystore.StandardScryptP)

	// rpcClient initialization
	client.ethRpcClient, err = rpc.Dial(host)
	if err != nil {
		return nil, err
	}

	// log initialization
	logPath := filepath.Join(logDir, "eth.log")
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return nil, err
	}

	client.l = &log.Logger{
		Out:       logFile,
		Formatter: new(log.JSONFormatter),
	}

	return client, nil
}

// Ping
func (client *Client) Ping() error {
	return nil
}

// GetBlockCount
func (client *Client) GetBlockCount() (int64, error) {
	var num string
	err := client.ethRpcClient.CallContext(context.Background(), &num, "eth_blockNumber")
	if err != nil {
		return 0, err
	}

	big, err := hexutil.DecodeBig(num)
	if err != nil {
		return 0, err
	}

	return big.Int64(), nil
}

// GetAddress - default address
func (client *Client) GetAddress(account string) (string, error) {
	address, ok := client.accountAddressMap[account]
	if !ok {
		return "", keeper.ErrAccountNotFound
	}

	return address, nil
}

// Create Account
func (client *Client) CreateAccount(account string) (keeper.Account, error) {
	address, _ := client.GetAddress(account)
	log.Println(address)
	if len(address) > 0 {
		return keeper.Account{}, keeper.ErrAccountExists
	}

	acc, err := client.store.NewAccount(PASSWORD)
	if err != nil {
		return keeper.Account{}, err
	}

	client.accountAddressLock.Lock()
	client.accountAddressMap[account] = acc.Address.Hex()
	client.accountAddressLock.Unlock()

	err = client.persistAccountMap()
	if err != nil {
		return keeper.Account{}, err
	}

	return keeper.Account{
		Account: account,
		Balance: 0,
		Addresses: []string{
			acc.Address.Hex(),
		},
	}, nil
}

// GetAccountInfo
func (client *Client) GetAccountInfo(account string, minConf int) (keeper.Account, error) {
	address, found := client.accountAddressMap[account]
	if !found {
		return keeper.Account{}, keeper.ErrAccountNotFound
	}

	var balance hexutil.Big
	err := client.ethRpcClient.CallContext(context.Background(), &balance, "eth_getBalance", common.HexToAddress(address), "latest")
	if err != nil {
		return keeper.Account{}, err
	}

	float64Value, _ := weiToEther(balance.ToInt()).Float64()
	return keeper.Account{
		Account:   account,
		Balance:   float64Value,
		Addresses: []string{address},
	}, nil
}

func (client *Client) GetNewAddress(account string) (string, error) {
	return "", keeper.ErrNotSupport
}

// GetAddressesByAccount
func (client *Client) GetAddressesByAccount(account string) ([]string, error) {
	address, ok := client.accountAddressMap[account]
	if !ok {
		return []string{}, keeper.ErrAccountNotFound
	}

	return []string{address}, nil
}

// ListAccountsMinConf
func (client *Client) ListAccountsMinConf(conf int) (map[string]float64, error) {
	accounts := make(map[string]float64, len(client.accountAddressMap))
	for name, address := range client.accountAddressMap {
		var balance hexutil.Big
		err := client.ethRpcClient.CallContext(context.Background(), &balance, "eth_getBalance", common.HexToAddress(address), "latest")
		if err != nil {
			client.l.Errorf("[ListAccountsMinConf] %s", err)

			accounts[name] = 0
		} else {
			float64Value, _ := weiToEther(balance.ToInt()).Float64()
			accounts[name] = float64Value
		}
	}

	return accounts, nil
}

// SendToAddress
func (client *Client) SendToAddress(address string, amount float64) error {
	return keeper.ErrNotSupport
}

// TODO check validity of account and have sufficent balance
func (client *Client) SendFrom(account, address string, amount float64) error {
	return nil
}

// ListUnspentMin
func (client *Client) ListUnspentMin(minConf int) ([]btcjson.ListUnspentResult, error) {
	return []btcjson.ListUnspentResult{}, keeper.ErrNotSupport
}

// Move
func (client *Client) Move(from, to string, amount float64) (bool, error) {
	return true, nil
}

// persistAccountMap write `accountAddressMap` into file `client.accountAddressMap`,
// `accountAddressMap` will persist into file with json format,
//
// Error - return if `client.accountFilePath` not found or write permission not right.
func (client *Client) persistAccountMap() error {
	stat, err := os.Stat(client.accountFilePath)
	if err != nil && os.IsNotExist(err) {
		return err
	}

	if !stat.Mode().IsRegular() {
		return ErrNotValidAccountFile
	}

	file, err := os.OpenFile(client.accountFilePath, os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(client.accountAddressMap)
}

// loadAccountMap from filesystem.
func (client *Client) loadAccountMap() error {
	client.accountAddressMap = make(map[string]string)
	file, err := os.Open(client.accountFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&client.accountAddressMap)
	if err != nil {
		return err
	}

	return nil
}

func weiToEther(wei *big.Int) *big.Float {
	result := new(big.Float)
	weiEther := new(big.Float).SetFloat64(float64(1 / params.Ether))
	weiInFloat := new(big.Float).SetFloat64(float64(wei.Int64()))
	result.Mul(weiInFloat, weiEther)
	return result
}
