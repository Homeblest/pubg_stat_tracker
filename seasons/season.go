package seasons

// SeasonData contains a list of all the PUBG seasons
type SeasonData struct {
	Seasons []Season `json:"data"`
}

// Season describes the properties of a PUBG season
type Season struct {
	ID         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}

// Attributes describes the attributes of a PUBG season
type Attributes struct {
	IsCurrentSeason bool `json:"isCurrentSeason"`
}
