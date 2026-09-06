package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gt=0"`
}

func helloHandler(c *gin.Context) {
	c.String(200, "Hello, baby)))\n")
}
func createUserHandler(c *gin.Context) {
	var u User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Bad Request\n")
		return
	}

	c.String(200, "Данные нового пользователя: Имя - %s, возраст - %d\n", u.Name, u.Age)
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
func main() {
	router := gin.Default()
	router.GET("/hello", helloHandler)
	router.GET("/bye", byeHandler)
	router.GET("/sum", sumHandler)
	router.GET("/users/:id", idHandler)
	router.POST("/users", createUserHandler)
	router.Run(":8080")

}
