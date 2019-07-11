package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

// FoursquareExplore return the 10 best venues around you
func FoursquareExplore(geoloc string, query string) (IFoursquareResponse, error) {
	var FoursquareAPIClientID string
	if len(os.Getenv("FoursquareAPIClientID")) != 0 {
		FoursquareAPIClientID = os.Getenv("FoursquareAPIClientID")
	} else {
		panic("process env FoursquareAPIClientID is Mandatory")
	}
	var FoursquareAPIClientSecret string
	if len(os.Getenv("FoursquareAPIClientSecret")) != 0 {
		FoursquareAPIClientSecret = os.Getenv("FoursquareAPIClientSecret")
	} else {
		panic("process env FoursquareAPIClientSecret is Mandatory")
	}
	request, err := http.NewRequest("GET", fmt.Sprintf("https://api.foursquare.com/v2/venues/explore?client_id=%v&client_secret=%v&v=20180323&limit=10&ll=%v&query=%v&openNow=1", FoursquareAPIClientID, FoursquareAPIClientSecret, geoloc, query), nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Fatal("Error sending request. ", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	var result IFoursquareResponse
	json.Unmarshal(body, &result)

	return result, err
}

// FoursquareRandomVenue return a random venue of a foursquare API response
func FoursquareRandomVenue(response IFoursquareResponse) IVenue {
	var venues []IVenue
	for _, group := range response.Response.Groups {
		for _, item := range group.Items {
			venues = append(venues, item.Venue)
		}
	}
	return venues[rand.Intn(len(venues))]
}
