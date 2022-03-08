package controller

import (
	"api-final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type genreInput struct {
	Name string `json:"name"`
}

// GetAllGenre godoc
// @Summary Get all Genre.
// @Description Get a list of Genre.
// @Tags Genre
// @Produce json
// @Success 200 {object} []models.Genre
// @Router /genre [get]
func GetAllGenre(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var genres []models.Genre
	db.Find(&genres)

	c.JSON(http.StatusOK, gin.H{"data": genres})
}

// CreateGenre godoc
// @Summary Create New Genre.
// @Description Creating a new Genre.
// @Tags Genre
// @Param Body body genreInput true "the body to create a new Genre"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Genre
// @Router /genre [post]
func CreateGenre(c *gin.Context) {
	// Validate input
	var input genreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Genre
	genre := models.Genre{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&genre)

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

// GetGenreById godoc
// @Summary Get Genre.
// @Description Get an Genre by id.
// @Tags Genre
// @Produce json
// @Param id path string true "Genre id"
// @Success 200 {object} models.Genre
// @Router /genre/{id} [get]
func GetGenreById(c *gin.Context) {
	// Get model if exist
	var genre models.Genre

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

// GetGameByGenreId godoc
// @Summary Get Game.
// @Description Get all Game by GenreId.
// @Tags Genre
// @Produce json
// @Param id path string true "Genre id"
// @Success 200 {object} []models.Game
// @Router /genre/{id}/game [get]
func GetGameByGenreId(c *gin.Context) {
	// Get model if exist
	var games []models.Game

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("genre_id = ?", c.Param("id")).Find(&games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": games})
}

// UpdateGenre godoc
// @Summary Update Genre.
// @Description Update Genre by id.
// @Tags Genre
// @Produce json
// @Param id path string true "Genre id"
// @Param Body body genreInput true "the body to update genre"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Genre
// @Router /genre/{id} [patch]
func UpdateGenre(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var genre models.Genre
	if err := db.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// Validate Input
	var input developerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Genre
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&genre).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

// DeleteGenre godoc
// @Summary Delete one Genre.
// @Description Delete a Genre by id.
// @Tags Genre
// @Produce json
// @Param id path string true "Genre id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /genre/{id} [delete]
func DeleteGenre(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var genre models.Genre
	if err := db.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}

	db.Delete(&genre)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
