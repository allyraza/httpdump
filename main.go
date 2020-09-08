package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))

		fmt.Printf("%s %s\n\n", r.Method, r.URL)
		fmt.Println("HEADERS")
		for k := range r.Header {
			fmt.Printf("%s: %s\n", k, r.Header.Get(k))
		}

		fmt.Println("")
	})

	http.ListenAndServe("localhost:2000", mux)
}
