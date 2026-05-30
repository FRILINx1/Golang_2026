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

	nc.Subscribe("pipeline.even", func(m *nats.Msg) {
		var msg Message
		json.Unmarshal(m.Data, &msg)

		if msg.Done {
			nc.Publish("pipeline.squared", m.Data)
			return
		}

		squaredValue := msg.Value * msg.Value
		log.Printf("%d^2 = %d", msg.Value, squaredValue)
		
		msg.Value = squaredValue
		data, _ := json.Marshal(msg)
		nc.Publish("pipeline.squared", data)
	})

	log.Println("Square чекає")
	select {}
}