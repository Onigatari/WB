package main

import (
	"github.com/nats-io/stan.go"
	"io"
	"log"
	"os"
)

func main() {
	sc, _ := stan.Connect("test-cluster", "publisher")
	defer sc.Close()

	//file, _ := os.Open("./models/model.json")
	//data, _ := io.ReadAll(file)
	//_ = sc.Publish("json_channel", data)

	file, _ := os.Open("./models/test.json")
	data, _ := io.ReadAll(file)
	_ = sc.Publish("json_channel", data)

	log.Println("Publish was successful!")
}
