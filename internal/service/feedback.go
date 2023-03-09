package service

import (
	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/Cheyzie/golang-test/internal/repository"
)

type FeedbackService struct {
	repo repository.Feedback
}

func NewFeedbackService(repo repository.Feedback) *FeedbackService {
	return &FeedbackService{repo: repo}
}

func (s *FeedbackService) GetFeedbackById(id int) (model.Feedback, error) {
	return s.repo.GetById(id)
}

func (s *FeedbackService) GetAllFeedbacks() ([]model.Feedback, error) {
	return s.repo.GetAll()
}

func (s *FeedbackService) GetAllFeedbacksPaginate(limit, offset string) (model.AllFeedbacksResponse, error) {

	feedbacks, err := s.repo.GetAllPaginate(limit, offset)

	if err != nil {
		return model.AllFeedbacksResponse{}, err
	}

	count, err := s.repo.GetCount()

	if err != nil {
		return model.AllFeedbacksResponse{}, err
	}

	response := model.AllFeedbacksResponse{
		Meta: model.Meta{
			Limit:  limit,
			Offset: offset,
			Total:  count,
		},
		Feedbacks: feedbacks,
	}
	return response, nil
}

func (s *FeedbackService) CreateFeedback(feedback model.Feedback) (int, error) {
	return s.repo.Create(feedback)
}
