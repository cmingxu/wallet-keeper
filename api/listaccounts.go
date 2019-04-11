package api

import (
	"fmt"
	"net/http"

	"github.com/cmingxu/wallet-keeper/keeper"
	"github.com/cmingxu/wallet-keeper/keeper/btc"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ListAccounts func
func (api *ApiServer) ListAccounts(c *gin.Context) {
	value, _ := c.Get(KEEPER_KEY) // sure about the presence of this value
	keeper := value.(keeper.Keeper)

	accounts, err := keeper.ListAccountsMinConf(btc.DEFAULT_CONFIRMATION)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, R(fmt.Sprint(err)))
	} else {
		c.JSON(http.StatusOK, R(accounts))
	}
}
