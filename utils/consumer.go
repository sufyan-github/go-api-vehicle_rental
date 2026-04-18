package utils

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "booking-events",
		GroupID: "booking-group",
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		log.Println("Received event:", string(msg.Value))
	}
}