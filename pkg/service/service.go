package service

import (
	"github.com/marisasha/kinolog"
	"github.com/marisasha/kinolog/pkg/repository"
)

type Authorization interface {
	CreateUser(user *kinolog.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
