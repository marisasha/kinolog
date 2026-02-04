package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/marisasha/kinolog/pkg/models"
)

type Authorization interface {
	CreateUser(user *models.User) error
	GetUser(username, password string) (models.User, error)
}

type Movies interface {
	AddMovie(movie *models.Movie, user_id *int) error
	GetAllMovies(user_id *int) ([]*models.Movie, error)
	GetMovie(movie_id *int) (*models.Movie, error)
}

type Repository struct {
	Authorization
	Movies
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Movies:        NewMoviePostgres(db),
	}
}
