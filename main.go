package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	os.Stdout.Write([]byte(`"`))
	defer os.Stdout.Write([]byte("\"\n"))

	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			return
		}
		escaped := []byte(strconv.Quote(string(r)))
		if _, err := os.Stdout.Write(escaped[1 : len(escaped)-1]); err != nil {
			log.Fatal(err)
		}
	}
}
