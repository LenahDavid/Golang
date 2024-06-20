package controllers

import (
	"awesomeProject1/database"
	"awesomeProject1/models"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

// SignUp @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Param input body  controllers.SignUpInput true "Sign up input"
// @Accept  json
// @Produce  json
// @Success 201
// @Router /signup [post]
func SignUp(c *gin.Context) {
	var input struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Username  string `json:"username" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required"`
		Role      string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	var role models.Role
	if err := database.DB.Where("role = ?", input.Role).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}
	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  input.Username,
		Email:     input.Email,
		Password:  string(hashedPassword),
		RoleID:    role.ID,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// SignIn @Summary Sign in a user
// @Description Sign in a user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200
// @param input body controllers.SignInInput true "Sign in input"
// @Router /signin [post]
func SignIn(c *gin.Context) {
	var user models.User
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signed in successfully"})
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// GetUsers @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

// GetUser @Summary Get a user
// @Description Get a user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200
// @Router /user/{id} [get]
// @param id path string true "User ID"
func GetUser(c *gin.Context) {
	var user models.User
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)

}

// UpdateUser @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200
// @Router /user/{id} [put]
// @param id path string true "User ID"
// @Param input body  controllers.SignUpInput true "Sign up input"
func UpdateUser(c *gin.Context) {
	var existingUser models.User
	id := c.Param("id")
	if err := database.DB.First(&existingUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var updateInput SignUpInput
	if err := c.BindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update data"})
		return
	}
	existingUser.FirstName = updateInput.FirstName
	existingUser.LastName = updateInput.LastName
	existingUser.Email = updateInput.Email
	existingUser.Username = updateInput.Username
	existingUser.Password = updateInput.Password

	if err := database.DB.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, existingUser)
}

// DeleteUser @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200
// @Router /user/{id} [delete]
// @param id path string true "User ID"
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

type SignUpInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
