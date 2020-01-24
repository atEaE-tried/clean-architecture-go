package infrastructure

import (
	gin "github.com/gin-gonic/gin"

	"github.com/atEaE-tried/clean-architecture-go/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserControlller(NewSQLHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })

	Router = router
}
