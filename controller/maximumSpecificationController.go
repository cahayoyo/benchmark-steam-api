package controller

import (
	"api-final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type maximumSpecificationInput struct {
	OS        string `json:"os"`
	Processor string `json:"processor"`
	Memory    string `json:"memory"`
	Graphics  string `json:"graphics"`
	DirectX   string `json:"directx"`
	Network   string `json:"network"`
	Storage   string `json:"storage"`
	SoundCard string `json:"soundcard"`
	GameID    uint   `json:"game_id"`
}

// GetAllMaximumSpecification godoc
// @Summary Get all MaximumSpecification.
// @Description Get a list of MaximumSpecification.
// @Tags MaximumSpecification
// @Produce json
// @Success 200 {object} []models.MaximumSpecification
// @Router /maximumspecification [get]
func GetAllMaximumSpecification(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var maximumspec []models.MaximumSpecification
	db.Find(&maximumspec)

	c.JSON(http.StatusOK, gin.H{"data": maximumspec})
}

// CreateMaximumSpecification godoc
// @Summary Create New MaximumSpecification.
// @Description Creating a new MaximumSpecification.
// @Tags MaximumSpecification
// @Param Body body maximumSpecificationInput true "the body to create a new MaximumSpecification"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.MaximumSpecification
// @Router /maximumspecification [post]
func CreateMaximumSpecification(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input maximumSpecificationInput
	var game models.Game
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not found"})
	}

	// Create MaximumSpecification
	maximumSpec := models.MaximumSpecification{
		OS:        input.OS,
		Processor: input.Processor,
		Memory:    input.Memory,
		Graphics:  input.Graphics,
		DirectX:   input.DirectX,
		Network:   input.Network,
		Storage:   input.Storage,
		SoundCard: input.SoundCard,
		GameID:    input.GameID,
	}

	db.Create(&maximumSpec)

	c.JSON(http.StatusOK, gin.H{"data": maximumSpec})
}

// GetMaximumSpecificationById godoc
// @Summary Get MaximumSpecification.
// @Description Get an MaximumSpecification by id.
// @Tags MaximumSpecification
// @Produce json
// @Param id path string true "MaximumSpecification id"
// @Success 200 {object} models.MaximumSpecification
// @Router /maximumspecification/{id} [get]
func GetMaximumSpecificationById(c *gin.Context) {
	// Get model if exist
	var maximumspec models.MaximumSpecification

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&maximumspec).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": maximumspec})
}

// UpdateMaximumSpecification godoc
// @Summary Update MaximumSpecification.
// @Description Update MaximumSpecification by id.
// @Tags MaximumSpecification
// @Produce json
// @Param id path string true "MaximumSpecification id"
// @Param Body body maximumSpecificationInput true "the body to update MaximumSpecification"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.MaximumSpecification
// @Router /maximumspecification/{id} [patch]
func UpdateMaximumSpecification(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var maximumspec models.MaximumSpecification
	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&maximumspec).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// Validate Input
	var input maximumSpecificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not found"})
	}

	var updatedInput models.MaximumSpecification
	updatedInput.OS = input.OS
	updatedInput.Processor = input.Processor
	updatedInput.Memory = input.Memory
	updatedInput.Graphics = input.Graphics
	updatedInput.DirectX = input.DirectX
	updatedInput.Network = input.Network
	updatedInput.Storage = input.Storage
	updatedInput.SoundCard = input.SoundCard
	updatedInput.GameID = input.GameID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&maximumspec).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": maximumspec})
}

// DeleteMaximumSpecification godoc
// @Summary Delete one MaximumSpecification.
// @Description Delete a MaximumSpecification by id.
// @Tags MaximumSpecification
// @Produce json
// @Param id path string true "MaximumSpecification id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /maximumspecification/{id} [delete]
func DeleteMaximumSpecification(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var maximumSpec models.MaximumSpecification
	if err := db.Where("id = ?", c.Param("id")).First(&maximumSpec).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}

	db.Delete(&maximumSpec)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
