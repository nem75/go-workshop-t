package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: tac filename [filename ...]\n")
		os.Exit(1)
	}

	// Not the best solution, all files are read into memory first.
	// Better: reading from EOF to last newline. Not trivial to detect though.
	var output []string

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s: %v\n", filename, err)
			os.Exit(2)
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			output = append([]string{scanner.Text()}, output...)
		}
	}

	for _, line := range output {
		fmt.Println(line)
	}
}
