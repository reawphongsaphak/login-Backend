package main

import (
	"main/src/router"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	router.SetupRoutes(r)
	r.Run(":8080")
}