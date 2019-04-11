# Gorrito
League of Legends (LOL) API V4 wrapper for GOlang

###  Example
```golang
const token = "MysecretToken"
const region = "EUW"

gorrito := Gorrito{region: region, token: token}

var positions []LeaguePosition
status := gorrito.leaguePositions(&positions, "RANKED_SOLO_5x5", "GOLD", "I", "TOP", 0)
fmt.Println(status, positions)
```
All the endpoints are supported except the tournament ones.
These are the functions with their corresponding structures and parameters

```golang
func championRotations(champions *Champions)

func (gorrito *Gorrito) leagueChallengerByQueue(leagueItem *LeagueItem, queue string)
func (gorrito *Gorrito) leagueGrandMasterByQueue(leagueItem *LeagueItem, queue string)
func (gorrito *Gorrito) leagueMasterByQueue(leagueItem *LeagueItem, queue string)
func (gorrito *Gorrito) leagueLeagues(leagueItem *LeagueItem, leagueID string)
func (gorrito *Gorrito) leaguePositional(leagues *[]string)
func (gorrito *Gorrito) leaguePositionsBySummoner(positions *[]LeaguePosition, summonerID string)
func (gorrito *Gorrito) leaguePositions(positions *[]LeaguePosition, queue string, tier string, division string, position string, page int)

func (gorrito *Gorrito) championMasteryBySummonerID(masteries *[]Mastery, account string)
func (gorrito *Gorrito) championMasteryByBySummonerIDChampion(mastery *Mastery, account string, champion int)
func (gorrito *Gorrito) championMasteryScoresBySummonerID(score *int, account string)

func (gorrito *Gorrito) matchByID(match *Match, matchID int64
func (gorrito *Gorrito) matchesByAccountID(matchList *MatchList, account string, champion int, queue int, season int, endTime int64, beginTime int64, endIndex int, beginIndex int)
func (gorrito *Gorrito) matchTimeline(timeline *MatchTimeline, matchID int64)
func (gorrito *Gorrito) matchesByTournament(matches *[]int64, tournamentCode string)
func (gorrito *Gorrito) matchByTournament(match *Match, matchID int64, tournamentCode string)

func (gorrito *Gorrito) spectatorActiveGames(currentGameInfo *CurrentGameInfo, account string)
func (gorrito *Gorrito) spectatorFeaturedGames(featuredGames *FeaturedGames)

func (gorrito *Gorrito) lolStatus(shardStatus *ShardStatus)

func (gorrito *Gorrito) lolStatus(shardStatus *ShardStatus)

func (gorrito *Gorrito) SummonerByAccountID(summoner *Summoner, account string)
func (gorrito *Gorrito) SummonerByName(summoner *Summoner, account string)
func (gorrito *Gorrito) SummonerByPuuid(summoner *Summoner, account string)
func (gorrito *Gorrito) SummonerBySummonerID(summoner *Summoner, account string)

func (gorrito *Gorrito) thirdPartyCodeBySummonerID(code *string, account string)
```