package repository

import (
	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user model.User) (int, error)
	GetByEmail(email string) (model.User, error)
	GetByCredentials(email, password_hash string) (model.User, error)
}

type Feedback interface {
	Create(feedback model.Feedback) (int, error)
	GetAllPaginate(limit, offset string) ([]model.Feedback, error)
	GetAll() ([]model.Feedback, error)
	GetCount() (int, error)
	GetById(id int) (model.Feedback, error)
}

type Repository struct {
	User
	Feedback
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:     NewUserPostgres(db),
		Feedback: NewFeedbackPostgres(db),
	}
}
