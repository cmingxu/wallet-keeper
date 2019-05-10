package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cmingxu/wallet-keeper/keeper"

	"github.com/gin-gonic/gin"
)

// Create func
func (api *ApiServer) GetAccountInfo(c *gin.Context) {
	value, _ := c.Get(KEEPER_KEY) // sure about the presence of this value
	keeper := value.(keeper.Keeper)

	// retrive account from query
	account, found := c.GetQuery("account")
	if !found {
		c.JSON(http.StatusBadRequest, R("no account name specified"))
		return
	}

	confarg, found := c.GetQuery("minconf")
	if !found {
		confarg = "6"
	}

	conf, err := strconv.ParseUint(confarg, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, R(fmt.Sprint(err)))
		return
	}

	accountInfo, err := keeper.GetAccountInfo(account, int(conf))
	if err != nil {
		c.JSON(http.StatusNotFound, R("account not found"))
	} else {
		c.JSON(http.StatusOK, R(accountInfo))
	}
}
