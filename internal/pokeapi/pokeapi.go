package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

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

func (c *Client) GetPokemonByName(linkString string, name *string) (GetPokemon, error) {
	url, err := url.Parse(linkString)
	if err != nil {
		return GetPokemon{}, err
	}
	if name != nil {
		url = url.JoinPath(*name + "/")
	}
	bytes, ok := c.cache.Get(url.String())
	if !ok {
		res, err := http.Get(url.String())
		if err != nil {
			return GetPokemon{}, err
		}
		if res.StatusCode > 200 {
			return GetPokemon{}, fmt.Errorf("Status code != 200. Got: %d\n", res.StatusCode)
		}
		bytes, err = io.ReadAll(res.Body)
		if err != nil {
			return GetPokemon{}, err
		}
		c.cache.Add(url.String(), bytes)
		defer res.Body.Close()
	}
	var dto GetPokemon
	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		return GetPokemon{}, err
	}
	return dto, nil
}