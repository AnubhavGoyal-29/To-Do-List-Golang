package routes

import (
	"ToDoListGolang/internal/controllers/taskcontroller"
	"ToDoListGolang/internal/controllers/usercontroller"
	"ToDoListGolang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	// User Routes

	r := gin.Default()
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/create", usercontroller.CreateUser) // No middleware
		userRoutes.POST("/login", usercontroller.LoginUser)   // No middleware

		// Middleware applied to all routes except login and create
		userRoutes.Use(middleware.AuthMiddleware())

		userRoutes.GET("/profile", usercontroller.GetUser)
		userRoutes.PUT("/update/:id", usercontroller.UpdateUser)
		userRoutes.DELETE("/delete/:id", usercontroller.DeleteUser)
		userRoutes.POST("/logout", usercontroller.LogoutUser)
	}

	// Task Routes (Protected)
	taskRoutes := r.Group("/tasks").Use(middleware.AuthMiddleware())
	{
		taskRoutes.GET("/", taskcontroller.GetTasks)
		taskRoutes.GET("/:id", taskcontroller.GetTask)
		taskRoutes.POST("/create", taskcontroller.CreateTask)
		taskRoutes.PUT("/update/:id", taskcontroller.UpdateTask)
		taskRoutes.DELETE("/delete/:id", taskcontroller.DeleteTask)
	}

	return r
}

