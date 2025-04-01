package midleware

import (
	"gin-starter/core/domain"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	user := &domain.User{}
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" || len(authHeader) <= len("Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorize",
		})
		return
	}
	token := authHeader[len("Bearer "):]

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	user.ID = uint(claims["auth"].(float64))

	c.Set("user", user)
	c.Next()
}
