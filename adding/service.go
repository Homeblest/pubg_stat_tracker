package adding

import (
	"github.com/homeblest/pubg_stat_tracker/matches"
	"github.com/homeblest/pubg_stat_tracker/players"
)

// Service provides player or match adding operations
type Service interface {
	AddMatch(match matches.Match)
	AddPlayer(player players.Player)
}

type service struct {
	matchRepo  matches.Repository
	playerRepo players.Repository
}

// New creates an adding service which handles the matches
func New(matchRepo matches.Repository, playerRepo players.Repository) Service {
	return &service{matchRepo, playerRepo}
}

// AddMatch adds the given match to the service repository
func (s *service) AddMatch(match matches.Match) {
	_ = s.matchRepo.Add(match) // TODO: Error handling
}

// AddPlayer adds the given player to the player repository
func (s *service) AddPlayer(player players.Player) {
	_ = s.playerRepo.Add(player) // TODO: Error handling
}
