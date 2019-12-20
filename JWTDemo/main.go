package main
//go:generate swagger generate spec

import (
	"github.com/gorilla/mux"
	"github.com/nik/JWTDemo/bo"
	"github.com/nik/JWTDemo/driver/postgres"
	"github.com/nik/JWTDemo/repository/users"
	"github.com/nik/JWTDemo/services/registry"
	"log"
	"net/http"
	_ "github.com/nik/JWTDemo/docs"
)

func main() {
	router := mux.NewRouter()

	//instantiate handlers and dbinstance
	//for relevant handlers use JWTMiddleware
	db:= postgres.InitDB("postgres://nenyejsk:80PuDCmjzIAlMJb8xLaS7rgBwE7gXHeN@rajje.db.elephantsql.com:5432/nenyejsk")
	signUpRepo := users.NewLoginPostGresRepo(db)
	signUpUseCase:=bo.NewLoginUseCase(signUpRepo)


	router.HandleFunc("/signupv1",registry.SignUpHandlerV1).Methods("POST")
	router.HandleFunc("/signupv2",registry.SignUpHandlerV2).Methods("POST")
	router.HandleFunc("/signupv3",registry.SignUpHandlerV3).Methods("POST")
	router.HandleFunc("/signupv3",registry.SignUpHandlerV3).Methods("POST")
	registry.NewLoginHandler(router,signUpUseCase)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}