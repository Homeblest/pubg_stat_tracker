package regions

// regionByShards is a mapping between PUBG server shards and Region identifiers
var regionByShards = map[string]string{
	"XAS":  "xbox-as",
	"XEU":  "xbox-eu",
	"XBNA": "xbox-na",
	"XBOC": "xbox-oc",
	"KR":   "pc-krjp",
	"NA":   "pc-na",
	"EU":   "pc-eu",
	"OC":   "pc-oc",
	"KA":   "pc-kakao",
	"SE":   "pc-sea",
	"SA":   "pc-sa",
	"AS":   "pc-as",
}

// GetShardIDFromRegion returns the ID of the PUBG server shard
func GetShardIDFromRegion(region string) string {
	return regionByShards[region]
}
