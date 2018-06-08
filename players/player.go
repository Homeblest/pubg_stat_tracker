package players

import (
	"errors"

	"github.com/homeblest/pubg_stat_tracker/matches"
)

// Player defines the properties of a PUBG character
type Player struct {
	ID      string
	Name    string
	Region  string
	Matches []matches.Match
}

// ErrorPlayerNotFound is used when trying to access a player that doesn't exist in the players.Repository
var ErrorPlayerNotFound = errors.New("Player not found")

// Repository provides access to the list of players
type Repository interface {
	Get(id string) (Player, error)
	Add(Player) error
}
