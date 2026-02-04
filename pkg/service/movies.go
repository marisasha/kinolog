package service

import (
	"github.com/marisasha/kinolog/pkg/models"
	"github.com/marisasha/kinolog/pkg/repository"
)

type MoviesService struct {
	repos repository.Movies
}

func NewMoviesService(repos repository.Movies) *MoviesService {
	return &MoviesService{repos: repos}
}

func (s *MoviesService) AddMovie(movie *models.Movie, user_id *int) error {
	return s.repos.AddMovie(movie, user_id)
}
func (s *MoviesService) GetAllMovies(user_id *int) ([]*models.Movie, error) {
	return s.repos.GetAllMovies(user_id)
}
func (s *MoviesService) GetMovie(movie_id *int) (*models.Movie, error) {
	return s.repos.GetMovie(movie_id)
}
