package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var user = "userxyz"
var password = ",q:=jW>,nG8#NTA"
var dbname = "coronadb"

type allstats struct {
	ActiveCases    int `json:"ActiveCases"`
	TotalRecovered int `json:"TotalRecovered"`
	TotalDeaths    int `json:"TotalDeaths"`
	TotalCases     int `json:"TotalCases"`
}

type data struct {
	Name  string   `json:"Name"`
	Stats allstats `json:"Stats"`
}

type worldwideData struct {
	Stats        allstats `json:"Stats"`
	AllCountries []data   `json:"AllCountries"`
}

func getDB() *sql.DB {
	db, err := sql.Open("mysql", user+":"+password+"@/"+dbname)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getcountrynames() *[]string {
	db := getDB()
	// defer the close till after the function has finished executing
	defer db.Close()
	var countrynames []string
	countries, err := db.Query("SELECT DISTINCT CountryName FROM CountryWiseData")
	for countries.Next() {
		var cname string
		err = countries.Scan(&cname)
		if err != nil {
			panic(err.Error())
		}
		countrynames = append(countrynames, cname)
	}
	return &countrynames
}

func getalldata() *worldwideData {

	db := getDB()
	// defer the close till after the function has finished executing
	defer db.Close()

	var result worldwideData
	var count int
	err := db.QueryRow(`SELECT COUNT(DISTINCT CountryName) From CountryWiseData`).Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	query := fmt.Sprintf(`SELECT CountryName, SUM(ActiveCases), SUM(TotalRecovered), SUM(TotalDeaths) 
							FROM CountryWiseData GROUP BY Date, CountryName ORDER BY Date DESC LIMIT %d`, count)
	countries, err := db.Query(query)

	for countries.Next() {
		var currentdata data
		err = countries.Scan(
			&currentdata.Name,
			&currentdata.Stats.ActiveCases,
			&currentdata.Stats.TotalRecovered,
			&currentdata.Stats.TotalDeaths)
		if err != nil {
			panic(err.Error())
		}
		result.AllCountries = append(result.AllCountries, currentdata)
	}
	return &result

}

func getcountrydata(country string) *data {

	db := getDB()
	// defer the close till after the function has finished executing
	defer db.Close()

	var result data
	result.Name = country
	query := fmt.Sprintf(`SELECT SUM(ActiveCases), SUM(TotalRecovered), SUM(TotalDeaths) 
							FROM CountryWiseData Where CountryName='%s' GROUP BY Date ORDER BY Date DESC LIMIT 1;`, country)
	err := db.QueryRow(query).Scan(&result.Stats.ActiveCases, &result.Stats.TotalRecovered, &result.Stats.TotalDeaths)
	if err != nil {
		panic(err.Error())
	}
	return &result
}
