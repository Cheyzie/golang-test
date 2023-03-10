package broker

import (
	"github.com/Cheyzie/golang-test/internal/model"
)

type FeedbackProducer interface {
	SendFeedback(topic string, feedback model.Feedback)
}

type Producer struct {
	FeedbackProducer
}

func NewProducer(d *KafkaDriver) *Producer {
	return &Producer{
		FeedbackProducer: NewFeedbackKafkaProducer(d),
	}
}
