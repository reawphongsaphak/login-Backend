package models

import (

)

type User struct {
	Email 		string `json:"email" binding:"required"`
	UserName	string `json:"username"`
	Password	string `json:"password" binding:"required"`
}