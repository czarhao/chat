package routes

import (
	"chat/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.GET("/index", controllers.Index)
	route.GET("/room/:name", controllers.Chat)
	return route
}
