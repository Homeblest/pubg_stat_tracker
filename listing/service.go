package listing

import "github.com/homeblest/pubg_stat_tracker/players"

// Service provides player listing operations
type Service interface {
	GetPlayer(string) (players.Player, error)
}

type service struct {
	playerRepo players.Repository
}

// New creates a player listing service
func New(playerRepo players.Repository) Service {
	return &service{playerRepo}
}

// GetPlayer returns a player in the playerRepository
func (s *service) GetPlayer(name string) (players.Player, error) {
	return s.playerRepo.Get(name)
}
