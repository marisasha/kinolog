package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/marisasha/kinolog/pkg/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/marisasha/kinolog/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		movies := api.Group("/movies")
		{
			movies.GET("/", h.getAllMovies)
			movies.GET("/:id", h.getMovie)
			movies.GET("/ai/search", h.getMovieInformation)
			movies.POST("/add", h.addMovie)
			movies.POST("/status/change", h.changeMovieStatus)
			movies.DELETE("/delete", h.deleteMovie)
		}

		friends := api.Group("/friends")
		{
			friends.GET("/", h.getAllFriends)
			friends.GET("/:id", h.getFriend)
			friends.POST("/add", h.addFriend)
			friends.PUT("/accept", h.acceptFriendRequest)
			friends.DELETE("/delete", h.refuseFriendRequest)
			friends.DELETE("/refuse", h.deleteFriend)
		}
	}

	return router
}
