package main

import (
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", root)
	rtr.HandleFunc("/home", worldwidestats).Methods("GET")
	rtr.HandleFunc("/all", all).Methods("GET")
	rtr.HandleFunc("/countrynames", countrynames).Methods("GET")
	rtr.HandleFunc("/country/{countryname}", countrywise).Methods("GET")
	http.Handle("/", rtr)
	http.ListenAndServe(":8000", nil)
}
