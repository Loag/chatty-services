package broadcaster

import (
	"fmt"
	"net"
	"os"
	"slices"
	"time"
)

func Start() {
	broadcastAddress := "255.255.255.255:8000"

	conn, err := net.Dial("udp", broadcastAddress)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	hostname, err := os.Hostname()
	// Infinite loop to keep sending broadcast messages
	for {
		message := []byte(hostname + ":8080")
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		time.Sleep(2 * time.Second)
	}
}

func Listen(servers *[]string) {
	address := ":8000"
	hostname, err := os.Hostname()

	// Create UDP Address
	udpAddress, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	// Infinite loop to keep listening for messages
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Received %s from %s\n", string(buf[:n]), addr)

		s := string(buf[:n])
		if !slices.Contains(*servers, s) && s != (hostname+":8080") {
			*servers = append(*servers, s)
		}
	}
}
