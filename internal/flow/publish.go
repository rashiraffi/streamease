package flow

import "fmt"

func PublishFlow() {
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Select an option:")
		fmt.Println("1. Publish to a Topic\n2. Publish to a Topic with Key\n3. Publish to a Topic with Key and Partition\n4. Exit")
		fmt.Print("Enter your option: ")
		selectedOption := ""
		fmt.Scan(&selectedOption)
		switch selectedOption {
		case "1":
			fmt.Println("Publish to a Topic")
		case "2":
			fmt.Println("Publish to a Topic with Key")
		case "3":
			fmt.Println("Publish to a Topic with Key and Partition")
		case "4":
			return
		default:
			fmt.Println("Invalid Option")
		}
	}
}
