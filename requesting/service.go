package requesting

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/homeblest/pubg_stat_tracker/players"
)

// Service takes care of requesting data from the PUBG API
type Service interface {
	RequestPlayer(string, string) (*players.Player, error)
}

type service struct {
	APIKey string
}

// NewService creates a request service
func NewService(APIKey string) Service {
	return &service{
		APIKey: APIKey,
	}
}

func (s *service) RequestPlayer(shard, name string) (*players.Player, error) {
	var player players.Player

	parameters := url.Values{
		"filter[playerNames]": {name},
	}

	endpointURL := fmt.Sprintf("https://api.playbattlegrounds.com/shards/%s/players?%s", shard, parameters.Encode())

	fmt.Println(endpointURL)

	buffer, err := httpRequest(endpointURL, s.APIKey)

	if err != nil {
		return nil, err
	}

	fmt.Printf("data:\n%s\n", buffer)

	return &player, nil
}

// Request makes a request to the PUBG API
func httpRequest(url, key string) (*bytes.Buffer, error) {
	// Create the request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	// Configure the  request
	req.Header.Set("Authorization", key)
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("Accept-Encoding", "gzip")

	// Send the request
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		response.Body.Close()
		return nil, fmt.Errorf("HTTP request failed: %s", response.Status)
	}

	var reader io.ReadCloser

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
	default:
		reader = response.Body
	}

	var buffer bytes.Buffer
	buffer.ReadFrom(reader)

	return &buffer, nil
}
