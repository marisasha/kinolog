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
	AddMovie(movie *models.Movie) (int, error)
	GetAllMovies(userId *int) ([]*models.Movie, error)
	GetMovie(movie_id *int) (*models.Movie, error)
	DeleteMovie(movie_id *int) error
	ChangeMovieStatus(userId, movieId, mark *int, newStatus, review *string) error
	SearchMovieInDB(title *string, year *int) (int, error)
	AddUserMovie(userId, movieId *int) error
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
