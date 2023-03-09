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

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
