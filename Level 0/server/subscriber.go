package server

import (
	"Level0/database"
	"Level0/parser"
	"github.com/nats-io/stan.go"

	"log"
	"sync"
)

func ListenToNATSStreaming() {
	sc, err := stan.Connect("test-cluster", "subscriber")
	if err != nil {
		log.Fatal(err)
	}

	defer sc.Close()

	log.Println("Nats streaming chanel listen")
	sc.Subscribe("json_channel", func(m *stan.Msg) {
		order := parser.ParseJsonByteArray(m.Data)
		database.InsertOrderToDB(order)
		DataBaseOrdersCache[order.OrderUID] = order
	})

	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
