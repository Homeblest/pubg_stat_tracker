package players

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/homeblest/pubg_stat_tracker/matches"
)

// Player defines the properties of a PUBG character
type Player struct {
	ID           string           `jsonapi:"primary,player"`
	Name         string           `jsonapi:"attr,name"`
	ShardID      string           `jsonapi:"attr,shardId"`
	CreatedAt    time.Time        `jsonapi:"attr,createdAt,iso8601"`
	UpdatedAt    time.Time        `jsonapi:"attr,updatedAt,iso8601"`
	PatchVersion string           `jsonapi:"attr,patchVersion"`
	TitleID      string           `jsonapi:"attr,titleId"`
	Matches      []*matches.Match `jsonapi:"relation,matches"`
}

// ErrorPlayerNotFound is used when trying to access a player that doesn't exist in the players.Repository
var ErrorPlayerNotFound = errors.New("Player not found")

// Repository provides access to the list of players
type Repository interface {
	Get(id string) (Player, error)
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
