package players

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/homeblest/pubg_stat_tracker/links"
	"github.com/homeblest/pubg_stat_tracker/matches"
)

// Data defines the data properties of a player request from the PUBG API
type Data struct {
	Players []Player    `json:"data"`
	Links   links.Link  `json:"links"`
	Meta    interface{} `json:"meta"` // Not used by the PUBG API right now, but we must declare it to unmarshal the JSON
}

// attributes defines the attributes of a PUBG player.
type attributes struct {
	CreatedAt    time.Time   `json:"createdAt"`
	Name         string      `json:"name"`
	PatchVersion string      `json:"patchVersion"`
	ShardID      string      `json:"shardId"`
	Stats        interface{} `json:"stats"` // Not used by the PUBG API right now, stats are actually in the seasons.
	TitleID      string      `json:"titleId"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

type assets struct {
	Data []interface{} `json:"data"`
}

type relationships struct {
	Assets  assets       `json:"assets"`
	Matches matches.Lite `json:"matches"`
}

// Player defines the properties of a PUBG character
type Player struct {
	Type          string        `json:"type"`
	ID            string        `json:"id"`
	Attributes    attributes    `json:"attributes"`
	Relationships relationships `json:"relationships"`
	Links         links.Link    `json:"links"`
}

// ErrorPlayerNotFound is used when trying to access a player that doesn't exist in the players.Repository
var ErrorPlayerNotFound = errors.New("Player not found")

// Repository provides access to the list of players
type Repository interface {
	Get(id string) (*Player, error)
	Add(Player) error
}

// Used to avoid recursion in UnmarshalJSON below.
type player Player

// UnmarshalJSON takes care of casting a JSON object into a Player object
func (p *Player) UnmarshalJSON(b []byte) (err error) {
	player := player{}
	if err = json.Unmarshal(b, &player); err == nil {
		*p = Player(player)
		return
	}
	return
}
