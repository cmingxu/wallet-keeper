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
		Host:         "192.168.0.101:8332",
		User:         "foo",
		Pass:         "qDDZdeQ5vw9XXFeVnXT4PZ--tGN2xNjjR4nrtyszZx0=",
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

	err = client.Ping()
	log.Println(err)

	blockChainInfo, err := client.GetBlockChainInfo()
	if err != nil {
		log.Error(err)
	}
	log.Println(blockChainInfo)

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
		log.Println("GetBestBlock", err)
	}
	log.Println(hash)
	log.Println(height)

	blockCount, err = client.GetBlockCount()
	if err != nil {
		log.Error("blocCount: ", err)
	}
	log.Printf("Block count: %d", blockCount)

	v, err := client.Version()
	if err != nil {
		log.Error("Version :", err)
	}
	log.Printf("v: %+v", v)

	hash, err = client.GetBlockHash(1)
	if err != nil {
		log.Error("GetBlockHash", err)
	}

	log.Println(hash)

	verboseHeader, err := client.GetBlockHeaderVerbose(hash)
	if err != nil {
		log.Error("verboseHeader", err)
	}
	log.Println(verboseHeader)

	rawResponse, err := client.RawRequest("getnetworkinfo", nil)
	if err != nil {
		log.Error("RawRequest", err)
	}
	log.Println(rawResponse)
	return client, nil
}
