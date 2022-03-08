package controller

import (
	"api-final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type minimumSpecificationInput struct {
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

// GetAllMinimumSpecification godoc
// @Summary Get all MinimumSpecification.
// @Description Get a list of MinimumSpecification.
// @Tags MinimumSpecification
// @Produce json
// @Success 200 {object} []models.MinimumSpecification
// @Router /minimumspecification [get]
func GetAllMinimumSpecification(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var minimumspec []models.MinimumSpecification
	db.Find(&minimumspec)

	c.JSON(http.StatusOK, gin.H{"data": minimumspec})
}

// CreateMinimumSpecification godoc
// @Summary Create New MinimumSpecification.
// @Description Creating a new MinimumSpecification.
// @Tags MinimumSpecification
// @Param Body body minimumSpecificationInput true "the body to create a new MinimumSpecification"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.MinimumSpecification
// @Router /minimumspecification [post]
func CreateMinimumSpecification(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input minimumSpecificationInput
	var game models.Game
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not found"})
	}

	// Create MinimumSpecification
	minimumspec := models.MinimumSpecification{
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

	db.Create(&minimumspec)

	c.JSON(http.StatusOK, gin.H{"data": minimumspec})
}

// GetMinimumSpecificationById godoc
// @Summary Get MinimumSpecification.
// @Description Get an MinimumSpecification by id.
// @Tags MinimumSpecification
// @Produce json
// @Param id path string true "MinimumSpecification id"
// @Success 200 {object} models.MinimumSpecification
// @Router /minimumspecification/{id} [get]
func GetMinimumSpecificationById(c *gin.Context) {
	// Get model if exist
	var minimumspec models.MinimumSpecification

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&minimumspec).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": minimumspec})
}

// UpdateMinimumSpecification godoc
// @Summary Update MinimumSpecification.
// @Description Update MinimumSpecification by id.
// @Tags MinimumSpecification
// @Produce json
// @Param id path string true "MinimumSpecification id"
// @Param Body body minimumSpecificationInput true "the body to update MinimumSpecification"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.MinimumSpecification
// @Router /minimumspecification/{id} [patch]
func UpdateMinimumSpecification(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var minimumspec models.MinimumSpecification
	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&minimumspec).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// Validate Input
	var input minimumSpecificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not found"})
	}

	var updatedInput models.MinimumSpecification
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

	db.Model(&minimumspec).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": minimumspec})
}

// DeleteMinimumSpecification godoc
// @Summary Delete one MinimumSpecification.
// @Description Delete a MinimumSpecification by id.
// @Tags MinimumSpecification
// @Produce json
// @Param id path string true "MinimumSpecification id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /minimumspecification/{id} [delete]
func DeleteMinimumSpecification(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var minimumspec models.MinimumSpecification
	if err := db.Where("id = ?", c.Param("id")).First(&minimumspec).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}

	db.Delete(&minimumspec)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
