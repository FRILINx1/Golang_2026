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

	var totalSum int

	nc.Subscribe("pipeline.squared", func(m *nats.Msg) {
		var msg Message
		json.Unmarshal(m.Data, &msg)

		if msg.Done {
			log.Printf("Сума чисел = %d", totalSum)
			return
		}

		totalSum += msg.Value
		log.Printf("Додано %d. Теперішня сума: %d", msg.Value, totalSum)
	})

	log.Println("Sum чекає")
	select {}
}