package main

import "net/http"
import "fmt"
import "time"

func handleSpecificPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, you've got a specific handler for a %s\n", r.URL.Path)
}

func main() {
	handler := http.HandlerFunc(handleSpecificPath)
	s := &http.Server{
		Addr:           ":8082",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.Handle("/mypath/newpath", handler)

	s.ListenAndServe()
}
