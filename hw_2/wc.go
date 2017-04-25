package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Check arguments.
	if len(os.Args) < 2 || os.Args[1] == "-l" && len(os.Args) < 3 {
		fmt.Printf("usage: wc [-l] filename\n")
		os.Exit(1)
	}

	// Open file and create scanner.
	var filename string
	if os.Args[1] == "-l" {
		filename = os.Args[2]
	} else {
		filename = os.Args[1]
	}
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("error opening file %s: %v\n", filename, err)
		os.Exit(2)
	}

	scanner := bufio.NewScanner(file)

	// If -l is set, use ScanWords instead of ScanLines.
	if os.Args[1] != "-l" {
		scanner.Split(func(in []byte, eof bool) (pos int, s []byte, err error) {
			pos, s, err = bufio.ScanWords(in, eof)
			return
		})
	}

	var output uint
	for scanner.Scan() {
		output++
	}

	fmt.Println(output)
}
