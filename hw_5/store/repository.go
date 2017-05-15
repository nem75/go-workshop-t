package store

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// FileRepository is a concrete implementation of a key value store data
// repository in the file system.
type FileRepository struct {
	Filename string
}

// NewFileRepository returns a new FileRepository.
// Uses the given filename and path to actually store/read stuff.
func NewFileRepository(filename string) FileRepository {
	return FileRepository{filename}
}

// ReadDb returns the data found in the repository's physical file.
func (f FileRepository) ReadDb() Data {
	file, err := os.Open(f.Filename)

	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(f.Filename)
			if err != nil {
				fmt.Printf("error creating file %s: %v\n", f.Filename, err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("error opening file %s: %v\n", f.Filename, err)
			os.Exit(2)
		}
	}
	defer file.Close()

	db := Data{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "=")
		db[split[0]] = split[1]
	}

	return db
}

// WriteDb writes the given Data struct to the repository's physical file.
func (f FileRepository) WriteDb(db Data) {
	var data string

	err := os.Truncate(f.Filename, 0)
	if err != nil {
		fmt.Printf("error truncating file %s: %v\n", f.Filename, err)
		os.Exit(3)
	}

	for key, value := range db {
		data += key + "=" + value + "\n"
	}

	err = ioutil.WriteFile(f.Filename, []byte(data), 0666)
	if err != nil {
		fmt.Printf("error writing file %s: %v\n", f.Filename, err)
		os.Exit(4)
	}
}

// virtualFile is used in testing the Store with a VirtualRepository.
var virtualFile = Data{}

// VirtualRepository is an implementation of a key value store data repository
// in memory. Used for testing.
type VirtualRepository struct {
}

// NewVirtualRespository returns a new VirtualRepository.
func NewVirtualRespository() VirtualRepository {
	return VirtualRepository{}
}

// ReadDb returns the contents of the current virtualFile.
func (v VirtualRepository) ReadDb() (db Data) {
	return virtualFile
}

// WriteDb writes the given Data struct to the virtualFile.
func (v VirtualRepository) WriteDb(db Data) {
	virtualFile = db
}
