package service

import (
	"github.com/Cheyzie/golang-test/internal/cache"
	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/Cheyzie/golang-test/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(email, password string) (model.Token, error)
	ParseToken(access_token string) (int, error)
}

type Feedback interface {
	CreateFeedback(feedback model.Feedback) (int, error)
	GetAllFeedbacks() ([]model.Feedback, error)
	GetAllFeedbacksPaginate(limit, offset string) (model.AllFeedbacksResponse, error)
	GetFeedbackById(id int) (model.Feedback, error)
}

type Service struct {
	Authorization
	Feedback
}

func NewService(repo *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repo.User),
		Feedback:      NewFeedbackService(repo.Feedback, cache.Feedback),
	}
}
