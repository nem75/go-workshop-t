package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const filename = "kv.db"
const delim = "="

func main() {
	file, err := os.Open(filename)

	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(filename)
			if err != nil {
				fmt.Printf("error creating file %s: %v\n", filename, err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("error opening file %s: %v\n", filename, err)
			os.Exit(2)
		}
	}

	defer file.Close()

	if len(os.Args) < 2 {
		// Just output the file.
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("error reading contents of file %s: %v\n", filename, err)
			os.Exit(3)
		}

		fmt.Printf("%s", contents)

	} else {
		// Read db file into memory.
		var db = make(map[string]string)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			split := strings.Split(scanner.Text(), delim)
			db[split[0]] = split[1]
		}

		// Check if read or write.
		var writeArg = regexp.MustCompile(`^[^=]+=[^=]*$`)
		writeMode := writeArg.MatchString(os.Args[1])

		if writeMode {
			// Write data from args into map.
			for _, arg := range os.Args[1:] {
				split := strings.Split(arg, "=")
				db[split[0]] = split[1]
			}

			// Rewrite file.
			var data string
			file.Close()
			os.Truncate(filename, 0)
			for key, value := range db {
				data += key + delim + value + "\n"
			}
			ioutil.WriteFile(filename, []byte(data), 0666)

		} else {
			// Output data for given key.
			for _, arg := range os.Args[1:] {
				fmt.Println(db[arg])
			}
		}
	}
}
