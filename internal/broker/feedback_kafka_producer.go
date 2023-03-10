package broker

import (
	"encoding/json"

	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type FeedbackKafkaProducer struct {
	driver *KafkaDriver
}

func NewFeedbackKafkaProducer(d *KafkaDriver) *FeedbackKafkaProducer {
	return &FeedbackKafkaProducer{driver: d}
}

func (fp *FeedbackKafkaProducer) SendFeedback(topic string, feedback model.Feedback) {

	kp, err := fp.driver.GetProducer()

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	message_value, err := json.Marshal(feedback)

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	delivery_chan := make(chan kafka.Event, 100000)

	err = kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message_value),
	}, delivery_chan)

	if err != nil {
		logrus.Error(err.Error())
		return
	}
	e := <-delivery_chan
	message := e.(*kafka.Message)

	if message.TopicPartition.Error != nil {
		logrus.Error("Feedback push to queue failed: " + message.TopicPartition.Error.Error())
	}
	close(delivery_chan)
	kp.Close()
}
