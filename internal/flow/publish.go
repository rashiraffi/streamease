package flow

// func PublishFlow() {
// 	selectedTopic := ""
// 	topics := kafka.GetAllTopics()
// 	for {
// 		fmt.Println("Select a topic from the list below to publish to:")
// 		for i := 0; i < len(topics); i++ {
// 			fmt.Printf("%d.\t%s\n", i+1, topics[i])
// 		}
// 		fmt.Print("Enter the topic number / Input 'Exit' to exit : ")
// 		selectedTopic = ""
// 		fmt.Scan(&selectedTopic)

// 		if strings.ToLower(selectedTopic) == "exit" {
// 			clear()
// 			return
// 		}

// 		selectedTopicint, err := strconv.Atoi(selectedTopic)
// 		if err != nil || selectedTopicint > len(topics) || selectedTopicint < 1 {
// 			clear()
// 			fmt.Println("ERROR :: Invalid Option, please try again.")
// 			continue
// 		}
// 		selectedTopic = topics[selectedTopicint-1]
// 		break
// 	}
// 	clear()
// 	publish(selectedTopic)
// }

// func publish(topic string) {
// 	var headers map[string]string = nil
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for {
// 		fmt.Println("Select an option:")
// 		fmt.Println("1. Set Headers and Publish\n2. Publish\n3. Exit")
// 		fmt.Print("Enter your option: ")
// 		selectedOption := ""
// 		fmt.Scan(&selectedOption)
// 		switch selectedOption {
// 		case "1":
// 			headerString := ""
// 			clear()

// 			fmt.Println("Enter headers in the json format eg: {\"key1\":\"value1\",\"key2\":\"value2\"}")

// 			scanner.Scan()
// 			headerString = scanner.Text()

// 			err := json.Unmarshal([]byte(headerString), &headers)
// 			if err != nil {
// 				clear()
// 				fmt.Println("ERROR :: Invalid headers, please try again.")
// 				continue
// 			}

// 			clear()
// 			kafka.Publish(topic, headers)
// 		case "2":
// 			clear()
// 			kafka.Publish(topic, headers)
// 		case "3":
// 			clear()
// 			return
// 		default:
// 			clear()
// 			fmt.Println("ERROR :: Invalid Option, please try again.")
// 		}
// 	}
// }
