package statistics

// SeasonStatistics contains the attributes of a player specific season statistics object
type SeasonStatistics struct {
	Data struct {
		Attributes seasonStatisticAttributes `json:"attributes"`
	} `json:"data"`
}

// seasonStatisticAttributes contains the game mode statistics object for a PUBG season
type seasonStatisticAttributes struct {
	GameModeStats GameModeStatistics `json:"gameModeStats"`
}

// GameModeStatistics contains statistic objects for every kind of game mode in PUBG
type GameModeStatistics struct {
	Solo     Statistics `json:"solo"`
	SoloFPP  Statistics `json:"solo-fpp"`
	Duo      Statistics `json:"duo"`
	DuoFPP   Statistics `json:"duo-fpp"`
	Squad    Statistics `json:"squad"`
	SquadFPP Statistics `json:"squad-fpp"`
}
