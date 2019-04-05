package api

import (
	"strings"

	"github.com/cmingxu/wallet-keeper/keeper"
	"github.com/cmingxu/wallet-keeper/keeper/btc"

	"github.com/gin-gonic/gin"
)

const KEEPER_KEY = "keeper"
const COIN_TYPE_HEADER = "CoinType"

// http api list
var METHODS_SUPPORTED = map[string]string{
	// misc
	"/ping":   "check if api service valid and backend bitcoin service healthy",
	"/health": "check system status",
	"/help":   "display this message",

	// useful APIs here
	"/getblockcount": "return height of the blockchain",
	"/getaddress":    "return address of specified account or default",
	"/getnewaddress": "return a new address of specified account or default",
	"/getbalance":    "sum balances of all accounts",
	"/listaccounts":  "list accounts with amount, minconf is 6",
	//"/getaddress_with_balances": "all addresses together with balances",
}

type ApiServer struct {
	httpListenAddr string
	btcKeeper      keeper.Keeper
	usdtKeeper     keeper.Keeper
}

//TODO valid host is valid
func (api *ApiServer) InitBtcClient(host, user, pass string) (err error) {
	api.btcKeeper, err = btc.NewClient(host, user, pass)
	return err
}

func NewApiServer(addr string) (*ApiServer, error) {
	return &ApiServer{
		httpListenAddr: addr,
	}, nil
}

func (api *ApiServer) HttpListen() error {
	r := gin.Default()

	// with midlleware determine which currency is active
	// within this very session, should be either `btc` or  `usdt`.
	// if none of these present, abort this request and return caller
	// with bad request instantly.
	r.Use(func(c *gin.Context) {
		coin_type := strings.ToLower(c.Request.Header.Get(COIN_TYPE_HEADER))
		switch coin_type {
		case "btc":
			c.Set(KEEPER_KEY, api.btcKeeper)
			break
		case "usdt":
			c.Set(KEEPER_KEY, api.btcKeeper)
			break
		default:
			c.JSON(400, gin.H{"message": "no coin type specified, should be btc or usdt"})
		}
	})

	// APIs related to wallet
	r.GET("/getblockcount", api.GetBlockCount)
	r.GET("/getaddress", api.GetAddress)
	r.GET("/getaddresses", api.GetAddresses)
	r.GET("/getnewaddress", api.GetNewAddress)
	r.GET("/listaccounts", api.ListAccounts)

	// misc API
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	r.GET("/help", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"methods": METHODS_SUPPORTED,
		})
	})

	return r.Run(api.httpListenAddr)
}
