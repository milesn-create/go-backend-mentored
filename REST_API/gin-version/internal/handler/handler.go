package handler

import (
	"fmt"
	"net/http"
	"rest-api-gin/internal/models"
	"rest-api-gin/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var UserService *service.UserService

// createUserhandler godoc
// @Summary Создание нового пользователя
// @Tags users
// @Accept json
// @Produce plain
// @Param user body User true "Данные нового пользоавателя"
// @Success 200 {string} string "Пользователь успешно создан"
// @Failure      400  {string}  string  "Некорректный JSON в теле запроса"
// @Failure      409  {string}  string  "Пользователь с таким именем уже существует"
// @Router /users [post]
func CreateUserHandler(c *gin.Context) {
	var u models.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Bad Request\n")
		return
	}
	createdUser, err := UserService.CreateUser(u)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}

	c.String(200, "Данные нового пользователя: ID - %d Имя - %s, возраст - %d\n", createdUser.ID, createdUser.Name, createdUser.Age)
}
func ByeHandler(c *gin.Context) {
	c.String(200, "Bye(((\n")
}
func SumHandler(c *gin.Context) {
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
func IdHandler(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "ID пользователя: %s\n", id)

}
func LoggerMiddleware(c *gin.Context) {
	fmt.Println("Пришел запрос: ", c.Request.Method, c.Request.URL.Path)
}
func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.String(http.StatusUnauthorized, "Unauthorized\n")
		c.Abort()
		return

	}
	c.Next()

}

func HelloHandler(c *gin.Context) {
	c.String(200, "Hello, baby)))\n")

}
