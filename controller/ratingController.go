package controller

import (
	"api-final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ratingInput struct {
	Rating int  `json:"rating"`
	GameID uint `json:"game_id"`
}

// GetAllRatingGame godoc
// @Summary Get all RatingGame.
// @Description Get a list of RatingGame.
// @Tags RatingGame
// @Produce json
// @Success 200 {object} []models.Rating
// @Router /ratinggame [get]
func GetAllRatingGame(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var rating []models.Rating
	db.Find(&rating)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// CreateRatingGame godoc
// @Summary Create New RatingGame.
// @Description Creating a new RatingGame.
// @Tags RatingGame
// @Param Body body ratingInput true "the body to create a new game"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Game
// @Router /ratinggame [post]
func CreateRatingGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate Input
	var input ratingInput
	var game models.Game
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if input.Rating > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum 100 for Rating"})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not found!"})
	}

	// Create Rating Game
	rating := models.Rating{
		Rating: input.Rating,
		GameID: input.GameID,
	}
	db.Create(&rating)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// GetRatingGameById godoc
// @Summary Get RatingGame.
// @Description Get a RatingGame by id.
// @Tags RatingGame
// @Produce json
// @Param id path string true "ratinggame id"
// @Success 200 {object} models.Rating
// @Router /ratinggame/{id} [get]
func GetRatingGameById(c *gin.Context) { // Get model if exist
	var rating models.Rating

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// UpdateRatingGame godoc
// @Summary Update RatingGame.
// @Description Update RatingGame by id.
// @Tags RatingGame
// @Produce json
// @Param id path string true "ratinggame id"
// @Param Body body ratingInput true "the body to update an RatingGame"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Rating
// @Router /ratinggame/{id} [patch]
func UpdateRatingGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get Model If Exist
	var rating models.Rating
	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// Validate Input
	var input ratingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if input.Rating > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum 100 for Rating"})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not Found!"})
	}

	var updatedInput models.Rating
	updatedInput.Rating = input.Rating
	updatedInput.GameID = input.GameID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&rating).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// DeleteRatingGame godoc
// @Summary Delete one RatingGame.
// @Description Delete a RatingGame by id.
// @Tags RatingGame
// @Produce json
// @Param id path string true "ratinggame id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /ratinggame/{id} [delete]
func DeleteRatingGame(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var rating models.Rating
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&rating)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
