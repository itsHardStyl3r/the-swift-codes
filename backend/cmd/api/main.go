package api

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() error {
	v1 := router.Group("/v1")
	bySwiftCode(v1)
	byCountryCode(v1)
	postSwiftCode(v1)
	deleteBySwiftCode(v1)

	err := router.Run("127.0.0.1:8080")
	return err
}
