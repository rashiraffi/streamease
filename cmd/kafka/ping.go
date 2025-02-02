/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package kafka

import (
	"fmt"

	"github.com/rashiraffi/streamease/internal/kafka"
	"github.com/spf13/cobra"
)

var (
	urlPath string
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check if the Kafka broker is up and running",
	Long:  `The ping command checks if the Kafka broker is up and running by sending a ping request to the broker. If the broker is up and running, the command returns a success message. Otherwise, it returns an error message.`,
	Run: func(cmd *cobra.Command, args []string) {
		hostAddr, err := getBrokerAddr(broker, profile)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		client := kafka.New(hostAddr)
		err = client.Ping()
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Kafka broker is up and running")
		}

	},
}

func init() {

	pingCmd.Flags().StringVarP(&profile, "profile", "p", "", "profile name")
	pingCmd.Flags().StringVarP(&broker, "broker", "b", "", "kafka broker address")

	KafkaCmd.AddCommand(pingCmd)

}
