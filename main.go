package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var PORT *string
var NAME *string

func root_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, *NAME)
}

func health_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func environ_handler(w http.ResponseWriter, r *http.Request) {
	var env string
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "DUMMY_") {
			env += e
		}
	}
	fmt.Fprintf(w, env)
}

func init() {
	PORT = flag.String("p", "8888", "listen port")
	NAME = flag.String("n", "Dummy Web Service", "service name")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", root_handler)
	http.HandleFunc("/environ", environ_handler)
	http.HandleFunc("/health", health_handler)
	http.ListenAndServe(":"+*PORT, nil)
}
