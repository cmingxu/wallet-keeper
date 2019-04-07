package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cmingxu/wallet-keeper/keeper"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (api *ApiServer) Move(c *gin.Context) {
	value, _ := c.Get(KEEPER_KEY) // sure about the presence of this value
	keeper := value.(keeper.Keeper)

	from, fromFound := c.GetQuery("from")
	to, toFound := c.GetQuery("to")
	amountarg, amountFound := c.GetQuery("amount")

	if !fromFound || !toFound || !amountFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "from account/to account/amount are all mandatory field",
		})
		return
	}

	amount, err := strconv.ParseFloat(amountarg, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	result, err := keeper.Move(from, to, amount)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	if result {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "fail",
		})
	}
}
