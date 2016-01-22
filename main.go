package main

import (
	"flag"
	"fmt"
	"net/http"
)

var PORT *string
var NAME *string

func root_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, *NAME)
}

func health_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func init() {
	PORT = flag.String("p", "8888", "listen port")
	NAME = flag.String("n", "Dummy Web Service", "service name")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", root_handler)
	http.HandleFunc("/health", health_handler)
	http.ListenAndServe(":"+*PORT, nil)
}
