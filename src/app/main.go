package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Name    string
	Address string
	LatLng  string
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	var LatLng string
	if len(request.URL.Query().Get("LatLng")) != 0 {
		LatLng = request.URL.Query().Get("LatLng")
	} else {
		LatLng = "45.515320,-73.563945"
	}
	var Query string
	if len(request.URL.Query().Get("Query")) != 0 {
		Query = request.URL.Query().Get("Query")
	} else {
		Query = "beer"
	}
	response, _ := FoursquareExplore(LatLng, Query)
	venue := FoursquareRandomVenue(response)

	payload, _ := json.Marshal(Result{Name: venue.Name, Address: venue.Location.Address, LatLng: fmt.Sprintf("%v, %v", venue.Location.Lat, venue.Location.Lng)})
	writer.Header().Add("content-type", "text/html; charset=utf-8")
	writer.Write(payload)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
