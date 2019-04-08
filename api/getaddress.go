package api

import (
	"fmt"
	"net/http"

	"github.com/cmingxu/wallet-keeper/keeper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (api *ApiServer) GetAddress(c *gin.Context) {
	value, _ := c.Get(KEEPER_KEY) // sure about the presence of this value
	keeper := value.(keeper.Keeper)

	address, err := keeper.GetAddress("")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": address,
		})
	}
}
