package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	target := "127.0.0.1"
	// Define the range of ports you want to scan
	startPort := 1
	endPort := 1024

	fmt.Printf("Scanning ports on %s...\n", target)

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", target, p)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// Port is closed or filtered
				// fmt.Printf("Port %d is close\n", p)
				return
			}
			defer conn.Close()
			fmt.Printf("Port %d is open\n", p)
		}(port)
	}

	wg.Wait()
}
