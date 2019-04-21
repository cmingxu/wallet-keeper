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
		c.JSON(http.StatusBadRequest, R("no account name specified"))
		return
	}

	// account validation
	if !accountValid(account) {
		c.JSON(http.StatusBadRequest, R("account not valid, ^([a-z|A-Z|0-9])+$"))
		return
	}

	if _, err := keeper.GetAccountInfo(account, 0); err == nil {
		c.JSON(http.StatusBadRequest, R("account exists"))
		return
	}

	created, err := keeper.CreateAccount(account)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, R(fmt.Sprint(err)))
	} else {
		c.JSON(http.StatusOK, R(created))
	}
}

func accountValid(account string) bool {
	return ACCOUNT_REGEXP.MatchString(account)
}
