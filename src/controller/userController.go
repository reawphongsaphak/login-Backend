package controller

import (
	"net/http"

	"main/src/model"
	"main/src/service"

	"github.com/gin-gonic/gin"
)

func RegisterNewUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.RegisterNewUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"id":      result.InsertedID,
	})
}

func LoginUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := service.LoginUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusCreated, gin.H{
	// 	"token":      token,
	// })

	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user,_ := c.Get("user")
	
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}