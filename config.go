package gorrito

const (
	UriSummonerByAccountID  = "/lol/summoner/v4/summoners/by-account/{encryptedAccountId}"
	UriSummonerByName       = "/lol/summoner/v4/summoners/by-name/{summonerName}666"
	UriSummonerByPuuid      = "/lol/summoner/v4/summoners/by-puuid/{encryptedPUUID}"
	UriSummonerBySummonerID = "/lol/summoner/v4/summoners/{encryptedSummonerId}"

	UriChampionMasteryBySummonerID           = "/lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}"
	UriChampionMasteryByBySummonerIDChampion = "/lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}/by-champion/{championId}"
	UriChampionMasteryScoresBySummonerID     = "/lol/champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}"

	UriChampionRotations = "/lol/platform/v3/champion-rotations"

	UriLeagueChallengerByQueue  = "/lol/league/v4/challengerleagues/by-queue/{queue}"
	UriLeagueGrandMasterByQueue = "/lol/league/v4/grandmasterleagues/by-queue/{queue}"
	UriLeagueMasterByQueue      = "/lol/league/v4/masterleagues/by-queue/{queue}"
	UriLeagueLeagues            = "/lol/league/v4/leagues/{leagueId}"
	UriLeagueBySummonerID       = "/lol/league/v4/entries/by-summoner/{encryptedSummonerId}"
	UriLeagueQueueTierDivision  = "/lol/league/v4/entries/{queue}/{tier}/{division}"

	UriSpectatorActiveGames   = "/lol/spectator/v4/active-games/by-summoner/{encryptedSummonerId}"
	UriSpectatorFeaturedGames = "/lol/spectator/v4/featured-games"

	UriLolStatus = "/lol/status/v4/platform-data"

	UriMatchByID     = "/lol/match/v5/matches/{matchId}"
	UriMatchByPuuid  = "/lol/match/v5/matches/by-puuid/{puuid}/ids"
	UriMatchTimeline = "/lol/match/v5/matches/{matchId}/timeline"
)

var regions = map[string]string{
	"BR":   "br1.api.riotgames.com",
	"EUNE": "eun1.api.riotgames.com",
	"EUW":  "euw1.api.riotgames.com",
	"JP":   "jp1.api.riotgames.com",
	"KR":   "kr.api.riotgames.com",
	"LAN":  "la1.api.riotgames.com",
	"LAS":  "la2.api.riotgames.com",
	"NA":   "na1.api.riotgames.com",
	"OCE":  "oc1.api.riotgames.com",
	"TR":   "tr1.api.riotgames.com",
	"RU":   "ru.api.riotgames.com",
	"PBE":  "pbe1.api.riotgames.com",
}

var regionsV5 = map[string]string{
	"EUROPE":   "europe.api.riotgames.com",
	"AMERICAS": "americas.api.riotgames.com",
	"ASIA":     "asia.api.riotgames.com",
}

const (
	RANKED_SOLO_5x5 = "RANKED_SOLO_5x5"
	RANKED_FLEX_SR  = "RANKED_FLEX_SR"
	RANKED_FLEX_TT  = "RANKED_FLEX_TT"
)
