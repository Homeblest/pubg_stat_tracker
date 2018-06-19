package status

import "time"

// API describes the status of the PUBG API
type API struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			ReleasedAt time.Time `json:"releasedAt"`
			Version    string    `json:"version"`
		} `json:"attributes"`
	} `json:"data"`
}
