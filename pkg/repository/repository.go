package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/marisasha/kinolog"
)

type Authorization interface {
	CreateUser(user *kinolog.User) error
	GetUser(username, password string) (kinolog.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
