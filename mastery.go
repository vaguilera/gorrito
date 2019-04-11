package gorrito

import "strconv"

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

func (gorrito *Gorrito) championMasteryBySummonerID(masteries *[]Mastery, account string) int {

	return gorrito.requestAPI(championMasteryBySummonerID, masteries, map[string]string{"encryptedSummonerId": account})
}

func (gorrito *Gorrito) championMasteryByBySummonerIDChampion(mastery *Mastery, account string, champion int) int {

	return gorrito.requestAPI(championMasteryByBySummonerIDChampion, mastery, map[string]string{"encryptedSummonerId": account, "championId": strconv.Itoa(champion)})
}

func (gorrito *Gorrito) championMasteryScoresBySummonerID(score *int, account string) int {
	return gorrito.requestAPI(championMasteryScoresBySummonerID, score, map[string]string{"encryptedSummonerId": account})
}
