package httpdump

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"log"
)

func formatHeaders(r *http.Request) string {
	headers := "\n\n"
	for k := range r.Header {
		headers += fmt.Sprintf("%s: %s\n", k, r.Header.Get(k))
	}

	return headers
}

func formatBody(r *http.Request) (string, error) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	if len(body) == 0 {
		return "", nil
	}

	return fmt.Sprintf("\n%s", string(body)), nil
}

// RequestDumpHandler dumps request to stdout
func RequestDumpHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := formatBody(r)
		if err != nil {
			log.Printf("body: invalid request body %s", err)
		}

		fmt.Printf("%s %s%s%s\n", r.Method, r.URL, formatHeaders(r), body)
	})
}

// RequestDumpHandlerFunc dumps request to stdout
func RequestDumpHandlerFunc(next http.HandlerFunc) http.Handler {
	return RequestDumpHandler(http.HandlerFunc(next))
}
