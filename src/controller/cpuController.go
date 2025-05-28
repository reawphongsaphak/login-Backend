package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"main/src/service"
)

func GetAllCPU (c *gin.Context) {
	cpus, err := service.GetAllCPU()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cpus)
}