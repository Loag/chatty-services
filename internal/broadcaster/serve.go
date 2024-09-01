package broadcaster

import (
	"fmt"
	"net"
	"os"
	"slices"
	"sync"
	"time"
)

func Start(wg *sync.WaitGroup) {
	broadcastAddress := "255.255.255.255:8000"

	conn, err := net.Dial("udp", broadcastAddress)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	hostname, _ := os.Hostname()

	i := 0
	for {
		i++
		message := []byte(hostname + ":8080")
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		time.Sleep(2 * time.Second)

		if i == 5 {
			break
		}
	}
	wg.Done()
}

func Listen(servers *[]string, wg *sync.WaitGroup) {
	address := ":8000"
	hostname, _ := os.Hostname()

	udpAddress, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	i := 0

	for {
		i++
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Received %s from %s\n", string(buf[:n]), addr)

		s := string(buf[:n])
		// check that it is not the sender
		if !slices.Contains(*servers, s) && s != (hostname+":8080") {
			*servers = append(*servers, s)
		}

		if i == 50 {
			break
		}
	}

	wg.Done()
}
