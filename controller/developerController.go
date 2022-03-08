package controller

import (
	"api-final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type developerInput struct {
	Name string `json:"name"`
}

// GetAllDeveloper godoc
// @Summary Get all Developer.
// @Description Get a list of Developer.
// @Tags Developer
// @Produce json
// @Success 200 {object} []models.Developer
// @Router /developer [get]
func GetAllDeveloper(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var developers []models.Developer
	db.Find(&developers)

	c.JSON(http.StatusOK, gin.H{"data": developers})
}

// CreateDeveloper godoc
// @Summary Create New Developer.
// @Description Creating a new Developer.
// @Tags Developer
// @Param Body body developerInput true "the body to create a new Developer"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Developer
// @Router /developer [post]
func CreateDeveloper(c *gin.Context) {
	// Validate input
	var input developerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Developer
	developer := models.Developer{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&developer)

	c.JSON(http.StatusOK, gin.H{"data": developer})
}

// GetDeveloperById godoc
// @Summary Get Developer.
// @Description Get an Developer by id.
// @Tags Developer
// @Produce json
// @Param id path string true "Developer id"
// @Success 200 {object} models.Developer
// @Router /developer/{id} [get]
func GetDeveloperById(c *gin.Context) {
	// Get model if exist
	var developer models.Developer

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&developer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": developer})
}

// GetGameByDeveloperId godoc
// @Summary Get Game.
// @Description Get all Game by DeveloperId.
// @Tags Developer
// @Produce json
// @Param id path string true "Developer id"
// @Success 200 {object} []models.Game
// @Router /developer/{id}/game [get]
func GetGameByDeveloperId(c *gin.Context) {
	// Get model if exist
	var games []models.Game

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("developer_id = ?", c.Param("id")).Find(&games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": games})
}

// UpdateDeveloper godoc
// @Summary Update Developer.
// @Description Update Developer by id.
// @Tags Developer
// @Produce json
// @Param id path string true "Developer id"
// @Param Body body developerInput true "the body to update developer"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Developer
// @Router /developer/{id} [patch]
func UpdateDeveloper(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var developer models.Developer
	if err := db.Where("id = ?", c.Param("id")).First(&developer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// Validate Input
	var input developerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Developer
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&developer).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": developer})
}

// DeleteDeveloper godoc
// @Summary Delete one Developer.
// @Description Delete a Developer by id.
// @Tags Developer
// @Produce json
// @Param id path string true "Developer id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /developer/{id} [delete]
func DeleteDeveloper(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var developer models.Developer
	if err := db.Where("id = ?", c.Param("id")).First(&developer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}

	db.Delete(&developer)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
