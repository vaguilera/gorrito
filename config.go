package gorrito

const (
	summonerByAccountID  = "/lol/summoner/v4/summoners/by-account/{encryptedAccountId}"
	summonerByName       = "/lol/summoner/v4/summoners/by-name/{summonerName}"
	summonerByPuuid      = "/lol/summoner/v4/summoners/by-puuid/{encryptedPUUID}"
	summonerBySummonerID = "/lol/summoner/v4/summoners/{encryptedSummonerId}"

	championMasteryBySummonerID           = "/lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}"
	championMasteryByBySummonerIDChampion = "/lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}/by-champion/{championId}"
	championMasteryScoresBySummonerID     = "/lol/champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}"

	championRotations = "/lol/platform/v3/champion-rotations"

	leagueChallengerByQueue   = "/lol/league/v4/challengerleagues/by-queue/{queue}"
	leagueGrandMasterByQueue  = "/lol/league/v4/grandmasterleagues/by-queue/{queue}"
	leagueMasterByQueue       = "/lol/league/v4/masterleagues/by-queue/{queue}"
	leagueLeagues             = "/lol/league/v4/leagues/{leagueId}"
	leaguePositional          = "/lol/league/v4/positional-rank-queues"
	leaguePositionsBySummoner = "/lol/league/v4/positions/by-summoner/{encryptedSummonerId}"
	leaguePositions           = "/lol/league/v4/positions/{positionalQueue}/{tier}/{division}/{position}/{page}"

	spectatorActiveGames   = "/lol/spectator/v4/active-games/by-summoner/{encryptedSummonerId}"
	spectatorFeaturedGames = "/lol/spectator/v4/featured-games"

	thirdPartyCode = "/lol/platform/v4/third-party-code/by-summoner/{encryptedSummonerId}"
	lolStatus      = "/lol/status/v3/shard-data"

	matchByID           = "/lol/match/v4/matches/{matchId}"
	matchByAccountID    = "/lol/match/v4/matchlists/by-account/{encryptedAccountId}"
	matchTimeline       = "/lol/match/v4/timelines/by-match/{matchId}"
	matchesByTournament = "/lol/match/v4/matches/by-tournament-code/{tournamentCode}/ids"
	matchByTournament   = "/lol/match/v4/matches/{matchId}/by-tournament-code/{tournamentCode}"
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
