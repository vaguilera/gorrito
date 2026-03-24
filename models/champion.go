package models

type Champions struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

type Mastery struct {
	ChampionLevel                int    `json:"championLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionPoints               int    `json:"championPoints"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	SummonerID                   string `json:"summonerId"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChampionID                   int    `json:"championId"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
}
