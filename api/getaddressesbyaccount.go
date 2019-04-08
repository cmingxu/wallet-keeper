package api

import (
	"fmt"
	"net/http"

	"github.com/cmingxu/wallet-keeper/keeper"
	"github.com/cmingxu/wallet-keeper/keeper/btc"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (api *ApiServer) GetAddressesByAccount(c *gin.Context) {
	value, _ := c.Get(KEEPER_KEY) // sure about the presence of this value
	keeper := value.(keeper.Keeper)

	account, found := c.GetQuery("account")
	if !found {
		account = btc.DEFAULT_ACCOUNT
	}

	addresses, err := keeper.GetAddressesByAccount(account)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, R(fmt.Sprint(err)))
	} else {
		c.JSON(http.StatusOK, R(addresses))
	}
}
