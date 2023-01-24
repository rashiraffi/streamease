package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rashiraffi/streamease/internal/flow"
	"github.com/rashiraffi/streamease/internal/kafka"
)

func main() {

	var (
		args []string
		host string = ""
		mode string = ""
	)

	args = os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-host":
			host = args[i+1]
			i++
		case "-mode":
			mode = args[i+1]
		}
	}

	kafka.InitClient(host)
	kafka.Ping()

	switch strings.ToLower(mode) {
	case "interactive":
		flow.InteractiveFlow()
	case "producer":
		flow.PublishFlow()
	case "consumer":
		fmt.Println("Consumer mode not implemented yet.")
	default:
		flow.InteractiveFlow()
	}

}
