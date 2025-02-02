package kafka

type kafkaClient struct {
	HostAddr string
}

type KafkaClient interface {
	Ping() error
	GetAllTopics() []string
}

func New(hostAddr string) KafkaClient {
	return &kafkaClient{
		HostAddr: hostAddr,
	}
}
