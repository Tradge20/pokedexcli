package pokeapi

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
)


func (c *Client) GetPokemon(pokemonName string ) (Pokemon, error ) {
	fullURL := baseURL + "/pokemon/" + pokemonName
	


	dat, ok := c.cache.Get(fullURL)
	if ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	locationArea := Pokemon{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, dat)
	return locationArea, nil
}