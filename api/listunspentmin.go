package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cmingxu/wallet-keeper/keeper"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (api *ApiServer) ListUnspentMin(c *gin.Context) {
	value, _ := c.Get(KEEPER_KEY) // sure about the presence of this value
	keeper := value.(keeper.Keeper)

	confarg, found := c.GetQuery("minconf")
	if !found {
		confarg = "0"
	}

	conf, err := strconv.ParseUint(confarg, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, R(fmt.Sprint(err)))
		return
	}

	result, err := keeper.ListUnspentMin(int(conf))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, R(fmt.Sprint(err)))
	} else {
		c.JSON(http.StatusOK, R(result))
	}
}
