package api

import (
	"github.com/gin-gonic/gin"
	"os"
)

var router = gin.Default()

func Run() error {
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Next()
	})
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
