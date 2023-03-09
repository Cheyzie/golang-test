package repository

import (
	"fmt"

	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetByEmail(email string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT user.id, user.username, user.email FROM %s user WHERE user.id = $1;", usersTable)
	err := r.db.Get(&user, query, email)

	return user, err
}

func (r *UserPostgres) GetByCredentials(email, password_hash string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id, username, email FROM %s WHERE email = $1 AND password_hash =$2;", usersTable)
	err := r.db.Get(&user, query, email, password_hash)

	return user, err
}

func (r *UserPostgres) Create(user model.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id;", usersTable)
	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
