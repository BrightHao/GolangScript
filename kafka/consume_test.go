package kafka

import (
	"testing"
)

func TestConsume(t *testing.T) {
	brokers := []string{"10.116.252.15:9092", "10.116.252.26:9092", "10.116.252.6:9092"}
	topic := "telegraf"

	Consume(brokers, topic)
}
