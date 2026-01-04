package kafka

import (
    "context"
    "github.com/segmentio/kafka-go"
)

func NewConsumer() *kafka.Reader {
    return kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "images",
        GroupID: "image-workers",
    })
}
