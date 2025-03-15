package middleware

import (
	"net/http"
	"ToDoListGolang/internal/utils"
	"ToDoListGolang/internal/database"
	"ToDoListGolang/internal/models"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT and extracts user_id
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Validate JWT
		userID, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		var userToken models.UserToken
		if err := database.DB.Where("user_id = ? AND token = ? AND is_valid = 1", userID, tokenString).First(&userToken).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}

		// Store user_id in context
		c.Set("userID", userID)
		c.Next()
	}
}

