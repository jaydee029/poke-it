package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Pokemonreq(arg string) (Pokemon, error) {
	end_url := "/pokemon/"
	full_url := base_url + end_url + arg

	//checking the cache memeory for the req endpoint
	cacheres, ok := c.cache.Get(full_url)
	if ok {
		fmt.Println("cache hit!")
		Pokeinfo := Pokemon{}
		err := json.Unmarshal(cacheres, &Pokeinfo)
		if err != nil {
			return Pokemon{}, err
		}

		return Pokeinfo, nil

	}

	req, err := http.NewRequest("GET", full_url, nil)

	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpclient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code : %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	pokeinfo := Pokemon{}
	err = json.Unmarshal(data, &pokeinfo)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(full_url, []byte(data))
	return pokeinfo, nil

}
