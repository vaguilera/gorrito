package gorrito

import (
	"bytes"
	"strconv"
)

type Match struct {
	SeasonID              int   `json:"seasonId"`
	QueueID               int   `json:"queueId"`
	GameID                int64 `json:"gameId"`
	ParticipantIdentities []struct {
		Player struct {
			CurrentPlatformID string `json:"currentPlatformId"`
			SummonerName      string `json:"summonerName"`
			MatchHistoryURI   string `json:"matchHistoryUri"`
			PlatformID        string `json:"platformId"`
			CurrentAccountID  string `json:"currentAccountId"`
			ProfileIcon       int    `json:"profileIcon"`
			SummonerID        string `json:"summonerId"`
			AccountID         string `json:"accountId"`
		} `json:"player"`
		ParticipantID int `json:"participantId"`
	} `json:"participantIdentities"`
	GameVersion string `json:"gameVersion"`
	PlatformID  string `json:"platformId"`
	GameMode    string `json:"gameMode"`
	MapID       int    `json:"mapId"`
	GameType    string `json:"gameType"`
	Teams       []struct {
		FirstDragon bool `json:"firstDragon"`
		Bans        []struct {
			PickTurn   int `json:"pickTurn"`
			ChampionID int `json:"championId"`
		} `json:"bans"`
		FirstInhibitor       bool   `json:"firstInhibitor"`
		Win                  string `json:"win"`
		FirstRiftHerald      bool   `json:"firstRiftHerald"`
		FirstBaron           bool   `json:"firstBaron"`
		BaronKills           int    `json:"baronKills"`
		RiftHeraldKills      int    `json:"riftHeraldKills"`
		FirstBlood           bool   `json:"firstBlood"`
		TeamID               int    `json:"teamId"`
		FirstTower           bool   `json:"firstTower"`
		VilemawKills         int    `json:"vilemawKills"`
		InhibitorKills       int    `json:"inhibitorKills"`
		TowerKills           int    `json:"towerKills"`
		DominionVictoryScore int    `json:"dominionVictoryScore"`
		DragonKills          int    `json:"dragonKills"`
	} `json:"teams"`
	Participants []struct {
		Spell1ID      int `json:"spell1Id"`
		ParticipantID int `json:"participantId"`
		Timeline      struct {
			Lane                        string             `json:"lane"`
			ParticipantID               int                `json:"participantId"`
			CsDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
			GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
			CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
			XpPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
			Role                        string             `json:"role"`
			DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
			DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
		} `json:"timeline"`
		Spell2ID int `json:"spell2Id"`
		TeamID   int `json:"teamId"`
		Stats    struct {
			NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
			VisionScore                     int  `json:"visionScore"`
			MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
			LargestMultiKill                int  `json:"largestMultiKill"`
			TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
			LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
			Perk1Var1                       int  `json:"perk1Var1"`
			Perk1Var3                       int  `json:"perk1Var3"`
			Perk1Var2                       int  `json:"perk1Var2"`
			TripleKills                     int  `json:"tripleKills"`
			Perk5                           int  `json:"perk5"`
			Perk4                           int  `json:"perk4"`
			PlayerScore9                    int  `json:"playerScore9"`
			PlayerScore8                    int  `json:"playerScore8"`
			Kills                           int  `json:"kills"`
			PlayerScore1                    int  `json:"playerScore1"`
			PlayerScore0                    int  `json:"playerScore0"`
			PlayerScore3                    int  `json:"playerScore3"`
			PlayerScore2                    int  `json:"playerScore2"`
			PlayerScore5                    int  `json:"playerScore5"`
			PlayerScore4                    int  `json:"playerScore4"`
			PlayerScore7                    int  `json:"playerScore7"`
			PlayerScore6                    int  `json:"playerScore6"`
			Perk5Var1                       int  `json:"perk5Var1"`
			Perk5Var3                       int  `json:"perk5Var3"`
			Perk5Var2                       int  `json:"perk5Var2"`
			TotalScoreRank                  int  `json:"totalScoreRank"`
			NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
			StatPerk1                       int  `json:"statPerk1"`
			StatPerk0                       int  `json:"statPerk0"`
			DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
			PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
			DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
			Perk2Var2                       int  `json:"perk2Var2"`
			Perk2Var3                       int  `json:"perk2Var3"`
			TotalUnitsHealed                int  `json:"totalUnitsHealed"`
			Perk2Var1                       int  `json:"perk2Var1"`
			Perk4Var1                       int  `json:"perk4Var1"`
			TotalDamageTaken                int  `json:"totalDamageTaken"`
			Perk4Var3                       int  `json:"perk4Var3"`
			WardsKilled                     int  `json:"wardsKilled"`
			LargestCriticalStrike           int  `json:"largestCriticalStrike"`
			LargestKillingSpree             int  `json:"largestKillingSpree"`
			QuadraKills                     int  `json:"quadraKills"`
			MagicDamageDealt                int  `json:"magicDamageDealt"`
			FirstBloodAssist                bool `json:"firstBloodAssist"`
			Item2                           int  `json:"item2"`
			Item3                           int  `json:"item3"`
			Item0                           int  `json:"item0"`
			Item1                           int  `json:"item1"`
			Item6                           int  `json:"item6"`
			Item4                           int  `json:"item4"`
			Item5                           int  `json:"item5"`
			Perk1                           int  `json:"perk1"`
			Perk0                           int  `json:"perk0"`
			Perk3                           int  `json:"perk3"`
			Perk2                           int  `json:"perk2"`
			Perk3Var3                       int  `json:"perk3Var3"`
			Perk3Var2                       int  `json:"perk3Var2"`
			Perk3Var1                       int  `json:"perk3Var1"`
			DamageSelfMitigated             int  `json:"damageSelfMitigated"`
			MagicalDamageTaken              int  `json:"magicalDamageTaken"`
			Perk0Var2                       int  `json:"perk0Var2"`
			FirstInhibitorKill              bool `json:"firstInhibitorKill"`
			TrueDamageTaken                 int  `json:"trueDamageTaken"`
			Assists                         int  `json:"assists"`
			Perk4Var2                       int  `json:"perk4Var2"`
			GoldSpent                       int  `json:"goldSpent"`
			TrueDamageDealt                 int  `json:"trueDamageDealt"`
			ParticipantID                   int  `json:"participantId"`
			PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
			SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
			TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
			PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
			TotalPlayerScore                int  `json:"totalPlayerScore"`
			Win                             bool `json:"win"`
			ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
			TotalDamageDealt                int  `json:"totalDamageDealt"`
			NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
			Deaths                          int  `json:"deaths"`
			WardsPlaced                     int  `json:"wardsPlaced"`
			PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
			PerkSubStyle                    int  `json:"perkSubStyle"`
			TurretKills                     int  `json:"turretKills"`
			FirstBloodKill                  bool `json:"firstBloodKill"`
			TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
			GoldEarned                      int  `json:"goldEarned"`
			KillingSprees                   int  `json:"killingSprees"`
			UnrealKills                     int  `json:"unrealKills"`
			FirstTowerAssist                bool `json:"firstTowerAssist"`
			FirstTowerKill                  bool `json:"firstTowerKill"`
			ChampLevel                      int  `json:"champLevel"`
			DoubleKills                     int  `json:"doubleKills"`
			InhibitorKills                  int  `json:"inhibitorKills"`
			FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
			Perk0Var1                       int  `json:"perk0Var1"`
			CombatPlayerScore               int  `json:"combatPlayerScore"`
			Perk0Var3                       int  `json:"perk0Var3"`
			VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
			PentaKills                      int  `json:"pentaKills"`
			TotalHeal                       int  `json:"totalHeal"`
			TotalMinionsKilled              int  `json:"totalMinionsKilled"`
			TimeCCingOthers                 int  `json:"timeCCingOthers"`
			StatPerk2                       int  `json:"statPerk2"`
		} `json:"stats"`
		ChampionID                int    `json:"championId"`
		HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
	} `json:"participants"`
	GameDuration int64 `json:"gameDuration"`
	GameCreation int64 `json:"gameCreation"`
}

