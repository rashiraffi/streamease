package flow

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/rashiraffi/streamease/internal/kafka"
)

func PublishFlow() {
	selectedTopic := ""
	topics := kafka.GetAllTopics()
	for {
		fmt.Println("Select a topic from the list below to publish to:")
		for i := 0; i < len(topics); i++ {
			fmt.Printf("%d. %s\n", i+1, topics[i])
		}
		fmt.Print("Enter the topic number / Input 'Exit' to exit : ")
		selectedTopic = ""
		fmt.Scan(&selectedTopic)

		if strings.ToLower(selectedTopic) == "exit" {
			return
		}

		selectedTopicint, err := strconv.Atoi(selectedTopic)
		if err != nil || selectedTopicint > len(topics) || selectedTopicint < 1 {
			fmt.Println("\033[2J")
			fmt.Println("Invalid Option, please try again.")
			continue
		}
		selectedTopic = topics[selectedTopicint-1]
		break
	}
	fmt.Println("\033[2J")
	publish(selectedTopic)
}

func publish(topic string) {
	var headers map[string]string = nil
	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Set Headers and Publish\n2. Publish\n3. Exit")
		fmt.Print("Enter your option: ")
		selectedOption := ""
		fmt.Scan(&selectedOption)
		switch selectedOption {
		case "1":
			headerString := ""
			fmt.Println("\033[2J")
			fmt.Println("Enter headers in the json format eg: {\"key1\":\"value1\",\"key2\":\"value2\"}")
			fmt.Scan(&headerString)

			err := json.Unmarshal([]byte(headerString), &headers)
			if err != nil {
				fmt.Println("\033[2J")
				fmt.Println("Invalid headers, please try again.")
			}

			fmt.Println("\033[2J")
			kafka.Publish(topic, headers)
		case "2":
			kafka.Publish(topic, headers)
		case "3":
			fmt.Println("\033[2J")
			return
		default:
			fmt.Println("\033[2J")
			fmt.Println("Invalid Option, please try again.")
		}
	}
}
