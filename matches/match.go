package matches

import (
	"errors"
	"time"
)

// Match defines the properties of a PUBG match
type Match struct {
	ID           string    `jsonapi:"primary,match"`
	CreatedAt    time.Time `jsonapi:"attr,createdAt,iso8601"`
	Duration     int       `jsonapi:"attr,duration"`
	GameMode     string    `jsonapi:"attr,gameMode"`
	PatchVersion string    `jsonapi:"attr,patchVersion"`
	ShardID      string    `jsonapi:"attr,shardId"`
	TitleID      string    `jsonapi:"attr,titleId"`
}

// ErrorMatchNotFound is used when trying to access a match that doesn't exist in the matches.Repository
var ErrorMatchNotFound = errors.New("Match not found")

// Repository provides access to the list of matches we have saved
type Repository interface {
	Get(id string) (Match, error)
	Add(Match) error
}
