package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gt=0"`
}

func main() {
	repo := NewUserRepository()
	userService = NewUserService(repo)
	router := gin.Default()
	router.Use(loggerMiddleware)
	router.GET("/hello", helloHandler)
	router.GET("/bye", byeHandler)
	router.GET("/sum", sumHandler)
	router.GET("/users/:id", idHandler)
	router.POST("/users", authMiddleware, createUserHandler)

	router.Run(":8080")

}
