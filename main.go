package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/arturoverbel/microservice_compra/connection"
	"github.com/arturoverbel/microservice_compra/model"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)


var prefixPath = "/api/compra"

func main() {
	r := mux.NewRouter()

	r.HandleFunc(prefixPath, CreateShoppingController).Methods("POST")
	r.HandleFunc(prefixPath, UpdateShoppingController).Methods("PUT")
	r.HandleFunc(prefixPath, DeleteShoppingController).Methods("DELETE")
	r.HandleFunc(prefixPath+"/{id}", FindShoppingByIDController).Methods("GET")
	r.HandleFunc(prefixPath+"/by-user/{id_user}", FindShoppingByUserController).Methods("GET")

	log.Printf("Listening...")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}