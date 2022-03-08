package controller

import (
	"api-final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewInput struct {
	Comment string `json:"comment"`
	GameID  uint   `json:"game_id"`
}

// GetAllReviewGame godoc
// @Summary Get all ReviewGame.
// @Description Get a list of ReviewGame.
// @Tags ReviewGame
// @Produce json
// @Success 200 {object} []models.Review
// @Router /reviewgame [get]
func GetAllReviewGame(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var review []models.Review
	db.Find(&review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// CreateReviewGame godoc
// @Summary Create New ReviewGame.
// @Description Creating a new ReviewGame.
// @Tags ReviewGame
// @Param Body body reviewInput true "the body to create a new reviewgame"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Review
// @Router /reviewgame [post]
func CreateReviewGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate Input
	var input reviewInput
	var game models.Game
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not found!"})
	}

	// Create Review Game
	review := models.Review{
		Comment: input.Comment,
		GameID:  input.GameID,
	}
	db.Create(&review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// GetReviewGameById godoc
// @Summary Get ReviewGame.
// @Description Get a ReviewGame by id.
// @Tags ReviewGame
// @Produce json
// @Param id path string true "reviewgame id"
// @Success 200 {object} models.Review
// @Router /reviewgame/{id} [get]
func GetReviewGameById(c *gin.Context) { // Get model if exist
	var review models.Review

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// UpdateReviewGame godoc
// @Summary Update ReviewGame.
// @Description Update ReviewGame by id.
// @Tags ReviewGame
// @Produce json
// @Param id path string true "reviewgame id"
// @Param Body body reviewInput true "the body to update an ReviewGame"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Review
// @Router /reviewgame/{id} [patch]
func UpdateReviewGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get Model If Exist
	var review models.Review
	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// Validate Input
	var input reviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not Found!"})
	}

	var updatedInput models.Review
	updatedInput.Comment = input.Comment
	updatedInput.GameID = input.GameID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&review).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DeleteReviewGame godoc
// @Summary Delete one ReviewGame.
// @Description Delete a ReviewGame by id.
// @Tags ReviewGame
// @Produce json
// @Param id path string true "reviewgame id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /reviewgame/{id} [delete]
func DeleteReviewGame(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&review)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
