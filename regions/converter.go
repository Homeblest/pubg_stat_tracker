package regions

// regionByShards is a mapping between PUBG server shards and Region identifiers
var regionByShards = map[string]string{
	"xbox-as":  "XAS",
	"xbox-eu":  "XEU",
	"xbox-na":  "XBNA",
	"xbox-oc":  "XBOC",
	"pc-krjp":  "KR",
	"pc-na":    "NA",
	"pc-eu":    "EU",
	"pc-oc":    "OC",
	"pc-kakao": "KA",
	"pc-sea":   "SE",
	"pc-sa":    "SA",
	"pc-as":    "AS",
}

// GetShardIDFromRegion returns the ID of the PUBG server shard
func GetShardIDFromRegion(region string) string {
	return regionByShards[region]
}
