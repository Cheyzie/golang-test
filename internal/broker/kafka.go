package broker

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaDriver struct {
	cfg *kafka.ConfigMap
}

func NewKafkaDriver(host string) *KafkaDriver {
	cfg := &kafka.ConfigMap{}
	cfg.SetKey("bootstrap.servers", host)
	return &KafkaDriver{cfg: cfg}
}

func (d *KafkaDriver) GetProducer() (*kafka.Producer, error) {
	return kafka.NewProducer(d.cfg)
}
