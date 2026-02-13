package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/marisasha/kinolog/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user *models.User) error {
	query := fmt.Sprintf("INSERT INTO %s (email,password_hash,first_name,last_name) VALUES ($1, $2, $3, $4)", userTable)
	_, err := r.db.Exec(query, user.Email, user.Password, user.FirstName, user.LastName)
	if err != nil {
		return err
	}
	return nil

}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {

	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err

}
