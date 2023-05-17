package kafka

import "github.com/segmentio/kafka-go"

// Ping checks if the kafka client is able to connect to the kafka broker.
func (k *kafkaClient) Ping() error {
	conn, err := kafka.Dial("tcp", k.HostAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
