package main

func worldwideservice() *worldwideData {
	recievedData := getalldata()
	length := len(recievedData.AllCountries)
	for i := 0; i < length; i++ {
		currentCountryData := &recievedData.AllCountries[i]
		recievedData.Stats.ActiveCases += currentCountryData.Stats.ActiveCases
		recievedData.Stats.TotalRecovered += currentCountryData.Stats.TotalRecovered
		recievedData.Stats.TotalDeaths += currentCountryData.Stats.TotalDeaths
		currentCountryData.Stats.TotalCases = currentCountryData.Stats.ActiveCases + currentCountryData.Stats.TotalRecovered + currentCountryData.Stats.TotalDeaths
		recievedData.Stats.TotalCases += currentCountryData.Stats.TotalCases
	}
	return recievedData
}

func countryservice(country string) *data {
	recievedData := getcountrydata(country)
	recievedData.Stats.TotalCases = recievedData.Stats.ActiveCases + recievedData.Stats.TotalRecovered + recievedData.Stats.TotalDeaths
	return recievedData
}

func nameservice() *[]string {
	recievedData := getcountrynames()
	return recievedData
}
