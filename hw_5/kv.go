package main

import (
	"fmt"
	st "go-workshop-t/hw_4/store"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var kvs *st.Store

func main() {
	args := os.Args[1:]
	store := st.NewStore(st.NewFileRepository("./kv.db"))

	switch getMode(args) {
	case "readAll":
		printAll(store)
	case "read":
		printValues(store, args)
	case "write":
		defer store.Flush()
		storeValues(store, args)
	case "daemon":
		kvs = store
		runServer()
	}
}

// Determine whether to read or write for the current run.
func getMode(args []string) (mode string) {
	mode = "read"

	if len(args) < 1 {
		mode = "readAll"
	} else if args[0] == "-d" {
		mode = "daemon"
	} else {
		writeArg := regexp.MustCompile(`^[^=]+=[^=]*$`)
		writeMode := writeArg.MatchString(args[0])
		if writeMode {
			mode = "write"
		}
	}

	return
}

// Split write parameters into key/value parts.
func splitWriteArg(arg string) (splitted [2]string) {
	split := strings.Split(arg, "=")
	splitted[0] = split[0]
	splitted[1] = split[1]
	return
}

// Print all key/values from the store.
func printAll(s *st.Store) {
	fmt.Println(s.GetAll())
}

// Print the values for the given keys.
func printValues(s *st.Store, args []string) {
	for _, arg := range args {
		fmt.Println(s.Get(arg))
	}
}

// Save the given key/values in the store.
func storeValues(s *st.Store, args []string) {
	for _, arg := range args {
		splitted := splitWriteArg(arg)
		s.Set(splitted[0], splitted[1])
	}
}

func runServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", mux)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlePost(w, r)
	case "GET":
		handleGet(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Print(r.Method)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer kvs.Flush()
	kvs.Set(key, string(value))

	w.WriteHeader(http.StatusCreated)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	value := kvs.Get(key)

	if len(value) < 1 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		io.WriteString(w, value)
	}
}
