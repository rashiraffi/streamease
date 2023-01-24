package main

import (
	"github.com/rashiraffi/streamease/internal/flow"
	"github.com/rashiraffi/streamease/internal/kafka"
)

func main() {

	kafka.InitClient("")
	kafka.Ping()

	flow.InitialFlow()

}
