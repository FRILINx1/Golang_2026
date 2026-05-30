package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

type Message struct {
	Value int  `json:"value"`
	Done  bool `json:"done"`
}

func main() {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.Subscribe("pipeline.numbers", func(m *nats.Msg) {
		var msg Message
		json.Unmarshal(m.Data, &msg)

		if msg.Done {
			nc.Publish("pipeline.even", m.Data)
			return
		}

		if msg.Value%2 == 0 {
			log.Printf("Пропущено  число: %d", msg.Value)
			nc.Publish("pipeline.even", m.Data)
		}
	})

	log.Println("Filter чекає.")
	select {} 
}