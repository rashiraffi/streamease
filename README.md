# Welcome to StreamEase!

StreamEase is a Go executable package that makes it easy to interact with Apache Kafka. With StreamEase, you can easily list all topics in your Kafka cluster, publish messages to specific topics, and consume messages from selected topics.

## Features:

- List all topics in your Kafka cluster
- Publish messages to specific topics
- Consume messages from selected topics

## Installation:

- First, make sure you have Go 1.19 or later installed on your system.
- Run the command `go get github.com/rashiraffi/streamease@latest` to download and install the latest version of StreamEase.

## Running the program:

- To start the program, run `streamease`
- Additional tags that can be used:
  - `-host` to specify the hostname/IP address of your Kafka cluster
  - `-mode` to specify the program mode. It can be:
    - `interactive` to start with interactive mode
    - `consumer` to start with consumer mode
    - `producer` to start with producer mode

In interactive mode, you will be prompted to select the actions you want to perform.
In consumer mode, you will be prompted to select the topic from which you want to consume messages.
In producer mode, you will be prompted to select the topic to which you want to publish messages.

We hope you find StreamEase useful and enjoy using it!
