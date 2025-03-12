package taskcontroller

import (
	"net/http"
	"strconv"
	"ToDoListGolang/internal/database"
	"ToDoListGolang/internal/models"

	"github.com/gin-gonic/gin"
)

// CreateTask - POST /tasks
func CreateTask(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("userID"))

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign logged-in user ID to the task
	task.UserID = uint(userID)

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetTasks - GET /tasks (Only gets tasks for logged-in user)
func GetTasks(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("userID"))

	var tasks []models.Task
	query := database.DB.Where("user_id = ?", userID)

	// Filtering by Status (e.g., /tasks?status=Completed)
	status := c.Query("status")
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Pagination (default: page 1, limit 10)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	query.Limit(limit).Offset(offset).Find(&tasks)
	var taskResponses []models.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, task.Serialize())
	}

	c.JSON(http.StatusOK, taskResponses)
}

// GetTask - GET /tasks/:id (Ensures user can only access their own tasks)
func GetTask(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("userID"))
	id := c.Param("id")

	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task.Serialize())
}

// UpdateTask - PUT /tasks/:id (Ensures user can only update their own tasks)
func UpdateTask(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("userID"))
	id := c.Param("id")

	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// DeleteTask - DELETE /tasks/:id (Ensures user can only delete their own tasks)
func DeleteTask(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("userID"))
	id := c.Param("id")

	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

