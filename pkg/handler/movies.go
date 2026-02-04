package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marisasha/kinolog/pkg/models"
)

type getAllMovieResponse struct {
	Data []*models.Movie
}
type getMovieResponse struct {
	Data *models.Movie
}

// @Summary Посмотреть все фильмы
// @Tags movies
// @Description Посмотреть все фильмы пользователя
// @ID get-movies
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /api/movies/ [get]
func (h *Handler) getAllMovies(c *gin.Context) {
	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	movies, err := h.services.Movies.GetAllMovies(&user_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	c.JSON(http.StatusAccepted, getAllMovieResponse{
		Data: movies,
	})
}

// @Summary Посмотреть фильм
// @Tags movies
// @Description Посмотреть фильм пользователя
// @ID get-movie
// @Accept json
// @Produce json
// @Param id path int true "Id movie"
// @Security ApiKeyAuth
// @Router /api/movies/{id} [get]
func (h *Handler) getMovie(c *gin.Context) {
	movie_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	movie, err := h.services.Movies.GetMovie(&movie_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	c.JSON(http.StatusAccepted, getMovieResponse{
		Data: movie,
	})
}

func (h *Handler) getMovieInformation(c *gin.Context) {

}

// @Summary Добавить кино
// @Tags movies
// @Description Добавляет новое кино
// @ID add-movie
// @Accept json
// @Produce json
// @Param request body models.Movie true "Информация о кино"
// @Security ApiKeyAuth
// @Router /api/movies/add [post]
func (h *Handler) addMovie(c *gin.Context) {
	var input models.Movie

	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Movies.AddMovie(&input, &user_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]string{
		"message": "movie sucsuccessfully created",
	})

}

func (h *Handler) changeMovieStatus(c *gin.Context) {

}

func (h *Handler) deleteMovie(c *gin.Context) {

}
