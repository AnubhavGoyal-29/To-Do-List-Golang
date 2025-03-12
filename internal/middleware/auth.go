package middleware

import (
	"net/http"
	"ToDoListGolang/internal/database"
	"context"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get logged-in user from Redis
		userID, err := database.RedisClient.Get(ctx, "logged_in_user").Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
			c.Abort()
			return
		}

		// Store user ID in context for later use
		c.Set("userID", userID)

		c.Next()
	}
}

