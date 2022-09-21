package main

import (
	"Level0/server"
	"sync"
)

func init() {
	server.СacheFromDatabase()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go server.StartServer()

	server.ListenToNATSStreaming()

	wg.Wait()
}
