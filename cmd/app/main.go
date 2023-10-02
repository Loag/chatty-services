package main

import (
	"encoding/json"
	"gossip/internal/broadcaster"
	"gossip/internal/server"
	"log"
	"net/http"
	"sync"
	"time"
)

type res struct {
	Id   string `json:"id"`
	Time string `json:"time"`
}

func main() {

	discovered_servers := []string{}

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go broadcaster.Start(&wg)

	wg.Add(1)
	go broadcaster.Listen(&discovered_servers, &wg)

	wg.Add(1)
	go server.Start(&wg)

	for {

		for _, server := range discovered_servers {

			resp, err := http.Get("http://" + server + "/")
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			var r res
			if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
				log.Fatal(err)
			}

			log.Printf("Server: %s Time received: %s", server, r.Time)

			time.Sleep(5 * time.Second)
		}
	}
}
