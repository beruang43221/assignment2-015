package routers

import (
	"github.com/beruang43221/assignment2-015/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/orders", controller.GetAllData)

	router.POST("/order", controller.CreateData)

	router.PUT("/order/:id", controller.UpdateData)

	router.DELETE("/order/:id", controller.DeleteData)

	return router
}