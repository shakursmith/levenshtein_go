package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shakursmith/goserver/middleware"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/distance", middleware.DistanceHandler).Methods("POST")
	r.HandleFunc("/hello", middleware.HelloHandler).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
