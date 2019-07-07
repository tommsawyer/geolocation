package geolocation

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var GeoApiKey string = ""
var YandexKey string = ""

type apiResponse struct {
	City          string  `json:"city"`
	ContinentCode string  `json:"continent_code"`
	ContinentName string  `json:"continent_name"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	IP            string  `json:"ip"`
	Latitude      float64 `json:"latitude"`
	Location      struct {
		CallingCode             string `json:"calling_code"`
		Capital                 string `json:"capital"`
		CountryFlag             string `json:"country_flag"`
		CountryFlagEmoji        string `json:"country_flag_emoji"`
		CountryFlagEmojiUnicode string `json:"country_flag_emoji_unicode"`
		GeonameID               int    `json:"geoname_id"`
		IsEu                    bool   `json:"is_eu"`
		Languages               []struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Native string `json:"native"`
		} `json:"languages"`
	} `json:"location"`
	Longitude  float64 `json:"longitude"`
	RegionCode string  `json:"region_code"`
	RegionName string  `json:"region_name"`
	Type       string  `json:"type"`
	Zip        string  `json:"zip"`
}

// Get will retrieve information about ip.
// ip should be string in format 100.100.100.100
func Get(ip string) (Info, error) {
	rsp, err := http.Get(fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s", ip, GeoApiKey))
	if err != nil {
		return Info{}, err
	}

	var r apiResponse
	err = json.NewDecoder(rsp.Body).Decode(&r)
	if err != nil {
		return Info{}, err
	}

	cityName, err := translate(r.City)
	if err != nil {
		return Info{}, err
	}

	return Info{
		IP:       r.IP,
		CityName: cityName,
	}, nil
}

type Info struct {
	IP       string
	CityName string
}
