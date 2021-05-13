package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shakursmith/goserver/controller"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/distance", controller.DistanceHandler).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