type MatchList struct {
	Matches []struct {
		Lane       string `json:"lane"`
		GameID     int64  `json:"gameId"`
		Champion   int    `json:"champion"`
		PlatformID string `json:"platformId"`
		Timestamp  int64  `json:"timestamp"`
		Queue      int    `json:"queue"`
		Role       string `json:"role"`
		Season     int    `json:"season"`
	} `json:"matches"`
	EndIndex   int `json:"endIndex"`
	StartIndex int `json:"startIndex"`
	TotalGames int `json:"totalGames"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MatchParticipantFrame struct {
	TotalGold           int      `json:"totalGold"`
	TeamScore           int      `json:"teamScore"`
	ParticipantID       int      `json:"participantId"`
	Level               int      `json:"level"`
	CurrentGold         int      `json:"currentGold"`
	MinionsKilled       int      `json:"minionsKilled"`
	DominionScore       int      `json:"dominionScore"`
	Position            Position `json:"position"`
	XP                  int      `json:"xp"`
	JungleMinionsKilled int      `json:"jungleMinionsKilled"`
}

type MatchTimeline struct {
	Frames []struct {
		Timestamp         int64                            `json:"timestamp"`
		ParticipantFrames map[string]MatchParticipantFrame `json:"participantFrames"`
		Events            []struct {
			EventType               string   `json:"eventType"`
			TowerType               string   `json:"towerType"`
			TeamID                  int      `json:"teamId"`
			AscendedType            string   `json:"ascendedType"`
			KillerID                int      `json:"killerId"`
			LevelUpType             string   `json:"levelUpType"`
			PointCaptured           string   `json:"pointCaptured"`
			AssistingParticipantIds []int    `json:"assistingParticipantIds"`
			WardType                string   `json:"wardType"`
			MonsterType             string   `json:"monsterType"`
			Type                    string   `json:"type"`
			SkillSlot               int      `json:"skillSlot"`
			VictimID                int      `json:"victimId"`
			Timestamp               int64    `json:"timestamp"`
			AfterID                 int      `json:"afterId"`
			MonsterSubType          string   `json:"monsterSubType"`
			LaneType                string   `json:"laneType"`
			ItemID                  int      `json:"itemId"`
			ParticipantID           int      `json:"participantId"`
			BuildingType            string   `json:"buildingType"`
			CreatorID               int      `json:"creatorId"`
			Position                Position `json:"position"`
			BeforeID                int      `json:"beforeId"`
		} `json:"events"`
	} `json:"frames"`
	FrameInterval int64 `json:"frameInterval"`
}

func (gorrito *Gorrito) matchByID(match *Match, matchID int64) int {

	return gorrito.requestAPI(matchByID, match, map[string]string{"matchId": strconv.FormatInt(matchID, 10)})
}

func (gorrito *Gorrito) matchesByAccountID(matchList *MatchList, account string, champion int, queue int, season int, endTime int64, beginTime int64, endIndex int, beginIndex int) int {

	var parameters bytes.Buffer
	parameters.WriteString("?")

	if champion != 0 {
		parameters.WriteString("&champion=" + strconv.Itoa(champion))
	}
	if queue != 0 {
		parameters.WriteString("&queue=" + strconv.Itoa(queue))
	}

	if season != 0 {
		parameters.WriteString("&season=" + strconv.Itoa(season))
	}

	if endTime != 0 {
		parameters.WriteString("&endTime=" + strconv.FormatInt(endTime, 10))
	}
	if beginTime != 0 {
		parameters.WriteString("&beginTime=" + strconv.FormatInt(beginTime, 10))
	}
	if endIndex != 0 {
		parameters.WriteString("&endIndex=" + strconv.Itoa(endIndex))
	}
	if beginIndex != 0 {
		parameters.WriteString("&beginIndex=" + strconv.Itoa(beginIndex))
	}

	return gorrito.requestAPI(matchByAccountID, matchList, map[string]string{"encryptedAccountId": account}, parameters.String())
}

func (gorrito *Gorrito) matchTimeline(timeline *MatchTimeline, matchID int64) int {

	return gorrito.requestAPI(matchTimeline, timeline, map[string]string{"matchId": strconv.FormatInt(matchID, 10)})
}

func (gorrito *Gorrito) matchesByTournament(matches *[]int64, tournamentCode string) int {

	return gorrito.requestAPI(matchesByTournament, matches, map[string]string{"tournamentCode": tournamentCode})
}

func (gorrito *Gorrito) matchByTournament(match *Match, matchID int64, tournamentCode string) int {

	return gorrito.requestAPI(matchByTournament, match, map[string]string{"matchId": strconv.FormatInt(matchID, 10), "tournamentCode": tournamentCode})
}
