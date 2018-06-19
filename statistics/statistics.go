package statistics

// Statistics contains a player's aggregated stats for a game mode in the context of a season
type Statistics struct {
	KillAssists         int     `json:"assists"`
	BoostsUsed          int     `json:"boosts"`
	DBNOs               int     `json:"dBNOs"`
	DailyKills          int     `json:"dailyKills"`
	DamageDealt         float32 `json:"damageDealt"`
	HeadshotKills       int     `json:"headshotKills"`
	Heals               int     `json:"heals"`
	KillPoints          float32 `json:"killPoints"`
	Kills               int     `json:"kills"`
	LongestKill         float32 `json:"longesKill"`
	LongestTimeSurvived float32 `json:"longestTimeSurvived"`
	Losses              int     `json:"losses"`
	LongestKillStreak   int     `json:"maxKillStreaks"`
	Revives             int     `json:"revives"`
	VehicleDistance     float32 `json:"rideDistance"`
	RoadKills           int     `json:"roadKills"`
	RoundMostKills      int     `json:"roundMostKills"`
	RoundsPlayed        int     `json:"roundsPlayed"`
	Suicides            int     `json:"suicides"`
	TeamKills           int     `json:"teamKills"`
	TotalTimeSurvived   float32 `json:"timeSurvived"`
	Top10s              int     `json:"top10s"`
	VehiclesDestroyed   int     `json:"vehicleDestroys"`
	DistanceWalked      float32 `json:"walkDistance"`
	WeaponsAcquired     int     `json:"weaponsAcquired"`
	WeeklyKills         int     `json:"weeklyKills"`
	WinPoints           float32 `json:"winPoints"`
	Wins                int     `json:"wins"`
}
