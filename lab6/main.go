package main

import (
	"encoding/json"
	"fmt"
)


func ToJSON(v any) (string, error) {

	jsonData, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return "", err
	}
	
	return string(jsonData), nil
}

type Server struct {
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	Debug      bool     `json:"debug"`
	AllowedIPs []string `json:"allowed_ips"`
}

func main() {

	myServer := Server{
		Host:       "localhost",
		Port:       8080,
		Debug:      true,
		AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
	}


	jsonResult, err := ToJSON(myServer)
	if err != nil {
		fmt.Println("Помилка :", err)
		return
	}


	fmt.Println(jsonResult)
}