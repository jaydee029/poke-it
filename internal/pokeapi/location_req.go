package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationArearesponse() (LocationAreas, error) {
	end_url := "/location-area"
	full_url := base_url + end_url

	req, err := http.NewRequest("GET", full_url, nil)

	if err != nil {
		return LocationAreas{}, err
	}

	res, err := c.httpclient.Do(req)

	if err != nil {
		return LocationAreas{}, err
	}

	if res.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("bad status code : %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	defer res.Body.Close()

	LocationAreaValues := LocationAreas{}
	err = json.Unmarshal(data, &LocationAreaValues)
	if err != nil {
		return LocationAreas{}, err
	}

	return LocationAreaValues, nil

}
