package api

import (
	"fmt"

	"github.com/cmingxu/wallet-keeper/keeper/eth"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (api *ApiServer) Test(c *gin.Context) {
	log.Println("xxxxxxxxxxxx")
	c.JSON(http.StatusOK, R("success"))
}
