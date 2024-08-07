package kafka

import (
	"encoding/json"
	"user_service/internal/models"

	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(broker, topic string) (*Producer, error) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &Producer{writer: writer}, nil
}

func (p *Producer) Produce(user models.User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(user.ID),
		Value: userBytes,
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
