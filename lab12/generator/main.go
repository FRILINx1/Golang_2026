package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

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

	time.Sleep(3 * time.Second)

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal("Помилка підключення NATS: ", err)
	}
	defer nc.Close()

	log.Println("Генератор запущено")

	for i := 1; i <= 100; i++ {
		msg := Message{Value: i, Done: false}
		data, _ := json.Marshal(msg)
		
		nc.Publish("pipeline.numbers", data)
		log.Printf("Згенеровано: %d", i)
		
		time.Sleep(10 * time.Millisecond) 
	}

	doneMsg := Message{Done: true}
	data, _ := json.Marshal(doneMsg)
	nc.Publish("pipeline.numbers", data)
	
	log.Println("Генерацію завершено.")
	time.Sleep(1 * time.Second)
}