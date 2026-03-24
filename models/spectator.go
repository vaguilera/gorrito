package models

type CurrentGameInfo struct {
	GameID            int64  `json:"gameId"`
	GameStartTime     int64  `json:"gameStartTime"`
	PlatformID        string `json:"platformId"`
	GameMode          string `json:"gameMode"`
	MapID             int    `json:"mapId"`
	GameType          string `json:"gameType"`
	GameQueueConfigID int    `json:"gameQueueConfigId"`
	Observers         struct {
		EncryptionKey string `json:"encryptionKey"`
	} `json:"observers"`
	Participants    []ParticipantGameInfo `json:"participants"`
	GameLength      int                   `json:"gameLength"`
	BannedChampions []struct {
		TeamID     int   `json:"teamId"`
		ChampionID int64 `json:"championId"`
		PickTurn   int64 `json:"pickTurn"`
	} `json:"bannedChampions"`
}

type ParticipantGameInfo struct {
	ProfileIconID            int    `json:"profileIconId"`
	ChampionID               int    `json:"championId"`
	SummonerName             string `json:"summonerName"`
	GameCustomizationObjects []struct {
		Category string `json:"category"`
		Content  string `json:"content"`
	} `json:"gameCustomizationObjects"`
	Bot   bool `json:"bot"`
	Perks struct {
		PerkStyle    int64   `json:"perkStyle"`
		PerkIds      []int64 `json:"perkIds"`
		PerkSubStyle int64   `json:"perkSubStyle"`
	} `json:"perks"`
	Spell2ID   int64  `json:"spell2Id"`
	TeamID     int64  `json:"teamId"`
	Spell1ID   int64  `json:"spell1Id"`
	SummonerID string `json:"summonerId"`
}

type FeaturedGames struct {
	ClientRefreshInterval int64 `json:"clientRefreshInterval"`
	GameList              []struct {
		GameID            int64  `json:"gameId"`
		GameStartTime     int64  `json:"gameStartTime"`
		PlatformID        string `json:"platformId"`
		GameMode          string `json:"gameMode"`
		MapID             int64  `json:"mapId"`
		GameType          string `json:"gameType"`
		GameQueueConfigID int64  `json:"gameQueueConfigId"`
		Observers         struct {
			EncryptionKey string `json:"encryptionKey"`
		} `json:"observers"`
		Participants    []ParticipantFeatured `json:"participants"`
		GameLength      int64                 `json:"gameLength"`
		BannedChampions []struct {
			TeamID     int   `json:"teamId"`
			ChampionID int64 `json:"championId"`
			PickTurn   int64 `json:"pickTurn"`
		} `json:"bannedChampions"`
	} `json:"gameList"`
}

type ParticipantFeatured struct {
	ProfileIconID int    `json:"profileIconId"`
	ChampionID    int    `json:"championId"`
	SummonerName  string `json:"summonerName"`
	Bot           bool   `json:"bot"`
	Spell2ID      int64  `json:"spell2Id"`
	TeamID        int64  `json:"teamId"`
	Spell1ID      int64  `json:"spell1Id"`
}
