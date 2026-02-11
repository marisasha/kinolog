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

func (s *MoviesService) GetAllMovies(userId *int) ([]*models.Movie, error) {
	return s.repos.GetAllMovies(userId)
}

func (s *MoviesService) GetMovie(movie_id *int) (*models.Movie, error) {
	return s.repos.GetMovie(movie_id)
}

func (s *MoviesService) DeleteMovie(movie_id *int) error {
	return s.repos.DeleteMovie(movie_id)
}

func (s *MoviesService) ChangeMovieStatus(userId, movieId, mark *int, newStatus, review *string) error {
	return s.repos.ChangeMovieStatus(userId, movieId, mark, newStatus, review)
}

func (s *MoviesService) SearchMovie(title *string, year, userId *int) (*models.Movie, error) {

	movieId, err := s.repos.SearchMovieInDB(title, year)
	if err != nil {
		return nil, err
	}

	if movieId != 0 {
		err = s.repos.AddUserMovie(userId, &movieId)
		if err != nil {
			return nil, err
		}
		return s.repos.GetMovie(&movieId)

	}

	movie, err := GetMovieInfoFromAI(title, year)
	if err != nil {
		return nil, err
	}

	movieId, err = s.repos.AddMovie(movie)
	if err != nil {
		return nil, err
	}

	err = s.repos.AddUserMovie(userId, &movieId)
	if err != nil {
		return nil, err
	}
	return s.repos.GetMovie(&movieId)
}
