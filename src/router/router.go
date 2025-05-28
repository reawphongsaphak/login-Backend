package router

import (
	"main/src/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1") 
	{
		api.GET("/cpus", controller.GetAllCPU)
	}	
}