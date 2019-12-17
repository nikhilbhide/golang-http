package main

import (
	"github.com/gorilla/mux"
	"github.com/nik/JWTDemo/services/registry"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signupv1",registry.SignUpHandlerV1).Methods("POST")
	router.HandleFunc("/signupv2",registry.SignUpHandlerV2).Methods("POST")
	router.HandleFunc("/signupv3",registry.SignUpHandlerV3).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}