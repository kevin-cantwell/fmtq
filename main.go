package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	useSingleQuotes := len(os.Args) > 1 && os.Args[1] == "-s"
	if useSingleQuotes {
		os.Stdout.Write([]byte(`'`))
		defer os.Stdout.Write([]byte("'\n"))
	} else {
		os.Stdout.Write([]byte(`"`))
		defer os.Stdout.Write([]byte("\"\n"))
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			return
		}
		var escaped []byte
		if useSingleQuotes {
			escaped = []byte(strconv.QuoteRune(r))
		} else {
			escaped = []byte(strconv.Quote(string(r)))
		}
		if _, err := os.Stdout.Write(escaped[1 : len(escaped)-1]); err != nil {
			log.Fatal(err)
		}
	}
}
