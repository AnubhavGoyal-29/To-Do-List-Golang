package routes

import (
	"ToDoListGolang/internal/controllers/taskcontroller"
	"ToDoListGolang/internal/controllers/usercontroller"
	"ToDoListGolang/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
    r := gin.Default()

    userRoutes := r.Group("/users")
    {
        userRoutes.POST("/create", usercontroller.CreateUser)
        userRoutes.POST("/login", usercontroller.LoginUser)
    }

    protectedUserRoutes := r.Group("/users").Use(middleware.AuthMiddleware())
    {
        protectedUserRoutes.GET("/profile", usercontroller.GetUser)
        protectedUserRoutes.PUT("/update/:id", usercontroller.UpdateUser)
        protectedUserRoutes.DELETE("/delete/:id", usercontroller.DeleteUser)
        protectedUserRoutes.POST("/logout", usercontroller.LogoutUser)
    }

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

