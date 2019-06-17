package main

import (
	"fmt"
	"net/http"
)

type response string

func (res response) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "This is hello world")
}

func main() {
	var res response
	http.ListenAndServe(":8080", res)
}
