package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/marisasha/kinolog/internal/service"
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

	router.Use(h.loggingMiddleware)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		movies := api.Group("/movies")
		{
			movies.GET("/", h.getAllMovies)
			movies.GET("/:id", h.getMovie)
			movies.GET("/ai/search", h.getMovieInformation)
			movies.PUT("/status/change", h.changeMovieStatus)
			movies.DELETE("/delete/:id", h.deleteMovie)
		}
	}

	return router
}
