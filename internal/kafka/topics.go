package kafka

import (
	"sort"

	"github.com/segmentio/kafka-go"
)

// GetAllTopics returns a slice of strings containing all the topics in the Kafka cluster.
func (k *kafkaClient) GetAllTopics() []string {
	conn, err := kafka.Dial("tcp", k.HostAddr)
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
		if p.Topic == "__consumer_offsets" {
			continue
		}
		topics = append(topics, p.Topic)
	}
	sort.Strings(topics)
	return topics
}
