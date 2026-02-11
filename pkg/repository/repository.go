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
	GetAllMovies(user_id *int) ([]*models.Movie, error)
	GetMovie(movie_id *int) (*models.Movie, error)
	DeleteMovie(movie_id *int) error
	ChangeMovieStatus(user_id, movie_id *int, newStatus *string) error
	SearchMovieInDB(title *string, year *int) (int, error)
	AddUserMovie(user_id, movieId *int) (int, error)
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
