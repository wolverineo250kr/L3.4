package kafka

import (
    "context"
    "github.com/segmentio/kafka-go"
)

var Writer = kafka.NewWriter(kafka.WriterConfig{
    Brokers: []string{"localhost:9092"},
    Topic:   "images",
})

func Send(id string) error {
    return Writer.WriteMessages(context.Background(),
        kafka.Message{Value: []byte(id)},
    )
}
