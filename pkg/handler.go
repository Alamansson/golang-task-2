package pkg

import (
	"github.com/gin-gonic/gin"
)


func InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		account := api.Group("/account")
		{
			account.GET("/", getAccount)
			account.DELETE("/:id", deleteAccount)
			account.POST("/", createAccount)
			account.PUT("/:id", updateAccount)
		}
	}
	return router
}
