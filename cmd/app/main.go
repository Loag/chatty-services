package main

import (
	"encoding/json"
	"gossip/internal/broadcaster"
	"gossip/internal/server"
	"log"
	"net/http"
)

type res struct {
	Id   string `json:"id"`
	Time string `json:"time"`
}

func main() {

	discovered_servers := []string{}

	go broadcaster.Start()
	go broadcaster.Listen(&discovered_servers)

	go server.Start()

	for {
		for _, server := range discovered_servers {
			resp, err := http.Get(server)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			var r res
			if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
				log.Fatal(err)
			}

			log.Printf("Time received: %s", r.Time)
		}
	}
}
