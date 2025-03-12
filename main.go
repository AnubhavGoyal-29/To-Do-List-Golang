package main

import (
	"fmt"
	"log"
	"ToDoListGolang/internal/database"
	"ToDoListGolang/internal/models"
	"ToDoListGolang/routes"
)

func main() {
	// Initialize MySQL and Redis
	database.InitDB()
	database.InitRedis()

	// Run Auto Migrations
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Task{})
	fmt.Println("Database migrated successfully")

	r := routes.RegisterRoutes()
	log.Fatal(r.Run(":8080"))
}

