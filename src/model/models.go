package models

import (

)

// type User struct {
// 	name string
// 	password string
// }

type CPU struct {
	ProductID 	string 	`json:"cpu_id"`
	Title		string 	`json:"title"`
	ImgURL		string 	`json:"imgUrl"`
	Price		int 	`json:"price"`
	Socket		string	`json:"Socket"`
	Brand		string	`json:"brand"`
	Quantity	int		`json:"quantity"`
}

type User struct {
	UserName	string `json:"username" binding:"required"`
	Password	string `json:"password" binding:"required"`
}