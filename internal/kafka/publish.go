package kafka

import (
	"context"
	"fmt"
	"strings"

	"github.com/rashiraffi/streamease/internal/config"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

func Publish(topic string, header map[string]string) {
	w := &kafka.Writer{
		Addr:        kafka.TCP(config.KafkaHost),
		Topic:       topic,
		Balancer:    &kafka.LeastBytes{},
		MaxAttempts: 3,
	}

	messageHeaders := []protocol.Header{}
	for k, v := range header {
		messageHeaders = append(messageHeaders, protocol.Header{
			Key:   k,
			Value: []byte(v),
		})
	}
	fmt.Println("\033[2J")
	fmt.Printf("Publishing to topic: %s, Input 'Exit' to exit\n", topic)
	for {
		var input string
		fmt.Scanln(&input)
		if strings.ToLower(input) == "exit" {
			break
		}

		// jsonMessage, err := json.Marshal(input)
		// if err != nil {
		// 	fmt.Println("ERROR:\t\tInvalid message. Please try again")
		// 	continue
		// }

		err := w.WriteMessages(
			context.Background(),
			kafka.Message{
				Key:     []byte("Key-A"),
				Value:   []byte(input),
				Headers: messageHeaders,
			},
		)

		if err != nil {
			fmt.Println("\033[2J")
			log.Error().Err(err).Msg("Error publishing message")
			return
		}
	}

}
