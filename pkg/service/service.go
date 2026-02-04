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
	AddMovie(movie *models.Movie, user_id *int) error
	GetAllMovies(user_id *int) ([]*models.Movie, error)
	GetMovie(movie_id *int) (*models.Movie, error)
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
