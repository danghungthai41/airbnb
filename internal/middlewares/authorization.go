package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeader(r *http.Request) (string, error) {

	return "", nil

}

func RequestedAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeader(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": token})

	}
}
