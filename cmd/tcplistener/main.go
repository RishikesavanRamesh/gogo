package main

import (
	// "bytes"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	// "os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {

	out := make(chan string, 1)
	go func() {
		defer fmt.Println("closing channel")
		defer close(out)
		defer f.Close()

		var str string = ""
		for {
			data := make([]byte, 8)

			n, err := f.Read(data)

			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i >= 0 {
				str += string(data[:i])
				data = data[i+1:]
				out <- str
				str = ""
			}

			str = str + string(data)
			if err == io.EOF {
				if len(str) > 0 {
					out <- str // send the last line
				}
				return
			}
		}
	}()

	return out
}

func main() {

	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}else {
			fmt.Println("accepted connection from", conn.RemoteAddr())
		}
		
		lines := getLinesChannel(conn)

		for line := range lines {
			fmt.Println(line)
		}
	}

}
