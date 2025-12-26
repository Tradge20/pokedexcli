package pokeapi

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocationAreas(pageURL *string ) (LocationAreaResp, error ) {
	fullURL := baseURL + "/location-area/?offset=0&limit=20/"
	if pageURL != nil {
		fullURL = *pageURL
	}


	dat, ok := c.cache.Get(fullURL)
	if ok {
		locationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(dat, &locationAreaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullURL, dat)
	return locationAreaResp, nil
}