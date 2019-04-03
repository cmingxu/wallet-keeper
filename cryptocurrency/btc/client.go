package btc

import (
	"log"

	btc "github.com/btcsuite/btcd/rpcclient"
	"github.com/pkg/errors"
)

type Client struct {
	btcClient *btc.Client
}

func NewClient() (*btc.Client, error) {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &btc.ConnConfig{
		Host:         "localhost:8332",
		User:         "test",
		Pass:         "test",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := btc.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
}
