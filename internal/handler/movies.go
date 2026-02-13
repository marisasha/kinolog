package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marisasha/kinolog/internal/models"
)

type getAllMovieResponse struct {
	Data []*models.Movie
}
type getMovieResponse struct {
	Data *models.Movie
}
type changeMovieStatusRequest struct {
	MovieId int    `json:"movie_id"`
	Status  string `json:"status"`
	Mark    int    `json:"mark"`
	Review  string `json:"review"`
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

// @Summary Найти фильм
// @Tags movies
// @Description Найти фильм
// @ID search-movie
// @Accept json
// @Produce json
// @Param title query string false "Название фильма"
// @Param year query integer false "Год выпуска"
// @Security ApiKeyAuth
// @Router /api/movies/ai/search [get]
func (h *Handler) getMovieInformation(c *gin.Context) {

	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	title := c.Query("title")
	yearStr := c.Query("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	movie, err := h.services.Movies.SearchMovie(&title, &year, &user_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	c.JSON(http.StatusCreated, getMovieResponse{
		Data: movie,
	})

}

// @Summary Обновить статус кино
// @Tags movies
// @Description Обновляет статус  кино
// @ID update-status-movie
// @Accept json
// @Produce json
// @Param request body  changeMovieStatusRequest true "Id movie and status"
// @Security ApiKeyAuth
// @Router /api/movies/status/change [put]
func (h *Handler) changeMovieStatus(c *gin.Context) {

	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	var input changeMovieStatusRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Movies.ChangeMovieStatus(&user_id, &input.MovieId, &input.Mark, &input.Status, &input.Review)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	c.JSON(http.StatusAccepted, map[string]string{
		"message": "movie status sucsuccessfully changed",
	})
}

// @Summary Удалить фильм
// @Tags movies
// @Description Посмотреть фильм пользователя
// @ID delete-movie
// @Accept json
// @Produce json
// @Param id path int true "Id movie"
// @Security ApiKeyAuth
// @Router /api/movies/delete/{id} [delete]
func (h *Handler) deleteMovie(c *gin.Context) {
	movie_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	err = h.services.Movies.DeleteMovie(&movie_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	c.JSON(http.StatusAccepted, map[string]string{
		"message": "movie sucsuccessfully deleted",
	})
}
