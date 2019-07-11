package main

type IVenue struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location struct {
		Address        string  `json:"address"`
		CrossStreet    string  `json:"crossStreet"`
		Lat            float64 `json:"lat"`
		Lng            float64 `json:"lng"`
		LabeledLatLngs []struct {
			Label string  `json:"label"`
			Lat   float64 `json:"lat"`
			Lng   float64 `json:"lng"`
		} `json:"labeledLatLngs"`
		Distance         int      `json:"distance"`
		PostalCode       string   `json:"postalCode"`
		Cc               string   `json:"cc"`
		City             string   `json:"city"`
		State            string   `json:"state"`
		Country          string   `json:"country"`
		FormattedAddress []string `json:"formattedAddress"`
	} `json:"location"`
	Categories []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		PluralName string `json:"pluralName"`
		ShortName  string `json:"shortName"`
		Icon       struct {
			Prefix string `json:"prefix"`
			Suffix string `json:"suffix"`
		} `json:"icon"`
		Primary bool `json:"primary"`
	} `json:"categories"`
	VenuePage struct {
		ID string `json:"id"`
	} `json:"venuePage"`
}

type IFoursquareResponse struct {
	Meta struct {
		Code      int    `json:"code"`
		RequestID string `json:"requestId"`
	} `json:"meta"`
	Response struct {
		Warning struct {
			Text string `json:"text"`
		} `json:"warning"`
		SuggestedRadius           int    `json:"suggestedRadius"`
		HeaderLocation            string `json:"headerLocation"`
		HeaderFullLocation        string `json:"headerFullLocation"`
		HeaderLocationGranularity string `json:"headerLocationGranularity"`
		TotalResults              int    `json:"totalResults"`
		SuggestedBounds           struct {
			Ne struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"ne"`
			Sw struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"sw"`
		} `json:"suggestedBounds"`
		Groups []struct {
			Type  string `json:"type"`
			Name  string `json:"name"`
			Items []struct {
				Reasons struct {
					Count int `json:"count"`
					Items []struct {
						Summary    string `json:"summary"`
						Type       string `json:"type"`
						ReasonName string `json:"reasonName"`
					} `json:"items"`
				} `json:"reasons"`
				Venue struct {
					ID       string `json:"id"`
					Name     string `json:"name"`
					Location struct {
						Address        string  `json:"address"`
						CrossStreet    string  `json:"crossStreet"`
						Lat            float64 `json:"lat"`
						Lng            float64 `json:"lng"`
						LabeledLatLngs []struct {
							Label string  `json:"label"`
							Lat   float64 `json:"lat"`
							Lng   float64 `json:"lng"`
						} `json:"labeledLatLngs"`
						Distance         int      `json:"distance"`
						PostalCode       string   `json:"postalCode"`
						Cc               string   `json:"cc"`
						City             string   `json:"city"`
						State            string   `json:"state"`
						Country          string   `json:"country"`
						FormattedAddress []string `json:"formattedAddress"`
					} `json:"location"`
					Categories []struct {
						ID         string `json:"id"`
						Name       string `json:"name"`
						PluralName string `json:"pluralName"`
						ShortName  string `json:"shortName"`
						Icon       struct {
							Prefix string `json:"prefix"`
							Suffix string `json:"suffix"`
						} `json:"icon"`
						Primary bool `json:"primary"`
					} `json:"categories"`
					VenuePage struct {
						ID string `json:"id"`
					} `json:"venuePage"`
				} `json:"venue"`
			} `json:"items"`
		} `json:"groups"`
	} `json:"response"`
}
