package storage

import (
	"github.com/homeblest/pubg_stat_tracker/players"
)

// MemoryPlayerStorage keeps player data in memory
type MemoryPlayerStorage struct {
	players []players.Player
}

// Add saves the given player to the memory repository
func (m *MemoryPlayerStorage) Add(player players.Player) error {
	m.players = append(m.players, player)

	return nil
}

// Get retrieves the player from memory, if it exists
func (m *MemoryPlayerStorage) Get(name string) (*players.Player, error) {
	var emptyPlayer players.Player

	for i := range m.players {
		if m.players[i].Attributes.Name == name {
			return &m.players[i], nil
		}
	}
	return &emptyPlayer, players.ErrorPlayerNotFound
}
