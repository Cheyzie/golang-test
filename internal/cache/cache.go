package cache

import (
	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/bradfitz/gomemcache/memcache"
)

type Feedback interface {
	SetFeedback(id int, value model.Feedback) error
	GetFeedback(id int) (model.Feedback, error)
}

type Cache struct {
	Feedback
}

func NewCache(mc *memcache.Client) *Cache {
	return &Cache{
		Feedback: NewFeedbackMemcache(mc),
	}
}
