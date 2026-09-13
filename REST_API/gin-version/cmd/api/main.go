package main

import (
	"log"
	_ "rest-api-gin/docs"
	"rest-api-gin/internal/database"
	"rest-api-gin/internal/handler"
	"rest-api-gin/internal/repository"
	"rest-api-gin/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User Management API
// @version 1.0
// @description REST API для управления пользователями
// @host localhost:8080
// @BasePath /
func main() {
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Successfully connected to database!")
	repo := repository.NewUserRepository()
	handler.UserService = service.NewUserService(repo)
	router := gin.Default()
	router.Use(handler.LoggerMiddleware)
	router.GET("/hello", handler.HelloHandler)
	router.GET("/bye", handler.ByeHandler)
	router.GET("/sum", handler.SumHandler)
	router.GET("/users/:id", handler.IdHandler)
	router.POST("/users", handler.AuthMiddleware, handler.CreateUserHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")

}
