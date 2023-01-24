package flow

import (
	"fmt"

	"github.com/rashiraffi/streamease/internal/kafka"
)

func listAllTopics() {
	topics := kafka.GetAllTopics()
	fmt.Println()
	for _, topic := range topics {
		if topic == "__consumer_offsets" {
			continue
		}
		fmt.Println(topic)
	}
	fmt.Println()
}
