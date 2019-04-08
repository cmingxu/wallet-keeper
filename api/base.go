package api

import (
	"github.com/gin-gonic/gin"
)

type Base struct {
}

func R(data interface{}) gin.H {
	return gin.H{
		"message": data,
	}
}
