package models

type Account struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type ActiveShard struct {
	Puuid string `json:"puuid"`
	Game  string `json:"game"`
	Shard string `json:"activeShard"`
}
