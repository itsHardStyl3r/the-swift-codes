package api

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() error {
	v1 := router.Group("/v1")
	BySwiftCode(v1)
	ByCountryCode(v1)
	PostSwiftCode(v1)
	DeleteBySwiftCode(v1)

	err := router.Run("127.0.0.1:8080")
	return err
}

func abortWithJSON(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, gin.H{"message": message})
}
