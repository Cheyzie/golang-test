package service

import (
	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/Cheyzie/golang-test/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(email, password string) (model.Token, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repo.User),
	}
}
