package api

import (
	"github.com/cmingxu/wallet-keeper/cryptocurrency/btc"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ApiServer struct {
	httpListenAddr string
}

func (api *ApiServer) InitBtcClient(btcAddr string) error {
	return nil
}

func NewApiServer(addr string) (*ApiServer, error) {
	return &ApiServer{
		httpListenAddr: addr,
	}, nil
}

func (api *ApiServer) HttpListen() error {
	cli, err := btc.NewClient()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		if err != nil {
			log.Error(err)
		}
		log.Println(cli)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r.Run(api.httpListenAddr)
}
