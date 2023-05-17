package kafka

import (
	"fmt"

	"github.com/rashiraffi/streamease/internal/kafka"
	"github.com/spf13/cobra"
)

var (
	createTopic bool
	listTopic   bool
)

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Manage Kafka topics",
	Long: `The topic command allows you to create and list Kafka topics. 
	Create a new topic by specifying the name with the --create flag. 
	List all topics with the --list flag. 
	For example:

	streamease kafka topic --create my-topic
	streamease kafka topic --list`,
	Run: func(cmd *cobra.Command, args []string) {
		if !createTopic && !listTopic {
			cmd.Help()
			return
		}

		hostAddr, err := getBrokerAddr(broker, profile)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		client := kafka.New(hostAddr)

		if createTopic {
			if len(args) != 1 {
				fmt.Println("Error: topic name is required")
				return
			}
		}
		if listTopic {
			topics := client.GetAllTopics()
			for _, topic := range topics {
				fmt.Println(topic)
			}
		}
	},
}

func init() {

	topicCmd.Flags().StringVarP(&profile, "profile", "p", "", "profile name")
	topicCmd.Flags().StringVarP(&broker, "broker", "b", "", "kafka broker address")

	topicCmd.Flags().BoolVarP(&createTopic, "create", "c", false, "create topic with name")
	topicCmd.Flags().BoolVarP(&listTopic, "list", "l", false, "list topic")

	KafkaCmd.AddCommand(topicCmd)

}
