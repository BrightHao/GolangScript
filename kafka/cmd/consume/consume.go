package main

import "kafka"

func main() {
	brokers := []string{"10.116.252.15:9092", "10.116.252.26:9092", "10.116.252.6:9092"}
	topic := "telegraf"

	kafka.Consume(brokers, topic)
}
