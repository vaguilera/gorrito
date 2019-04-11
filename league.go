package gorrito

import "strconv"

type LeagueItem struct {
	LeagueID string       `json:"leagueId"`
	Tier     string       `json:"tier"`
	Entries  []LeagueItem `json:"entries"`
	Queue    string       `json:"queue"`
	Name     string       `json:"name"`
}

type LeagueItemn struct {
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

type LeaguePosition struct {
	Tier         string     `json:"tier"`
	SummonerName string     `json:"summonerName"`
	HotStreak    bool       `json:"hotStreak"`
	MiniSeries   MiniSeries `json:"miniSeries"`
	Wins         int        `json:"wins"`
	Veteran      bool       `json:"veteran"`
	Losses       int        `json:"losses"`
	Rank         string     `json:"rank"`
	LeagueName   string     `json:"leagueName"`
	Inactive     bool       `json:"inactive"`
	FreshBlood   bool       `json:"freshBlood"`
	Position     string     `json:"position"`
	LeagueID     string     `json:"leagueId"`
	QueueType    string     `json:"queueType"`
	SummonerID   string     `json:"summonerId"`
	LeaguePoints int        `json:"leaguePoints"`
}

func (gorrito *Gorrito) leagueChallengerByQueue(leagueItem *LeagueItem, queue string) int {

	return gorrito.requestAPI(leagueChallengerByQueue, leagueItem, map[string]string{"queue": queue})
}

func (gorrito *Gorrito) leagueGrandMasterByQueue(leagueItem *LeagueItem, queue string) int {

	return gorrito.requestAPI(leagueGrandMasterByQueue, leagueItem, map[string]string{"queue": queue})
}

func (gorrito *Gorrito) leagueMasterByQueue(leagueItem *LeagueItem, queue string) int {

	return gorrito.requestAPI(leagueMasterByQueue, leagueItem, map[string]string{"queue": queue})
}

func (gorrito *Gorrito) leagueLeagues(leagueItem *LeagueItem, leagueID string) int {

	return gorrito.requestAPI(leagueLeagues, leagueItem, map[string]string{"leagueId": leagueID})
}

func (gorrito *Gorrito) leaguePositional(leagues *[]string) int {

	return gorrito.requestAPI(leaguePositional, leagues, nil)
}

func (gorrito *Gorrito) leaguePositionsBySummoner(positions *[]LeaguePosition, summonerID string) int {

	return gorrito.requestAPI(leaguePositionsBySummoner, positions, map[string]string{"encryptedSummonerId": summonerID})
}

func (gorrito *Gorrito) leaguePositions(positions *[]LeaguePosition, queue string, tier string, division string, position string, page int) int {

	return gorrito.requestAPI(leaguePositions, positions, map[string]string{
		"positionalQueue": queue,
		"tier":            tier,
		"division":        division,
		"position":        position,
		"page":            strconv.Itoa(page),
	})
}
