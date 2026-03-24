package models

type LeagueList struct {
	LeagueID string       `json:"leagueId"`
	Tier     string       `json:"tier"`
	Entries  []LeagueItem `json:"entries"`
	Queue    string       `json:"queue"`
	Name     string       `json:"name"`
}

type LeagueItem struct {
	SummonerName string     `json:"summonerName"`
	HotStreak    bool       `json:"hotStreak"`
	MiniSeries   MiniSeries `json:"miniSeries"`
	Wins         int        `json:"wins"`
	Veteran      bool       `json:"veteran"`
	Losses       int        `json:"losses"`
	Rank         string     `json:"rank"`
	Inactive     bool       `json:"inactive"`
	FreshBlood   bool       `json:"freshBlood"`
	SummonerID   string     `json:"summonerId"`
	LeaguePoints int        `json:"leaguePoints"`
}

type MiniSeries struct {
	Progress string `json:"progress"`
	Losses   string `json:"losses"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

type LeagueEntry struct {
	LeagueID  string `json:"leagueId"`
	QueueType string `json:"queueType"`
	Tier      string `json:"tier"`
	Rank      string `json:"rank"`
	LeagueItem
}
