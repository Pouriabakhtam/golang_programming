package main

import (
	"github.com/gin-gonic/gin"
	"restapi.ca/db"
	"restapi.ca/routes"
)

func main() {
	db.Initdb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8484") // localhost:8080
}


