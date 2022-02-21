package kafka

import (
	"log"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consume(brokers []string, topic string) {
	if len(topic) == 0 {
		log.Fatal("topic or tp empty.")
	}

	c, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     strings.Join(brokers, ","),
		"broker.address.family": "v4",
		"group.id":              "falcon-consumer-group-v2",
		"session.timeout.ms":    10000,
		// "auto.offset.reset":     "earliest",
		"auto.offset.reset":  "latest",
		"enable.auto.commit": false,
	})

	topics := strings.Split(topic, ",")
	log.Println(topics)
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("\nheaders: ", msg.Headers, "\n msg:", string(msg.Value))
		}
	}
}
