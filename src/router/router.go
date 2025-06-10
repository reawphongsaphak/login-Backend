package router

import (
	"main/src/controller"
	"main/src/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1") 
	{
		api.POST("/register", controller.RegisterNewUser)
		api.POST("/login", controller.LoginUser)
		api.GET("/validate", middleware.RequireAuth, controller.Validate)
	}	
}