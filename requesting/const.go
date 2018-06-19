package requesting

const (
	pubgAPIBaseURL      string = "https://api.playbattlegrounds.com"
	playersEndpoint     string = "/players"
	pubgAPIBaseShardURL string = pubgAPIBaseURL + "/shards/%s"
	seasonsEndpoint     string = "/seasons"
	statisticsEndpoint  string = pubgAPIBaseShardURL + playersEndpoint + "/%s" + seasonsEndpoint + "/%s"
	statusEndpoint      string = "/status"
)
