package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userService *UserService

func createUserHandler(c *gin.Context) {
	var u User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Bad Request\n")
		return
	}
	createdUser, err := userService.CreateUser(u)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}

	c.String(200, "Данные нового пользователя: ID - %d Имя - %s, возраст - %d\n", createdUser.ID, createdUser.Name, createdUser.Age)
}
func byeHandler(c *gin.Context) {
	c.String(200, "Bye(((\n")
}
func sumHandler(c *gin.Context) {
	a, err := strconv.Atoi(c.Query("a"))
	if err != nil {
		c.String(http.StatusBadRequest, "a - not an integer\n")
		return
	}
	b, err := strconv.Atoi(c.Query("b"))
	if err != nil {
		c.String(http.StatusBadRequest, "b - not an integer\n")
		return
	}
	sum := a + b
	c.String(200, "Сумма = %d\n", sum)
}
func idHandler(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "ID пользователя: %s\n", id)

}
func loggerMiddleware(c *gin.Context) {
	fmt.Println("Пришел запрос: ", c.Request.Method, c.Request.URL.Path)
}
func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.String(http.StatusUnauthorized, "Unauthorized\n")
		c.Abort()
		return

	}
	c.Next()

}

func helloHandler(c *gin.Context) {
	c.String(200, "Hello, baby)))\n")

}
