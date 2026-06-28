package routes

import (
	"github.com/gin-gonic/gin"
	middleware "restapi.ca/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/even", GetEvents)
	server.GET("/even/:id", GetSingleEvent)
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authorization)
	authenticated.POST("/even", CreateEvent)
	authenticated.PUT("/even/:id", UpdateEvent)
	authenticated.DELETE("even/:id", DeleteEvent)
	server.POST("/signup/", signup)
	server.POST("/login/", Login)
}
