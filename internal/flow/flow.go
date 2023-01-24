package flow

import (
	"fmt"
	"os"
)

func InitialFlow() {
	for {
		fmt.Println("Select an option:")
		fmt.Println("1. List All Topics\n2. Create a Topic\n3. Produce to a Topic\n4. Consume from a Topic\n5. Exit")
		fmt.Print("Enter your option: ")

		selectedOption := ""

		fmt.Scan(&selectedOption)

		switch selectedOption {
		case "1":
			fmt.Print("\033[H\033[2J")
			listAllTopics()
		case "2":
			fmt.Println("Create a Topic")
			fmt.Print("\033[H\033[2J")
		case "3":
			PublishFlow()
			fmt.Print("\033[H\033[2J")
		case "4":
			fmt.Println("Consume from a Topic")
			fmt.Print("\033[H\033[2J")
		case "5":
			fmt.Print("\033[H\033[2J")
			os.Exit(1)
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Invalid Option, please try again.")
		}
	}
}
