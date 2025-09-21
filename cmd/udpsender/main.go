package main

import (
	// "bytes"
	"bufio"
	// "bytes"
	"fmt"
	// "io"
	"log"
	"net"
	"os"
	// "os"
)


func main() {

	sender, err := net.ResolveUDPAddr("udp", ":42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	reader := bufio.NewReader(os.Stdin)

	conn, err := net.DialUDP("udp", nil, sender)
	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer conn.Close()

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error", "error", err)
		}
		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Fatal("error", "error", err)
		}
	}

}
