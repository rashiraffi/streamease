package kafka

import (
	"github.com/rashiraffi/streamease/internal/config"
	"github.com/segmentio/kafka-go"
)

func InitClient(hostAddr string) {
	if hostAddr == "" {
		config.KafkaHost = "localhost:9092"
	} else {
		config.KafkaHost = hostAddr
	}

}

func Ping() {
	conn, err := kafka.Dial("tcp", config.KafkaHost)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
}
