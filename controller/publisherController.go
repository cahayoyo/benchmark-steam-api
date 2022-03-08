package controller

import (
	"api-final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type publisherInput struct {
	Name string `json:"name"`
}

// GetAllPublisher godoc
// @Summary Get all Publisher.
// @Description Get a list of Publisher.
// @Tags Publisher
// @Produce json
// @Success 200 {object} []models.Publisher
// @Router /publisher [get]
func GetAllPublisher(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var publisher []models.Publisher
	db.Find(&publisher)

	c.JSON(http.StatusOK, gin.H{"data": publisher})
}

// CreatePublisher godoc
// @Summary Create New Publisher.
// @Description Creating a new Publisher.
// @Tags Publisher
// @Param Body body publisherInput true "the body to create a new Publisher"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Publisher
// @Router /publisher [post]
func CreatePublisher(c *gin.Context) {
	// Validate input
	var input publisherInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Publisher
	publisher := models.Publisher{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&publisher)

	c.JSON(http.StatusOK, gin.H{"data": publisher})
}

// GetPublisherById godoc
// @Summary Get Publisher.
// @Description Get an Publisher by id.
// @Tags Publisher
// @Produce json
// @Param id path string true "Publisher id"
// @Success 200 {object} models.Publisher
// @Router /publisher/{id} [get]
func GetPublisherById(c *gin.Context) {
	// Get model if exist
	var publisher models.Publisher

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&publisher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": publisher})
}

// GetGameByPublisherId godoc
// @Summary Get Game.
// @Description Get all Game by PublisherId.
// @Tags Publisher
// @Produce json
// @Param id path string true "Publisher id"
// @Success 200 {object} []models.Game
// @Router /publisher/{id}/game [get]
func GetGameByPublisherId(c *gin.Context) {
	// Get model if exist
	var games []models.Game

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("publisher_id = ?", c.Param("id")).Find(&games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": games})
}

// UpdatePublisher godoc
// @Summary Update Publisher.
// @Description Update Publisher by id.
// @Tags Publisher
// @Produce json
// @Param id path string true "Publisher id"
// @Param Body body publisherInput true "the body to update Publisher"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Publisher
// @Router /publisher/{id} [patch]
func UpdatePublisher(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var publisher models.Publisher
	if err := db.Where("id = ?", c.Param("id")).First(&publisher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// Validate Input
	var input publisherInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Publisher
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&publisher).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": publisher})
}

// DeletePublisher godoc
// @Summary Delete one Publisher.
// @Description Delete a Publisher by id.
// @Tags Publisher
// @Produce json
// @Param id path string true "Publisher id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /publisher/{id} [delete]
func DeletePublisher(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var publisher models.Publisher
	if err := db.Where("id = ?", c.Param("id")).First(&publisher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}

	db.Delete(&publisher)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
