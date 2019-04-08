package api

import (
	"net/http"

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

	accountInfo, err := keeper.GetAccountInfo(account)
	if err != nil {
		c.JSON(http.StatusNotFound, R("account not found"))
	} else {
		c.JSON(http.StatusOK, R(accountInfo))
	}
}
