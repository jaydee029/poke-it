package pokeapi

type LocationAreas struct { // in go json data is carried in bytes in a struct and
	Count    int     `json:"count"` // then unmarshaled , json to go tool used to get the shape of response body
	Next     *string `json:"next"`  //string pointger used in case of nextg and previous since it may be null
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
