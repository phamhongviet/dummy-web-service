package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var port *string
var name *string

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, *name)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func environHandler(w http.ResponseWriter, r *http.Request) {
	var env string
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "DUMMY_") {
			env += e + "\n"
		}
	}
	fmt.Fprintln(w, env)
}

func init() {
	port = flag.String("p", "8888", "listen port")
	name = flag.String("n", "Dummy Web Service", "service name")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/environ", environHandler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":"+*port, nil)
}
