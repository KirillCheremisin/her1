package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"workshop/internal/api"
	"workshop/internal/config"
)

const getJokePath = "/api?format=json"

type JokeClient struct {
	config *config.Config
}

func NewJokeClient(cfg *config.Config) *JokeClient {
	return &JokeClient{
		config: cfg,
	}
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.config.BaseURL + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %s", http.StatusText(resp.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
