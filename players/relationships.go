package players

import "github.com/homeblest/pubg_stat_tracker/matches"

// Relationships has references to resource objects related to a player
type Relationships struct {
	assets  interface{}
	matches matches.Data
}
