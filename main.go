package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ToDoListGolang/internal/database"
	"ToDoListGolang/internal/models"
)

func main() {
	database.InitDB()
	database.InitRedis()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Task{})
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to To-Do List API!"})
	})

	r.Run(":8080")
}

