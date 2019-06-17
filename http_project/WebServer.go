package main

import "net/http"
import "fmt"
import "time"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
}

func main() {
	handler := http.HandlerFunc(handleRequest)
	httpServerInstace := &http.Server{
		Addr:           ":8081",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	httpServerInstace.ListenAndServe()
}
