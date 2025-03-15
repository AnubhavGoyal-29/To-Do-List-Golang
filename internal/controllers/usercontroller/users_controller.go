package usercontroller

import (
	"net/http"
	"ToDoListGolang/internal/database"
	"ToDoListGolang/internal/models"
	"ToDoListGolang/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"context"
)

var ctx = context.Background()

func LogoutUser(c *gin.Context) {
	userId, _ := c.Get("userID")
	result := database.DB.Model(&models.UserToken{}).
		Where("user_id = ? AND is_valid = 1", userId).
		Update("is_valid", false)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or already logged out token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func LoginUser(c *gin.Context) {
	var user models.User
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Find user by email
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	userToken := models.UserToken{UserID: user.ID, Token: token, IsValid: true}
	database.DB.Create(&userToken)
	// Return token
	c.JSON(http.StatusOK, gin.H{"token": token})

}

func CreateUser(c *gin.Context) {
	var user models.User

	// Bind JSON input
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Create user in the database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	userToken := models.UserToken{UserID: user.ID, Token: token, IsValid: true}
	database.DB.Create(&userToken)
	// Return the user object and login success message
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered and logged in successfully",
		"user": gin.H{
			"token": token,
		},
	})
}


// GetUser - GET /users/:id
func GetUser(c *gin.Context) {
	id, _ := c.Get("userID")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	response := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser - PUT /users/:id
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser - DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

