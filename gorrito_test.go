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

const token = "RGAPI-a005a943-5df9-43f0-8738-23c9b40dc8d8"
const region = "EUW"

// func Test_summonersByName(t *testing.T) {
// 	gorrito := NewClient(token, &Config{Region: region})

// 	summoner, err := gorrito.SummonerByName("shötgan")
// 	if err != nil {
// 		t.Errorf("error %s", err)
// 	} else {
// 		if summoner.Puuid != "NYnSHHSBy0zRc6eN0JllMRYSJG0U8F3006zs3jROOo9t_FxzvWAa4l3WncWuevM8N5FwgoNyr41ocQ" {
// 			t.Errorf("got %s, want %s", summoner.Puuid, "NYnSHHSBy0zRc6eN0JllMRYSJG0U8F3006zs3jROOo9t_FxzvWAa4l3WncWuevM8N5FwgoNyr41ocQ")
// 		}
// 	}
// }

// func Test_championMasteryBySummonerID(t *testing.T) {
// 	gorrito := NewClient(token, &Config{Region: region})

// 	masteries, err := gorrito.championMasteryBySummonerID("oUKkIFPFei6Li6PeXDRCnjesF3OMZQWrVH51Z7tqV7lFhD-R")
// 	fmt.Println(err, masteries)
// }

// func Test_championMasteryByBySummonerIDChampion(t *testing.T) {
// 	gorrito := Gorrito{region: region, token: token}

// 	var mastery Mastery
// 	err := gorrito.championMasteryByBySummonerIDChampion(&mastery, "id5l4_4JjOEOn3VpCWdj5Cb4s5lj2nmP5BTr-0AempQccA0", 110)
// 	fmt.Println(err, mastery)
// }

// func Test_championMasteryScoresBySummonerID(t *testing.T) {
// 	gorrito := NewClient(token, &Config{Region: region})

// 	score, err := gorrito.championMasteryScoresBySummonerID("oUKkIFPFei6Li6PeXDRCnjesF3OMZQWrVH51Z7tqV7lFhD-R")
// 	fmt.Println(err, score)
// }

// func Test_championRotations(t *testing.T) {
// 	gorrito := NewClient(token, &Config{Region: region})

// 	champions, err := gorrito.championRotations()
// 	fmt.Println(err, champions)
// }

func Test_leaguePositions(t *testing.T) {
	gorrito := NewClient(token, &Config{Region: region})

	queue, err := gorrito.LeagueLeaguesByQueueTierDivision("RANKED_SOLO_5x5", "GOLD", "II")
	fmt.Println(err, queue)

}

// func Test_leaguePositions(t *testing.T) {
// 	gorrito := Gorrito{region: region, token: token}

// 	var positions []LeaguePosition
// 	err := gorrito.leaguePositions(&positions, "RANKED_SOLO_5x5", "GOLD", "I", "TOP", 0)
// 	fmt.Println(err, positions)

// }
// func Test_leaguePositionsBySummoner(t *testing.T) {
// 	gorrito := Gorrito{region: region, token: token}

// 	var positions []LeaguePosition
// 	err := gorrito.leaguePositionsBySummoner(&positions, "N6sJUal26619EiPeJBaJkw_QTEZsi5uL_ddPlANsP4xM--xZ")
// 	fmt.Println(err, positions)

// }

// func Test_spectatorFeaturedGames(t *testing.T) {
// 	gorrito := NewClient(token, &Config{Region: region})

// 	featuredGames, err := gorrito.spectatorFeaturedGames()
// 	fmt.Println(err, featuredGames)

// }

// func Test_spectatorActiveGames(t *testing.T) {
// 	gorrito := NewClient(token, &Config{Region: region})

// 	currentGameInfo, err := gorrito.SpectatorActiveGames("Wb_WHNe-stXP_9QRHYTlQfggrNvDKJnFRL8M6SWXbS1Kv7Gpopk_5gF1-g")
// 	fmt.Println(err, currentGameInfo)

// }

// func Test_LolStatus(t *testing.T) {
// 	gorrito := NewClient(token, &Config{Region: region})

// 	shard, err := gorrito.lolStatus()
// 	fmt.Println(err, shard)

// }

// func Test_matchByID(t *testing.T) {
// 	gorrito := Gorrito{region: region, token: token}

// 	var match Match
// 	err := gorrito.matchByID(&match, 3992336664)
// 	fmt.Println(err, match)

// }

// func Test_matchTimeline(t *testing.T) {
// 	gorrito := NewClient(token, &Config{Region: region})

// 	timeline, _ := gorrito.matchTimeline("EUW1_5506291862")
// 	// fmt.Println(err, timeline)

// 	fmt.Println(timeline.Info.Frames[0].ParticipantFrames.Two.ChampionStats.Health)

// }

// func Test_matchesByPuuid(t *testing.T) {

// 	gorrito := NewClient(token, &Config{Region: region})

// 	options := gorrito.NewMatchesByPuuidOptions().Queue(420).Type("ranked")

// 	score, err := gorrito.MatchesByPuuid("3mEtBCOupCYdPmxVjecBh_B0l4YELWZMK5vohgvxRmcNM3uDVOUqT7QsChCR1OM_26tTK-q9MHBSog", options)
// 	fmt.Println(err, score)

// }
