package matches

import "errors"

// Match defines the properties of a PUBG match
type Match struct {
	ID string
}

// ErrorMatchNotFound is used when trying to access a match that doesn't exist in the matches.Repository
var ErrorMatchNotFound = errors.New("Match not found")

// Repository provides access to the list of matches we have saved
type Repository interface {
	Get(id string) (Match, error)
	Add(Match) error
}
