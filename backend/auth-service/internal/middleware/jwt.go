package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "нет заголовка Authorization"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "недействительный токен"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID, _ := claims["user_id"].(string)
		role, _ := claims["role"].(string)

		c.Set("user_id", userID)
		c.Set("role", role)

		c.Next()
	}
}

// GetSecret возвращает секретный ключ для подписи JWT
func GetSecret() []byte {
	// Можно также использовать os.Getenv или config-файл
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// На dev-среде можно использовать дефолт
		secret = "default_secret_key"
	}
	return []byte(secret)
}
