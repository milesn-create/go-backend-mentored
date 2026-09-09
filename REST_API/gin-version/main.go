package main

import (
	_ "rest-api-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gt=0"`
}

// @title User Management API
// @version 1.0
// @description REST API для управления пользователями
// @host localhost:8080
// @BasePath /
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")

}
