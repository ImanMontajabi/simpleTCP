package main

import (
	"fmt"
	"net"
)

func main() {
	pc, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer pc.Close()

	fmt.Println("UDP Server is listening on port 8080...")
	buffer := make([]byte, 1024)

	for {
		n, addr, err := pc.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading packet:", err)
			continue
		}
		fmt.Printf("Packet received from %s: %s\n", addr.String(), string(buffer[:n]))
		_, err = pc.WriteTo([]byte("Message received!"), addr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
}
