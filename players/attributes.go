package players

// Attributes describes the attributes of a player object.
type Attributes struct {
	Name         string
	shardID      string
	stats        interface{}
	createdAt    string
	patchVersion string
	titleID      string
}
