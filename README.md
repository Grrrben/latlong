# latlong
Go microservice to obtain the latitude and longitude of an IP address

## usage
```
import (
  "github.com/grrrben/latlong"
)
// Latlong string ip, address of the ip on which a geolocation should be found.
// Returns Geodata object and nil or nil and error in case the geolocation cannot be retreived.
geo, err := latlong.Latlong(ipaddress)
```

## Geodata object

The Geodata object contains both Latitude and Longitude of the location that is related to the given IP address. 

```
// Geodata. Note that the JSON from the api endpoints should either give both latitude and longitude or
// give a comma seperated lat,lng Loc. The object contains a country which can be retrieved by the 
// GetCountry method. Optionally, a city var is available.
type Geodata struct {
	Ip string
	Latitude float64
	Longitude float64
	Loc string
	City string
	Country string
	Country_name string
}
```

There are helper functions that return the Latitude and Longitude as a string. 

```
// Latstring returns a string representation of the (float64) Latitude
func (g Geodata) Latstring() string {
	return strconv.FormatFloat(g.Latitude, 'f', 4, 64)
}

// Lngstring returns a string representation of the (float64) Longitude
func (g Geodata) Lngstring() string {
	return strconv.FormatFloat(g.Longitude, 'f', 4, 64)
}
```
