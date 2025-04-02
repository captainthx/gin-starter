package midleware

import (
	"fmt"
	"gin-starter/core/domain"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	user := &domain.User{}
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorize",
		})
		return
	}

	token := authHeader
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = authHeader[len("Bearer "):]
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !parsedToken.Valid {
		fmt.Println("error = ", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token expired",
		})
		return
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Token",
		})
		return
	}

	user.ID = uint(claims["auth"].(float64))
	c.Set("user", user)
	c.Next()
}
