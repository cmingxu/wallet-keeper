package api

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/cmingxu/wallet-keeper/keeper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var ACCOUNT_REGEXP = regexp.MustCompile("^([a-z|A-Z|0-9])+$")

// Create func
func (api *ApiServer) CreateAccount(c *gin.Context) {
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

	// account validation
	if !accountValid(account) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "account not valid, ^([a-z|A-Z|0-9])+$",
		})
		return
	}

	if _, err := keeper.GetAccountInfo(account); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "account exists",
		})
		return
	}

	created, err := keeper.CreateAccount(account)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": created,
		})
	}
}

func accountValid(account string) bool {
	return ACCOUNT_REGEXP.MatchString(account)
}
