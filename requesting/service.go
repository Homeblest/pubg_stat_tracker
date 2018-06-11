package requesting

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/homeblest/pubg_stat_tracker/players"
	"github.com/slemgrim/jsonapi"
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
	parameters := url.Values{
		"filter[playerNames]": {name},
	}

	endpointURL := fmt.Sprintf("https://api.playbattlegrounds.com/shards/%s/players?%s", shard, parameters.Encode())

	reader, err := httpRequest(endpointURL, s.APIKey)

	if err != nil {
		return nil, err
	}
	result, err := jsonapi.UnmarshalManyPayload(*reader, reflect.TypeOf(new(players.Player)))

	if err != nil {
		return nil, err
	}

	thePlayers := make([]*players.Player, len(result))

	for idx, elt := range result {
		player, ok := elt.(*players.Player)
		if !ok {
			return nil, errors.New("Failed to convert players")
		}
		thePlayers[idx] = player
	}
	player := *thePlayers[0]
	fmt.Println(player.Name)

	return &player, nil
}

// Request makes a request to the PUBG API
func httpRequest(url, key string) (*io.Reader, error) {
	// Create the request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	// Configure the  request
	req.Header.Set("Authorization", key)
	req.Header.Set("Accept", "application/vnd.api+json")

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

	// Retrieve response body
	var reader io.Reader
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
	default:
		reader = response.Body
	}

	return &reader, nil
}
