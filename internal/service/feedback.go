package service

import (
	"github.com/Cheyzie/golang-test/internal/broker"
	"github.com/Cheyzie/golang-test/internal/cache"
	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/Cheyzie/golang-test/internal/repository"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/sirupsen/logrus"
)

type FeedbackService struct {
	repo     repository.Feedback
	cache    cache.Feedback
	producer broker.FeedbackProducer
}

func NewFeedbackService(repo repository.Feedback, cache cache.Feedback, producer broker.FeedbackProducer) *FeedbackService {
	return &FeedbackService{repo: repo, cache: cache, producer: producer}
}

func (s *FeedbackService) GetFeedbackById(id int) (model.Feedback, error) {
	feedback, err := s.cache.GetFeedback(id)

	if err == nil {
		return feedback, err
	} else if err != memcache.ErrCacheMiss {
		logrus.Error(err.Error())
	}

	feedback, err = s.repo.GetById(id)
	if err != nil {
		return feedback, err
	}

	s.cache.SetFeedback(feedback.Id, feedback)

	return feedback, err
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
	id, err := s.repo.Create(feedback)
	if err != nil {
		return id, err
	}
	feedback.Id = id
	s.producer.SendFeedback("feedbacks", feedback)
	return id, err
}
