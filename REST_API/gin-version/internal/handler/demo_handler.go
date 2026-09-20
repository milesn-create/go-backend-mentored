package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.String(200, "Hello, baby)))\n")

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
