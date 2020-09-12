package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/allyraza/httpdump"
)

func main() {
	addr := flag.String("addr", ":8080", "address to listen on")
	flag.Parse()

	mux := http.NewServeMux()

	mux.Handle("/", httpdump.RequestDumpHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	log.Printf("started server on addr %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}
