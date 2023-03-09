package cache

import (
	"encoding/json"
	"fmt"

	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/bradfitz/gomemcache/memcache"
)

type FeedbackMemcache struct {
	mc *memcache.Client
}

func NewFeedbackMemcache(mc *memcache.Client) *FeedbackMemcache {
	return &FeedbackMemcache{mc: mc}
}

func (c *FeedbackMemcache) SetFeedback(id int, feedback model.Feedback) error {
	cacheValue, err := json.Marshal(feedback)
	if err != nil {
		return err
	}
	c.mc.Set(&memcache.Item{
		Key:   fmt.Sprintf("feedback:%d", id),
		Value: []byte(cacheValue),
	})
	return nil
}

func (c *FeedbackMemcache) GetFeedback(id int) (model.Feedback, error) {
	var feedback model.Feedback
	item, err := c.mc.Get(fmt.Sprintf("feedback:%d", id))
	if err != nil {
		return model.Feedback{}, err
	}
	json.Unmarshal([]byte(item.Value), &feedback)
	return feedback, nil
}
