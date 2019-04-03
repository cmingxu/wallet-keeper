package btc

import (
	btc "github.com/btcsuite/btcd/rpcclient"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	btcClient *btc.Client
}

func NewClient() (*btc.Client, error) {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &btc.ConnConfig{
		Host:         "localhost:8332",
		User:         "xcm",
		Pass:         "foobar",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := btc.New(connCfg, nil)
	if err != nil {
		log.Error(err)
	}
	defer client.Shutdown()

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Error(err)
	}
	log.Printf("Block count: %d", blockCount)

	accounts, err := client.ListAccounts()
	log.Println(len(accounts))
	for a := range accounts {
		log.Println(a)
	}

	hash, height, err := client.GetBestBlock()
	if err != nil {
		log.Println(err)
	}
	log.Println(hash)
	log.Println(height)

	amount, err := client.GetBalance("")
	if err != nil {
		log.Println(err)
	}
	log.Println(amount)

	blockCount, err = client.GetBlockCount()
	if err != nil {
		log.Error(err)
	}
	log.Printf("Block count: %d", blockCount)

	v, err := client.Version()
	if err != nil {
		log.Error(err)
	}
	log.Printf("v: %+v", v)
	return client, nil
}
