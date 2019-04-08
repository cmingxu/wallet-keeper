package api

import (
	"fmt"
	"net/http"

	"github.com/cmingxu/wallet-keeper/keeper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ListAccounts func
func (api *ApiServer) ListAccounts(c *gin.Context) {
	value, _ := c.Get(KEEPER_KEY) // sure about the presence of this value
	keeper := value.(keeper.Keeper)
	minConf := 6 // default min conf

	accounts, err := keeper.ListAccountsMinConf(minConf)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": accounts,
		})
	}
}
