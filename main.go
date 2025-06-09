package main

import (
	"log"
	
	"main/src/router"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

	r := gin.Default()
	r.Use(cors.Default())

	router.SetupRoutes(r)
	r.Run(":8080")
}