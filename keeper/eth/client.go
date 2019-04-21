package eth

import (
	"os"
	"path/filepath"

	"github.com/cmingxu/wallet-keeper/keeper"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

const PASSWORD = "password"

type Client struct {
	ethClient *ethclient.Client
	ks        *keystore.KeyStore
	l         *log.Logger
}

func NewClient(host, keyStorePath, logDir string) (*Client, error) {
	ethClient, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		return nil, err
	}

	client := &Client{ethClient: ethClient}
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

// Create Key store
func (client *Client) createKeyStore(keyStorePath string) error {
	ks := keystore.NewKeyStore(keyStorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(PASSWORD)
	if err != nil {
		return err
	}

	cient.ks = ks
}

// GetBlockCount
func (client *Client) GetBlockCount() (int64, error) {
	header, err := client.ethClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}

	return header.Number.Int64(), nil
}

// GetAddress - default address
func (client *Client) GetAddress(account string) (string, error) {
	return "", err
}

// Create Account
// Returns customized account info
func (client *Client) CreateAccount(account string, minConf int) (keeper.Account, error) {
	return nil, nil
}

// GetAccountInfo
func (client *Client) GetAccountInfo(account string) (keeper.Account, error) {
	return keeper.Account{}, nil
}

// TODO
// GetNewAddress does map to `getnewaddress` rpc call now
// rpcclient doesn't have such golang wrapper func.
func (client *Client) GetNewAddress(account string) (string, error) {
	return "", nil
}

// GetAddressesByAccount
func (client *Client) GetAddressesByAccount(account string) ([]string, error) {
	return []string{}, nil
}

// ListAccountsMinConf
func (client *Client) ListAccountsMinConf(conf int) (map[string]float64, error) {
	return make(map[string]float64), nil
}

// SendToAddress
func (client *Client) SendToAddress(address string, amount float64) error {
	return nil
}

// TODO check validity of account and have sufficent balance
func (client *Client) SendFrom(account, address string, amount float64) error {
	return nil
}

// Move
func (client *Client) Move(from, to string, amount float64) (bool, error) {
	return true, nil
}
