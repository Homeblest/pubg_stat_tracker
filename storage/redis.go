package storage

import (
	"encoding/json"
	"errors"

	"github.com/gomodule/redigo/redis"
	"github.com/homeblest/pubg_stat_tracker/players"
)

// RedisPlayerStorage keeps player data in memory
type RedisPlayerStorage struct {
}

// Add saves the given player to the memory repository
func (m *RedisPlayerStorage) Add(player players.Player) error {
	// Get the Redis connection object
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return errors.New("Redis Dial failed")
	}

	b, err := json.Marshal(player)
	if err != nil {
		return errors.New("JSON Marshal failed")
	}

	_, err = conn.Do("SET", player.Attributes.Name, string(b))
	if err != nil {
		return errors.New("Redis SET operation failed")
	}

	return nil
}

// Get retrieves the player from memory, if it exists
func (m *RedisPlayerStorage) Get(playerName string) (*players.Player, error) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return nil, errors.New("Redis Dial failed")
	}

	objStr, err := redis.String(conn.Do("GET", playerName))
	if err != nil {
		return nil, errors.New("Failed retrieving player from redis")
	}

	b := []byte(objStr)
	player := &players.Player{}
	err = json.Unmarshal(b, player)
	if err != nil {
		return nil, errors.New("Failed to unmarshal redis string")
	}

	return player, nil
}
