package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
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
	var fileContent [5][]byte

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
	fileContent[3], err = ioutil.ReadFile("json/order2.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileContent[4], err = ioutil.ReadFile("json/order3.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	for {
		// generate random index using math/rand
		index := rand.Intn(5)

		nc.Publish("orders", fileContent[index])
		nc.Flush()
	}

}
