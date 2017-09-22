package latlong

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"strings"
	"time"
	"errors"
)

var urls = []string{"https://ipinfo.io/%s/json", "https://ipapi.co/%s/json/"}

// Geodata. Note that the JSON from the api endpoints should either give both latitude and longitude or
// give a comma seperated lat,lng Loc. The object contains a country which can be retrieved by the GetCountry
// method. Optionally, a city var is available.
type Geodata struct {
	Ip string
	Latitude float64
	Longitude float64
	Loc string
	City string
	Country string
	Country_name string
}

// GeoConverter interface for the Geodata struct
type GeoConverter interface {
	Latstring() string
	Lngstring() string
	GetCountry() string
}

// Latstring returns a string representation of the (float64) Latitude
func (g Geodata) Latstring() string {
	return strconv.FormatFloat(g.Latitude, 'f', 4, 64)
}

// Lngstring returns a string representation of the (float64) Longitude
func (g Geodata) Lngstring() string {
	return strconv.FormatFloat(g.Longitude, 'f', 4, 64)
}

// GetCountry is an adaptor as the country is named differently in different API endpoints.
func (g Geodata) GetCountry() string {
	if g.Country != "" {
		return g.Country
	} else if g.Country_name != "" {
		return g.Country_name
	}
	return ""
}

var myClient = &http.Client{Timeout: 3 * time.Second}

// Latlong string ip, address of the ip on which a geolocation should be found.
// This is the main function.
// Returns Geodata object and nil or nil and error in case the geolocation cannot be retreived.
func Latlong(ip string) (*Geodata, error) {

	geo := new(Geodata)

	for _, url := range urls {
		formattedUrl := fmt.Sprintf(url, ip)
		err := getJson(formattedUrl, geo)
		// if there is no error, we can savely break the loop here
		if err == nil {
			break
		}
	}

	if geo.Longitude == 0 {
		geodata := strings.Split(geo.Loc, ",")

		lat, err := strconv.ParseFloat(geodata[0], 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Unable to parse latitude from geodata [%s]", geodata))
		}
		geo.Latitude = lat

		lng, err := strconv.ParseFloat(geodata[1], 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Unable to parse latitude from geodata [%s]", geodata))
		}
		geo.Longitude = lng
	}
	return geo, nil
}

// SetUrls can be used to override the preset urls from which the geolocation data is retrieved.
// The response of this API endpoint should be json and be compatible with the Geodata struct
func SetUrls (set []string){
	urls = set
}

// getJson sets the retrieved JSON values on the target.
// Returns nil or an Error when unable to fetch the json from the url
func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(target)
	return nil
}
