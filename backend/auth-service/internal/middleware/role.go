package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware проверяет, есть ли у пользователя нужная роль
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "доступ запрещен"})
			c.Abort()
			return
		}
		c.Next()
	}
}
