package main

import "net/http"
import "fmt"
import "time"

func getResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("MyKey", "This is my key")
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintf(w, "<h1>hello, you've got a specific handler for a %s\n</h1>", r.URL.Path)
}

func main() {
	handler := http.HandlerFunc(getResponse)
	s := &http.Server{
		Addr:           ":8082",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.Handle("/mypath/newpath", handler)

	s.ListenAndServe()
}
