package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agnivade/levenshtein"
)

type Strings struct {
	S1 string
	S2 string
}

func DistanceHandler(w http.ResponseWriter, r *http.Request) {
	// declare new struct
	var s Strings

	// decode json
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&s)
	if err != nil {
		panic(err)
	}

	// calculate distance and respond
	distance := levenshtein.ComputeDistance(s.S1, s.S2)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(distance)
	fmt.Printf("The distance between %s and %s is %d.\n", s.S1, s.S2, distance)
}
