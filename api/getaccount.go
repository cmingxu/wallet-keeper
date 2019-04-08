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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no account name specified",
		})
		return
	}

	accountInfo, err := keeper.GetAccountInfo(account)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "account not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": accountInfo,
		})
	}
}
