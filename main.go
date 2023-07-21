package main

import (
	"docker-compose/database"
	"docker-compose/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	database.MigrateDatabase()

	router.GET("/",routes.Home)
	router.POST("/",routes.Display)

	router.Run()
}