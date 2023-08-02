package pokeapi

import (
	"net/http"
	"time"

	"github.com/jaydee029/poke-it/internal/pokecache"
)

const base_url = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpclient http.Client
}

func NewClient(cacheDuration time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheDuration),
		httpclient: http.Client{
			Timeout: time.Minute, //kills the processs after a minute
		},
	}
}
