package requesting

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/homeblest/pubg_stat_tracker/players"
	"github.com/homeblest/pubg_stat_tracker/seasons"
	"github.com/homeblest/pubg_stat_tracker/statistics"
	"github.com/homeblest/pubg_stat_tracker/status"
)

// Service takes care of requesting data from the PUBG API
type Service interface {
	RequestPlayer(string, string) (*players.Player, error)
	RequestSeasons(string) (*seasons.SeasonData, error)
	RequestSeasonStatistics(string, string, string) (*statistics.SeasonStatistics, error)
	RequestStatus() (*status.API, error)
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

func (s *service) RequestStatus() (*status.API, error) {
	apiURL := pubgAPIBaseURL + statusEndpoint

	body, err := createRequest(apiURL, s.APIKey, url.Values{})
	if err != nil {
		return nil, err
	}

	status := &status.API{}

	err = json.NewDecoder(body).Decode(&status)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (s *service) RequestPlayer(name, shard string) (*players.Player, error) {
	players, err := s.RequestPlayers(name, shard)
	if err != nil {
		return nil, err
	}
	player := players[0]

	return &player, nil
}

func (s *service) RequestPlayers(name, shard string) ([]players.Player, error) {
	shardURL := fmt.Sprintf(pubgAPIBaseShardURL, string(shard))
	playersURL := strings.Join([]string{shardURL, playersEndpoint}, "")

	query := url.Values{"filter[playerNames]": {name}}

	body, err := createRequest(playersURL, s.APIKey, query)
	if err != nil {
		return nil, err
	}

	playersData := &players.Data{}

	err = json.NewDecoder(body).Decode(playersData)
	if err != nil {
		return nil, err
	}

	players := *playersData

	return players.Players, nil
}

func (s *service) RequestSeasons(shard string) (*seasons.SeasonData, error) {
	shardURL := fmt.Sprintf(pubgAPIBaseShardURL, string(shard))
	apiURL := strings.Join([]string{shardURL, seasonsEndpoint}, "")

	body, err := createRequest(apiURL, s.APIKey, url.Values{})
	if err != nil {
		return nil, err
	}

	seasonData := &seasons.SeasonData{}

	err = json.NewDecoder(body).Decode(seasonData)
	if err != nil {
		return nil, err
	}

	return seasonData, nil
}

func (s *service) RequestSeasonStatistics(playerID, seasonID, shard string) (*statistics.SeasonStatistics, error) {
	apiURL := fmt.Sprintf(statisticsEndpoint, string(shard), playerID, seasonID)

	body, err := createRequest(apiURL, s.APIKey, url.Values{})
	if err != nil {
		return nil, err
	}

	seasonStatistics := &statistics.SeasonStatistics{}

	err = json.NewDecoder(body).Decode(seasonStatistics)
	if err != nil {
		return nil, err
	}

	return seasonStatistics, nil
}

// createRequest makes a http GET request to the PUBG API
func createRequest(url, key string, query url.Values) (*bytes.Buffer, error) {
	// Create the request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	// Configure the  request
	req.Header.Set("Authorization", key)
	req.Header.Set("Accept", "application/vnd.api+json")

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	// Send the request
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusUnauthorized:
			return nil, errors.New("API key invalid or missing")
		case http.StatusNotFound:
			return nil, errors.New("The specified resource was not found")
		case http.StatusUnsupportedMediaType:
			return nil, errors.New("Content type incorrect or not specified")
		case http.StatusTooManyRequests:
			return nil, errors.New("Too many requests")
		default:
			return nil, errors.New(res.Status)
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(body), nil
}
