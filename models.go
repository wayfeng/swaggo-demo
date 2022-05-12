package main

type User struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Age   int    `json:"age" example:"42"`
	Email string `json:"email" binding:"required" example:"john.doe@example.com"`
}
