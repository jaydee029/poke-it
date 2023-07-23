package pokeapi

import (
	"net/http"
	"time"
)

const base_url = "https://pokeapi.co/api/v2"

type Client struct {
	httpclient http.Client
}

func newClient() Client {
	return Client{
		httpclient: http.Client{
			Timeout: time.Minute, //kills the processs after a minute
		},
	}
}
