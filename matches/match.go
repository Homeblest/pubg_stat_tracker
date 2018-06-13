package matches

import (
	"errors"
)

// Lite describes a list of LiteData match objects
type Lite struct {
	Data []LiteData `json:"data"`
}

// LiteData is a lite version of a match Data object, only containing a type and a match ID each
type LiteData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Match defines the properties of a PUBG match
type Match struct {
}

// ErrorMatchNotFound is used when trying to access a match that doesn't exist in the matches.Repository
var ErrorMatchNotFound = errors.New("Match not found")

// Repository provides access to the list of matches we have saved
type Repository interface {
	Get(id string) (Match, error)
	Add(Match) error
}
