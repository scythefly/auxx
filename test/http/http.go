package http

import (
	"fmt"
	"net/http"
)

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("--- handle redirect\n")

	fmt.Println("request.Host:", r.Host)
	fmt.Println("request.Header.Get(\"Host\"):", r.Header.Get("Host"))
	fmt.Println("request.URL.Host:", r.URL.Host)
	fmt.Println("request.URL.Hostname():", r.URL.Hostname())

	w.WriteHeader(http.StatusNotFound)
}
