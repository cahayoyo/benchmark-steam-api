package controller

import (
	"api-final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type gameInput struct {
	Title               string `json:"title"`
	About               string `json:"about"`
	Year                int    `json:"year"`
	DeveloperID         uint   `json:"developer_id"`
	PublisherID         uint   `json:"publisher_id"`
	GenreID             uint   `json:"genre_id"`
	AgeRatingCategoryID uint   `json:"age_rating_category_id"`
}

// GetAllGame godoc
// @Summary Get all game.
// @Description Get a list of Game.
// @Tags Game
// @Produce json
// @Success 200 {object} []models.Game
// @Router /game [get]
func GetAllGame(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var games []models.Game
	db.Find(&games)

	c.JSON(http.StatusOK, gin.H{"data": games})
}

// CreateGame godoc
// @Summary Create New Game.
// @Description Creating a new Game.
// @Tags Game
// @Param Body body gameInput true "the body to create a new game"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Game
// @Router /game [post]
func CreateGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate Input
	var input gameInput
	var developer models.Developer
	var publisher models.Publisher
	var genre models.Genre
	var rating models.AgeRatingCategory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.DeveloperID).First(&developer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeveloperID not found!"})
		return
	}

	if err := db.Where("id = ?", input.PublisherID).First(&publisher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PublisherID not found!"})
		return
	}

	if err := db.Where("id = ?", input.GenreID).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GenreID not found!"})
		return
	}

	if err := db.Where("id = ?", input.AgeRatingCategoryID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AgeRatingCategoryID not found!"})
		return
	}

	// Create Game
	game := models.Game{
		Title:               input.Title,
		About:               input.About,
		Year:                input.Year,
		DeveloperID:         input.DeveloperID,
		PublisherID:         input.PublisherID,
		GenreID:             input.GenreID,
		AgeRatingCategoryID: input.AgeRatingCategoryID,
	}
	db.Create(&game)

	c.JSON(http.StatusOK, gin.H{"data": game})

}

// GetGameById godoc
// @Summary Get Game.
// @Description Get a Game by id.
// @Tags Game
// @Produce json
// @Param id path string true "game id"
// @Success 200 {object} models.Game
// @Router /game/{id} [get]
func GetGameById(c *gin.Context) { // Get model if exist
	var game models.Game

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// UpdateGame godoc
// @Summary Update Game.
// @Description Update Game by id.
// @Tags Game
// @Produce json
// @Param id path string true "game id"
// @Param Body body gameInput true "the body to update an game"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Game
// @Router /game/{id} [patch]
func UpdateGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get Model If Exist
	var game models.Game
	var developer models.Developer
	var publisher models.Publisher
	var genre models.Genre
	var rating models.AgeRatingCategory
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// Validate Input
	var input gameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.DeveloperID).First(&developer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeveloperID not found!"})
		return
	}

	if err := db.Where("id = ?", input.PublisherID).First(&publisher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PublisherID not found!"})
		return
	}

	if err := db.Where("id = ?", input.GenreID).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GenreID not found!"})
		return
	}

	if err := db.Where("id = ?", input.AgeRatingCategoryID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AgeRatingCategoryID not found!"})
		return
	}

	var updatedInput models.Game
	updatedInput.Title = input.Title
	updatedInput.About = input.About
	updatedInput.Year = input.Year
	updatedInput.DeveloperID = input.DeveloperID
	updatedInput.PublisherID = input.PublisherID
	updatedInput.GenreID = input.GenreID
	updatedInput.AgeRatingCategoryID = input.AgeRatingCategoryID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&game).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": game})

}

// DeleteGame godoc
// @Summary Delete one Game.
// @Description Delete a Game by id.
// @Tags Game
// @Produce json
// @Param id path string true "game id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /game/{id} [delete]
func DeleteGame(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&game)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
