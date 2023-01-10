package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) (int, error) {
	scanner := bufio.NewScanner(r)
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	wc := 0

	for scanner.Scan() {
		wc++
		if err := scanner.Err(); err != nil {
			return 0, err
		}
	}
	return wc, nil
}
