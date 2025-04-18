package api

import (
	"github.com/gin-gonic/gin"
	"os"
)

var router = gin.Default()

func Run() error {
	v1 := router.Group("/v1")
	BySwiftCode(v1)
	ByCountryCode(v1)
	PostSwiftCode(v1)
	DeleteBySwiftCode(v1)

	err := router.Run(os.Getenv("httpListenOn"))
	return err
}

func abortWithJSON(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, gin.H{"message": message})
}
