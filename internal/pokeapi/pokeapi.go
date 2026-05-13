package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name       string `json:"name"`
		linkString string `json:"linkString"`
	} `json:"results"`
}

type LocationAreaByName struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationArea(linkString string) (LocationArea, error) {
	bytes, ok := c.cache.Get(linkString)
	if !ok {
		res, err := http.Get(linkString)
		if err != nil {
			return LocationArea{}, err
		}
		bytes, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, err
		}
		c.cache.Add(linkString, bytes)
		defer res.Body.Close()
	}
	var dto LocationArea
	err := json.Unmarshal(bytes, &dto)
	if err != nil {
		return LocationArea{}, err
	}
	return dto, nil
}

func (c *Client) GetLocationAreaByName(linkString string, name *string) (LocationAreaByName, error) {
	url, err := url.Parse(linkString)
	if err != nil {
		return LocationAreaByName{}, err
	}
	if name != nil {
		url = url.JoinPath(*name + "/")
	}
	bytes, ok := c.cache.Get(url.String())
	if !ok {
		res, err := http.Get(url.String())
		if err != nil {
			return LocationAreaByName{}, err
		}
		if res.StatusCode > 200 {
			return LocationAreaByName{}, fmt.Errorf("Status code != 200. Got: %d\n", res.StatusCode)
		}
		bytes, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaByName{}, err
		}
		c.cache.Add(url.String(), bytes)
		defer res.Body.Close()
	}
	var dto LocationAreaByName
	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		return LocationAreaByName{}, err
	}
	return dto, nil
}
