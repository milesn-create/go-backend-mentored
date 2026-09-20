package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
