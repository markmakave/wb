package main

import (
	"fmt"
	"io/ioutil"
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

	// file content buffer
	var fileContent [3][]byte

	// Some garbage
	fileContent[0] = []byte("Hello World!")

	fileContent[1], err = ioutil.ReadFile("json/model.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileContent[2], err = ioutil.ReadFile("json/order1.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		index := time.Now().UnixNano() % 3

		nc.Publish("orders", fileContent[index])
		nc.Flush()

		// wait 1 second
		time.Sleep(1 * time.Second)
	}

}
