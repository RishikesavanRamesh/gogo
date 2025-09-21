package main

import (
	// "bytes"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {

	out := make(chan string, 1)
	go func() {
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

	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	lines := getLinesChannel(f)

	for line := range lines {
		fmt.Println("read:", line)
	}

}
