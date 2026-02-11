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

func (s *MoviesService) GetAllMovies(user_id *int) ([]*models.Movie, error) {
	return s.repos.GetAllMovies(user_id)
}

func (s *MoviesService) GetMovie(movie_id *int) (*models.Movie, error) {
	return s.repos.GetMovie(movie_id)
}

func (s *MoviesService) DeleteMovie(movie_id *int) error {
	return s.repos.DeleteMovie(movie_id)
}

func (s *MoviesService) ChangeMovieStatus(user_id, movie_id *int, newStatus *string) error {
	return s.repos.ChangeMovieStatus(user_id, movie_id, newStatus)
}

func (s *MoviesService) SearchMovie(title *string, year, user_id *int) (int, error) {
	movieId, err := s.repos.SearchMovieInDB(title, year)
	if err != nil {
		return 0, err
	}
	if movieId != 0 {
		return s.repos.AddUserMovie(user_id, &movieId)
	}
	movie, err := GetMovieInfoFromAI(title, year)
	if err != nil {
		return 0, err
	}
	movieId, err = s.repos.AddMovie(movie)
	if err != nil {
		return 0, err
	}
	return s.repos.AddUserMovie(user_id, &movieId)

}
