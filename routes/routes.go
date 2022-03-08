package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"api-final-project/controller"
	"api-final-project/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// Auth
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.PATCH("/changepassword/:username", controller.ChangePassword)

	// Game
	r.GET("/game", controller.GetAllGame)
	r.GET("/game/:id", controller.GetGameById)

	gameMiddlewareRoute := r.Group("/game")
	gameMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	gameMiddlewareRoute.POST("/", controller.CreateGame)
	gameMiddlewareRoute.PATCH("/:id", controller.UpdateGame)
	gameMiddlewareRoute.DELETE("/:id", controller.DeleteGame)

	// Rating Game
	r.GET("/ratinggame", controller.GetAllRatingGame)
	r.GET("/ratinggame/:id", controller.GetRatingGameById)

	ratingGameMiddlewareRoute := r.Group("/ratinggame")
	ratingGameMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	ratingGameMiddlewareRoute.POST("/", controller.CreateRatingGame)
	ratingGameMiddlewareRoute.PATCH("/:id", controller.UpdateRatingGame)
	ratingGameMiddlewareRoute.DELETE("/:id", controller.DeleteRatingGame)

	// Review Game
	r.GET("/reviewgame", controller.GetAllReviewGame)
	r.GET("/reviewgame/:id", controller.GetReviewGameById)

	reviewMiddlewareRoute := r.Group("/reviewgame")
	reviewMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	reviewMiddlewareRoute.POST("/", controller.CreateReviewGame)
	reviewMiddlewareRoute.PATCH("/:id", controller.UpdateReviewGame)
	reviewMiddlewareRoute.DELETE("/:id", controller.DeleteRatingGame)

	// Age Rating Categories
	r.GET("/age-rating-categories", controller.GetAllRating)
	r.GET("/age-rating-categories/:id", controller.GetRatingById)
	r.GET("/age-rating-categories/:id/game", controller.GetGameByRatingId)

	ratingMiddlewareRoute := r.Group("/age-rating-categories")
	ratingMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	ratingMiddlewareRoute.POST("/", controller.CreateRating)
	ratingMiddlewareRoute.PATCH("/:id", controller.UpdateRating)
	ratingMiddlewareRoute.DELETE("/:id", controller.DeleteRating)

	// Developer
	r.GET("/developer", controller.GetAllDeveloper)
	r.GET("/developer/:id", controller.GetDeveloperById)
	r.GET("/developer/:id/game", controller.GetGameByDeveloperId)

	developerMiddlewareRoute := r.Group("/developer")
	developerMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	developerMiddlewareRoute.POST("/", controller.CreateDeveloper)
	developerMiddlewareRoute.PATCH("/:id", controller.UpdateDeveloper)
	developerMiddlewareRoute.DELETE("/:id", controller.DeleteDeveloper)

	// Publisher
	r.GET("/publisher", controller.GetAllPublisher)
	r.GET("/publisher/:id", controller.GetPublisherById)
	r.GET("/publisher/:id/game", controller.GetGameByPublisherId)

	publisherMiddlewareRoute := r.Group("/publisher")
	publisherMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	publisherMiddlewareRoute.POST("/", controller.CreatePublisher)
	publisherMiddlewareRoute.PATCH("/:id", controller.UpdatePublisher)
	publisherMiddlewareRoute.DELETE("/:id", controller.DeletePublisher)

	// Genre
	r.GET("/genre", controller.GetAllGenre)
	r.GET("/genre/:id", controller.GetGenreById)
	r.GET("/genre/:id/game", controller.GetGameByGenreId)

	genreMiddlewareRoute := r.Group("/genre")
	genreMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	genreMiddlewareRoute.POST("/", controller.CreateGenre)
	genreMiddlewareRoute.PATCH("/:id", controller.UpdateGenre)
	genreMiddlewareRoute.DELETE("/:id", controller.DeleteGenre)

	// Maximum Specification
	r.GET("/maximumspecification", controller.GetAllMaximumSpecification)
	r.GET("/maximumspecification/:id", controller.GetMaximumSpecificationById)

	maxspecMiddlewareRoute := r.Group("/maximumspecification")
	maxspecMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	maxspecMiddlewareRoute.POST("/", controller.CreateMaximumSpecification)
	maxspecMiddlewareRoute.PATCH("/:id", controller.UpdateMaximumSpecification)
	maxspecMiddlewareRoute.DELETE("/:id", controller.DeleteMaximumSpecification)

	// Minimum Specification
	r.GET("/minimumspecification", controller.GetAllMinimumSpecification)
	r.GET("/minimumspecification/:id", controller.GetMinimumSpecificationById)

	minspecMiddlewareRoute := r.Group("/minimumspecification")
	minspecMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	minspecMiddlewareRoute.POST("/", controller.CreateMinimumSpecification)
	minspecMiddlewareRoute.PATCH("/:id", controller.UpdateMinimumSpecification)
	minspecMiddlewareRoute.DELETE("/:id", controller.DeleteMinimumSpecification)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
