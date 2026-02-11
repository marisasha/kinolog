package service

import (
	"github.com/marisasha/kinolog/pkg/models"
	"github.com/marisasha/kinolog/pkg/repository"
)

type Authorization interface {
	CreateUser(user *models.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Movies interface {
	// AddMovie(movie *models.Movie) error
	GetAllMovies(user_id *int) ([]*models.Movie, error)
	GetMovie(movie_id *int) (*models.Movie, error)
	DeleteMovie(movie_id *int) error
	ChangeMovieStatus(user_id, movie_id *int, newStatus *string) error
	SearchMovie(title *string, year, user_id *int) (int, error)
}

type Service struct {
	Authorization
	Movies
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Movies:        NewMoviesService(repos.Movies),
	}
}
