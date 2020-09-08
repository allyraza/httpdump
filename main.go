package main

import (
	"fmt"
	"net/http"
)

func headers(r *http.Request) string {
	headers := ""
	for k := range r.Header {
		headers += fmt.Sprintf("%s: %s\n", k, r.Header.Get(k))
	}

	return headers
}

// RequestDumpHandler dumps request to stdout
func RequestDumpHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n\n%s\n", r.Method, r.URL, headers(r))
	})
}

// RequestDumpHandlerFunc dumps request to stdout
func RequestDumpHandlerFunc(next http.HandlerFunc) http.Handler {
	return RequestDumpHandler(http.HandlerFunc(next))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", RequestDumpHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))

	http.ListenAndServe("localhost:2000", mux)
}
