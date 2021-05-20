package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/shakursmith/goserver/middleware"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/distance", middleware.DistanceHandler).Methods("POST")
	r.HandleFunc("/hello", middleware.HelloHandler).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Starting server on the port 8080...")

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
