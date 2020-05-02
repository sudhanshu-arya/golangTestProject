package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Corona API")

}
func worldwidestats(w http.ResponseWriter, r *http.Request) {
	var result data
	result.Name = "Worldwide"
	result.Stats = worldwideservice().Stats
	// returning json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func all(w http.ResponseWriter, r *http.Request) {
	result := worldwideservice()
	// returning json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func countrywise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := vars["countryname"]
	result := countryservice(country)
	// returning json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func countrynames(w http.ResponseWriter, r *http.Request) {
	result := nameservice()
	// returning json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
