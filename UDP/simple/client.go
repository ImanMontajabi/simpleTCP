package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello UDP Server!"))
	if err != nil {
		fmt.Println("Failed to send message:", err)
		return
	}
	fmt.Println("Message sent to server.")

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}
	fmt.Printf("Response from the server: %s\n", string(buffer[:n]))
}
