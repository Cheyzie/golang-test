package model

type Meta struct {
	Limit  string
	Offset string
	Total  int
}

type Feedback struct {
	Id           int    `json:"id" db:"id"`
	CustomerName string `json:"customer_name" db:"customer_name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	FeedbackText string `json:"feedback_text" db:"feedback_text" binding:"required"`
	Source       string `json:"source" binding:"required"`
}

type AllFeedbacksResponse struct {
	Meta      Meta
	Feedbacks []Feedback
}
