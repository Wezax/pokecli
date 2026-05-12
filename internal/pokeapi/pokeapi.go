package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/Wezax/pokecli/internal/pokecache"
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

func GetLocationArea(url string, cache *pokecache.Cache) (LocationArea, error) {
	bytes, ok := cache.Get(url);
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return LocationArea{}, err
		}
		bytes, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, err
		}
		cache.Add(url, bytes)
		defer res.Body.Close()		
	}
	var dto LocationArea
	err := json.Unmarshal(bytes, &dto)
	if err != nil {
		return LocationArea{}, err
	}
	return dto, nil
}
