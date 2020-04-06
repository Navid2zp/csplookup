package csplookup

type IPLookup struct {
	Status    string `json:"Status"`
	ErrorCode string `json:"ErrorCode"`
	Message   string `json:"Message"`
	Result    struct {
		City struct {
			GeoNameID uint              `json:"GeoNameID"`
			Names     map[string]string `json:"Names"`
		} `json:"City"`
		Continent struct {
			Code      string            `json:"Code"`
			GeoNameID uint              `json:"GeoNameID"`
			Names     map[string]string `json:"Names"`
		} `json:"Continent"`
		Country struct {
			GeoNameID         uint              `json:"GeoNameID"`
			IsInEuropeanUnion bool              `json:"IsInEuropeanUnion"`
			IsoCode           string            `json:"IsoCode"`
			Names             map[string]string `json:"Names"`
		} `json:"Country"`
		Location struct {
			AccuracyRadius uint16  `json:"AccuracyRadius"`
			Latitude       float64 `json:"Latitude"`
			Longitude      float64 `json:"Longitude"`
			MetroCode      uint    `json:"MetroCode"`
			TimeZone       string  `json:"TimeZone"`
		} `json:"Location"`
		Postal struct {
			Code string `json:"Code"`
		} `json:"Postal"`
		RegisteredCountry struct {
			GeoNameID         uint              `json:"GeoNameID"`
			IsInEuropeanUnion bool              `json:"IsInEuropeanUnion"`
			IsoCode           string            `json:"IsoCode"`
			Names             map[string]string `json:"Names"`
		} `json:"RegisteredCountry"`
		RepresentedCountry struct {
			GeoNameID         uint              `json:"GeoNameID"`
			IsInEuropeanUnion bool              `json:"IsInEuropeanUnion"`
			IsoCode           string            `json:"IsoCode"`
			Names             map[string]string `json:"Names"`
			Type              string            `json:"Type"`
		} `json:"RepresentedCountry"`
		Subdivisions []struct {
			GeoNameID uint              `json:"GeoNameID"`
			IsoCode   string            `json:"IsoCode"`
			Names     map[string]string `json:"Names"`
		} `json:"Subdivisions"`
		Traits struct {
			IsAnonymousProxy    bool `json:"IsAnonymousProxy"`
			IsSatelliteProvider bool `json:"IsSatelliteProvider"`
		} `json:"Traits"`
	} `json:"Result"`
}

type Client struct {
	APIKey string
}
