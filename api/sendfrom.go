package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cmingxu/wallet-keeper/keeper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (api *ApiServer) SendFrom(c *gin.Context) {
	value, _ := c.Get(KEEPER_KEY) // sure about the presence of this value
	keeper := value.(keeper.Keeper)

	from, fromFound := c.GetQuery("from")
	address, addrFound := c.GetQuery("address")
	amountArg, amountFound := c.GetQuery("amount")
	if !fromFound || !addrFound || !amountFound {
		c.JSON(http.StatusBadRequest, R("from, address, amount are mandatory fields"))
		return
	}

	amount, err := strconv.ParseFloat(amountArg, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, R(fmt.Sprint(err)))
		return
	}

	// simple validation
	if amount <= 0 {
		c.JSON(http.StatusBadRequest, R("amount should at least greater than 0"))
		return
	}

	hash, err := keeper.SendFrom(from, address, amount)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, R(fmt.Sprint(err)))
	} else {
		c.JSON(http.StatusOK, R(gin.H{"address": address, "txid": hash}))
	}
}
