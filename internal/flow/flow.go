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
			fmt.Println("\033[2J")
			listAllTopics()
		case "2":
			fmt.Println("Create a Topic")
			fmt.Println("\033[2J")
		case "3":
			fmt.Println("\033[2J")
			PublishFlow()
			fmt.Println("\033[2J")
		case "4":
			fmt.Println("Consume from a Topic")
			fmt.Println("\033[2J")
		case "5":
			fmt.Println("\033[2J")
			os.Exit(1)
		default:
			fmt.Println("\033[2J")
			fmt.Println("Invalid Option, please try again.")
		}
	}
}
