package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer nc.Close()

	// read model.json file
	filename := "./model.json"
	plan, error := os.ReadFile(filename)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	// while true, publish message
	for {
		err = nc.Publish("channel", plan)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		nc.Flush()
		fmt.Println("Message published")

		time.Sleep(1 * time.Second)
	}

}
