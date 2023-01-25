package flow

import (
	"fmt"
	"os"
)

func InteractiveFlow() {
	for {
		fmt.Println("Select an option:")
		fmt.Println("1. List All Topics\n2. Create a Topic\n3. Produce to a Topic\n4. Consume from a Topic\n5. Exit")
		fmt.Print("Enter your option: ")

		selectedOption := ""

		fmt.Scan(&selectedOption)

		switch selectedOption {
		case "1":
			clear()
			listAllTopics()
		case "2":
			clear()
			fmt.Println("Create a Topic")
		case "3":
			clear()
			PublishFlow()
		case "4":
			clear()
			fmt.Println("Consume from a Topic")
		case "5":
			clear()
			os.Exit(0)
		default:
			clear()
			fmt.Println("Invalid Option, please try again.")
		}
	}
}
