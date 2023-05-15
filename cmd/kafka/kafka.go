/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package kafka

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// KafkaCmd represents the kafka command
var KafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Kafka is a pallet that contains all kafka commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.AllKeys())
		fmt.Println(viper.Get("kafka"))
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
