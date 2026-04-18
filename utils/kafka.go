package utils

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

var writer = kafka.NewWriter(kafka.WriterConfig{
	Brokers: []string{"localhost:9092"},
	Topic:   "booking-events",
})

func PublishBookingEvent(data interface{}) error {

	msg, _ := json.Marshal(data)

	return writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: msg,
		},
	)
}