package pokeapi

import (
	"encoding/json"
	"net/http"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationArea(url string) (LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()
	var dto LocationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&dto)
	if err != nil {
		return LocationArea{}, err
	}
	return dto, nil
}
