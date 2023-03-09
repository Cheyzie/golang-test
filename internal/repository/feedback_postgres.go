package repository

import (
	"fmt"

	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/jmoiron/sqlx"
)

type FeedbackPostgres struct {
	db *sqlx.DB
}

func NewFeedbackPostgres(db *sqlx.DB) *FeedbackPostgres {
	return &FeedbackPostgres{db: db}
}

func (r *FeedbackPostgres) Create(feedback model.Feedback) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (customer_name, email, feedback_text, source) VALUES ($1, $2, $3, $4) RETURNING id;", feedbacksTable)
	row := r.db.QueryRow(query, feedback.CustomerName, feedback.Email, feedback.FeedbackText, feedback.Source)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *FeedbackPostgres) GetAllPaginate(limit, offset string) ([]model.Feedback, error) {
	var feedbacks []model.Feedback

	query := fmt.Sprintf("SELECT id, customer_name, email, feedback_text, source FROM %s LIMIT $1 OFFSET $2;", feedbacksTable)
	err := r.db.Select(&feedbacks, query, limit, offset)

	return feedbacks, err
}

func (r *FeedbackPostgres) GetAll() ([]model.Feedback, error) {
	var feedbacks []model.Feedback

	query := fmt.Sprintf("SELECT id, customer_name, email, feedback_text, source FROM %s;", feedbacksTable)
	err := r.db.Select(&feedbacks, query)

	return feedbacks, err
}

func (r *FeedbackPostgres) GetCount() (int, error) {
	var count int

	query := fmt.Sprintf("SELECT count(*) FROM %s;", feedbacksTable)
	row := r.db.QueryRow(query)

	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *FeedbackPostgres) GetById(id int) (model.Feedback, error) {
	var feedback model.Feedback

	query := fmt.Sprintf("SELECT id, customer_name, email, feedback_text, source FROM %s WHERE ID = $1;", feedbacksTable)
	err := r.db.Get(&feedback, query, id)

	return feedback, err
}
