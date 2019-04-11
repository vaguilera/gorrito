package gorrito

import (
	"fmt"
	"testing"
)

/*
	This file isn't a formal unit test file
	I used these tests only to develop the library
	To properly do unit testing of the library a mock server will be needed
*/

const token = ""
const region = "EUW"

func Test_summonersByName(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var summoner Summoner
	err := gorrito.SummonerByName(&summoner, "shötgan")
	fmt.Println(err, summoner)

	/*    if total != 10 {
	      t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	   } */
}

func Test_championMasteryBySummonerID(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var masteries []Mastery
	err := gorrito.championMasteryBySummonerID(&masteries, "id5l4_4JjOEOn3VpCWdj5Cb4s5lj2nmP5BTr-0AempQccA0")
	fmt.Println(err, masteries)
}

func Test_championMasteryByBySummonerIDChampion(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var mastery Mastery
	err := gorrito.championMasteryByBySummonerIDChampion(&mastery, "id5l4_4JjOEOn3VpCWdj5Cb4s5lj2nmP5BTr-0AempQccA0", 110)
	fmt.Println(err, mastery)
}

func Test_championMasteryScoresBySummonerID(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var score int
	err := gorrito.championMasteryScoresBySummonerID(&score, "id5l4_4JjOEOn3VpCWdj5Cb4s5lj2nmP5BTr-0AempQccA0")
	fmt.Println(err, score)
}

func Test_championRotations(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var champions Champions
	err := gorrito.championRotations(&champions)
	fmt.Println(err, champions)
}

func Test_leaguePositions(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var positions []LeaguePosition
	err := gorrito.leaguePositions(&positions, "RANKED_SOLO_5x5", "GOLD", "I", "TOP", 0)
	fmt.Println(err, positions)

}
func Test_leaguePositionsBySummoner(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var positions []LeaguePosition
	err := gorrito.leaguePositionsBySummoner(&positions, "N6sJUal26619EiPeJBaJkw_QTEZsi5uL_ddPlANsP4xM--xZ")
	fmt.Println(err, positions)

}

func Test_spectatorFeaturedGames(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var featuredGames FeaturedGames
	err := gorrito.spectatorFeaturedGames(&featuredGames)
	fmt.Println(err, featuredGames)

}

func Test_spectatorActiveGames(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var currentGameInfo CurrentGameInfo
	err := gorrito.spectatorActiveGames(&currentGameInfo, "M6heuV0xaGxDMMhc0fcgWPSFn2Rxf6KufwtdPvtI5bA1P7zt")
	fmt.Println(err, currentGameInfo)

}

func Test_LolStatus(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var shardStatus ShardStatus
	err := gorrito.lolStatus(&shardStatus)
	fmt.Println(err, shardStatus)

}
func Test_matchByID(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var match Match
	err := gorrito.matchByID(&match, 3992336664)
	fmt.Println(err, match)

}

func Test_matchTimeline(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var matchTimeline MatchTimeline
	err := gorrito.matchTimeline(&matchTimeline, 3992336664)
	fmt.Println(err, matchTimeline)

}

func Test_matchesByAccountID(t *testing.T) {
	gorrito := Gorrito{region: region, token: token}

	var matchList MatchList

	err := gorrito.matchesByAccountID(&matchList, "6ormLWaHtYFYUzdwGN7Bhqv1uw_Q8xfqTHMukTe_Bw59hTs", 15, 0, 7, 0, 0, 0, 0)
	fmt.Println(err, matchList)

}
