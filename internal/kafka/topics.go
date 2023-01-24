package kafka

import (
	"github.com/rashiraffi/streamease/internal/config"
	"github.com/segmentio/kafka-go"
)

func GetAllTopics() []string {
	conn, err := kafka.Dial("tcp", config.KafkaHost)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	topics := []string{}

	for _, p := range partitions {
		topics = append(topics, p.Topic)
	}
	return topics
}
