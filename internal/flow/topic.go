package flow

import (
	"fmt"

	"github.com/rashiraffi/streamease/internal/kafka"
)

func listAllTopics() {
	topics := kafka.GetAllTopics()
	for _, topic := range topics {
		fmt.Println(topic)
	}
	fmt.Println()
}
