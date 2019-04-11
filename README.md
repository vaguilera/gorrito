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
championRotations(champions *Champions)

leagueChallengerByQueue(leagueItem *LeagueItem, queue string)
leagueGrandMasterByQueue(leagueItem *LeagueItem, queue string)
leagueMasterByQueue(leagueItem *LeagueItem, queue string)
leagueLeagues(leagueItem *LeagueItem, leagueID string)
leaguePositional(leagues *[]string)
leaguePositionsBySummoner(positions *[]LeaguePosition, summonerID string)
leaguePositions(positions *[]LeaguePosition, queue string, tier string, division string, position string, page int)

championMasteryBySummonerID(masteries *[]Mastery, account string)
championMasteryByBySummonerIDChampion(mastery *Mastery, account string, champion int)
championMasteryScoresBySummonerID(score *int, account string)

matchByID(match *Match, matchID int64
matchesByAccountID(matchList *MatchList, account string, champion int, queue int, season int, endTime int64, beginTime int64, endIndex int, beginIndex int)
matchTimeline(timeline *MatchTimeline, matchID int64)
matchesByTournament(matches *[]int64, tournamentCode string)
matchByTournament(match *Match, matchID int64, tournamentCode string)

spectatorActiveGames(currentGameInfo *CurrentGameInfo, account string)
spectatorFeaturedGames(featuredGames *FeaturedGames)

lolStatus(shardStatus *ShardStatus)

lolStatus(shardStatus *ShardStatus)

SummonerByAccountID(summoner *Summoner, account string)
SummonerByName(summoner *Summoner, account string)
SummonerByPuuid(summoner *Summoner, account string)
SummonerBySummonerID(summoner *Summoner, account string)

thirdPartyCodeBySummonerID(code *string, account string)
```
