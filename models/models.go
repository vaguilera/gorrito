package models

type Champions struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

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
	Rank      string `json:"Rank"`
	LeagueItem
}

type Mastery struct {
	ChampionLevel                int    `json:"championLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionPoints               int    `json:"championPoints"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	SummonerID                   string `json:"summonerId"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChampionID                   int    `json:"championId"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
}

type Match struct {
	Metadata MatchMetaData `json:"metadata"`
	Info     MatchInfo     `json:"info"`
}

type MatchMetaData struct {
	DataVersion  string   `json:"dataVersion"`
	MatchId      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type MatchInfo struct {
	GameDuration       int64              `json:"gameDuration"`
	GameCreation       int64              `json:"gameCreation"`
	GameEndTimestamp   int64              `json:"gameEndTimestamp"`
	GameId             int64              `json:"gameId"`
	GameMode           string             `json:"gameMode"`
	GameName           string             `json:"gameName"`
	GameStartTimestamp int64              `json:"gameStartTimestamp"`
	GameType           string             `json:"gameType"`
	GameVersion        string             `json:"gameVersion"`
	MapId              int                `json:"gameMapId"`
	Participants       []MatchParticipant `json:"participants"`
	PlatformId         string             `json:"platformId"`
	QueueId            int                `json:"queueId"`
	Teams              []MatchTeam        `json:"teams"`
	TournamentCode     string             `json:"tournamentCode"`
}

type MatchParticipant struct {
	Assists                     int    `json:"assists"`
	BaronKills                  int    `json:"baronKills"`
	BountyLevel                 int    `json:"bountyLevel"`
	ChampExperience             int    `json:"champExperience"`
	ChampLevel                  int    `json:"champLevel"`
	ChampionId                  int    `json:"championId"`
	ChampionName                string `json:"championName"`
	ChampionTransform           int    `json:"championTransform"`
	ConsumablesPurchased        int    `json:"consumablesPurchased"`
	DamageDealtToBuildings      int    `json:"damageDealtToBuildings"`
	DamageDealtToObjectives     int    `json:"damageDealtToObjectives"`
	DamageDealtToTurrets        int    `json:"damageDealtToTurrets"`
	DamageSelfMitigated         int    `json:"damageSelfMitigated"`
	Deaths                      int    `json:"deaths"`
	DetectorWardsPlaced         int    `json:"detectorWardsPlaced"`
	DoubleKills                 int    `json:"doubleKills"`
	DragonKills                 int    `json:"dragonKills"`
	FirstBloodAssist            bool   `json:"firstBloodAssist"`
	FirstBloodKill              bool   `json:"firstBloodKill"`
	FirstTowerAssist            bool   `json:"firstTowerAssist"`
	FirstTowerKill              bool   `json:"firstTowerKill"`
	GameEndedInEarlySurrender   bool   `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender        bool   `json:"gameEndedInSurrender"`
	GoldEarned                  int    `json:"goldEarned"`
	GoldSpent                   int    `json:"goldSpent"`
	IndividualPosition          string `json:"individualPosition"`
	InhibitorKills              int    `json:"inhibitorKills"`
	InhibitorTakedowns          int    `json:"inhibitorTakedowns"`
	InhibitorsLost              int    `json:"inhibitorsLost"`
	Item0                       int    `json:"item0"`
	Item1                       int    `json:"item1"`
	Item2                       int    `json:"item2"`
	Item3                       int    `json:"item3"`
	Item4                       int    `json:"item4"`
	Item5                       int    `json:"item5"`
	Item6                       int    `json:"item6"`
	ItemsPurchased              int    `json:"itemsPurchased"`
	KillingSprees               int    `json:"killingSprees"`
	Kills                       int    `json:"kills"`
	Lane                        string `json:"lane"`
	LargestCriticalStrike       int    `json:"largestCriticalStrike"`
	LargestKillingSpree         int    `json:"largestKillingSpree"`
	LargestMultiKill            int    `json:"largestMultiKill"`
	LongestTimeSpentLiving      int    `json:"longestTimeSpentLiving"`
	MagicDamageDealt            int    `json:"magicDamageDealt"`
	MagicDamageDealtToChampions int    `json:"magicDamageDealtToChampions"`
	MagicDamageTaken            int    `json:"magicDamageTaken"`
	NeutralMinionsKilled        int    `json:"neutralMinionsKilled"`
	NexusKills                  int    `json:"nexusKills"`
	NexusTakedowns              int    `json:"nexusTakedowns"`
	NexusLost                   int    `json:"nexusLost"`
	ObjectivesStolen            int    `json:"objectivesStolen"`
	ObjectivesStolenAssists     int    `json:"objectivesStolenAssists"`
	ParticipantId               int    `json:"participantId"`
	PentaKills                  int    `json:"pentaKills"`
	Perks                       struct {
		StatPerks struct {
			Defense int `json:"defense"`
			Flex    int `json:"flex"`
			Offense int `json:"offense"`
		} `json:"statPerks"`
		Styles []struct {
			Description string `json:"description"`
			Style       int    `json:"style"`
			Selections  []struct {
				Perk int `json:"perk"`
				Var1 int `json:"var1"`
				Var2 int `json:"var2"`
				Var3 int `json:"var3"`
			} `json:"selections"`
		} `json:"styles"`
	} `json:"perks"`
	PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
	ProfileIcon                    int    `json:"profileIcon"`
	Puuid                          string `json:"puuid"`
	QuadraKills                    int    `json:"quadraKills"`
	RiotIdName                     string `json:"riotIdName"`
	RiotIdTagline                  string `json:"riotIdTagline"`
	Role                           string `json:"role"`
	SightWardsBoughtInGame         int    `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int    `json:"spell1Casts"`
	Spell2Casts                    int    `json:"spell2Casts"`
	Spell3Casts                    int    `json:"spell3Casts"`
	Spell4Casts                    int    `json:"spell4Casts"`
	Summoner1Casts                 int    `json:"summoner1Casts"`
	Summoner1Id                    int    `json:"summoner1Id"`
	Summoner2Casts                 int    `json:"summoner2Casts"`
	Summoner2Id                    int    `json:"summoner2Id"`
	SummonerId                     string `json:"summonerId"`
	SummonerLevel                  int    `json:"summonerLevel"`
	SummonerName                   string `json:"summonerName"`
	TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
	TeamId                         int    `json:"teamId"`
	TeamPosition                   string `json:"teamPosition"`
	TimeCCingOthers                int    `json:"timeCCingOthers"`
	TimePlayed                     int    `json:"timePlayed"`
	TotalDamageDealt               int    `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int    `json:"totalDamageTaken"`
	TotalHeal                      int    `json:"totalHeal"`
	TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int    `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int    `json:"totalUnitsHealedv"`
	TripleKills                    int    `json:"tripleKills"`
	TrueDamageDealt                int    `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int    `json:"trueDamageTaken"`
	TurretKills                    int    `json:"turretKills"`
	TurretTakedowns                int    `json:"turretTakedowns"`
	TurretsLost                    int    `json:"turretsLost"`
	UnrealKills                    int    `json:"unrealKills"`
	VisionScore                    int    `json:"visionScore"`
	VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
	WardsKilled                    int    `json:"wardsKilled"`
	WardsPlaced                    int    `json:"wardsPlacedv"`
	Win                            bool   `json:"win"`
}

type MatchTeam struct {
	Bans []struct {
		PickTurn   int `json:"pickTurn"`
		ChampionID int `json:"championId"`
	} `json:"bans"`
	Objectives struct {
		Baron      MatchObjective `json:"baron"`
		Champion   MatchObjective `json:"champion"`
		Dragon     MatchObjective `json:"dragon"`
		Inhibitor  MatchObjective `json:"inhibitor"`
		RiftHerald MatchObjective `json:"riftHerald"`
		Tower      MatchObjective `json:"tower"`
	} `json:"objectives"`
	TeamId int  `json:"teamId"`
	Win    bool `json:"win"`
}

type MatchObjective struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
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

type MatchesList []string

type content struct {
	Locale  string `json:"locale"`
	Content string `json:"content"`
}

type status struct {
	ID                int       `json:"id"`
	MaintenanceStatus string    `json:"maintenance_status"`
	IncidentSeverity  string    `json:"incident_severity"`
	Titles            []content `json:"titles"`
	Updates           []struct {
		ID               int       `json:"id"`
		Author           string    `json:"author"`
		Publish          bool      `json:"publish"`
		PublishLocations []string  `json:"publish_locations"`
		Translations     []content `json:"translations"`
		CreatedAt        string    `json:"created_at"`
		UpdatedAt        string    `json:"updated_at"`
	} `json:"updates"`
	CreatedAt string   `json:"created_at"`
	ArchiveAt string   `json:"archive_at"`
	UpdatedAt string   `json:"updated_at"`
	Platforms []string `json:"platforms"`
}

type ShardStatus struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Locales      []string `json:"locales"`
	Maintenances []status `json:"maintenances"`
	Incidents    []status `json:"incidents"`
}

type Summoner struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"Name"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
	AccountID     string `json:"accountId"`
	ID            string `json:"id"`
	RevisionDate  int64  `json:"revisionDate"`
}

type MatchTimeLine struct {
	Metadata MatchMetaData     `json:"metadata"`
	Info     MatchTimeLineInfo `json:"info"`
}

type MatchTimeLineInfo struct {
	FrameInterval int                  `json:"frameInterval"`
	Frames        []MatchTimeLineFrame `json:"frames"`
	GameId        int                  `json:"gameId"`
	Participants  []struct {
		ParticipantId int    `json:"participantId"`
		Puuid         string `json:"puuid"`
	}
}

type MatchTimeLineFrame struct {
	events []struct {
		RealTimestamp int64  `json:"realTimestamp"`
		Timestamp     int64  `json:"timestamp"`
		Type          string `json:"type"`
	}
	ParticipantFrames struct {
		One   ParticipantFrame `json:"1"`
		Two   ParticipantFrame `json:"2"`
		Three ParticipantFrame `json:"3"`
		Four  ParticipantFrame `json:"4"`
		Five  ParticipantFrame `json:"5"`
		Six   ParticipantFrame `json:"6"`
		Seven ParticipantFrame `json:"7"`
		Eight ParticipantFrame `json:"8"`
		Nine  ParticipantFrame `json:"9"`
		Ten   ParticipantFrame `json:"10"`
	} `json:"participantFrames"`
	Timestamp int64 `json:"timestamp"`
}

type ParticipantFrame struct {
	ChampionStats struct {
		AbilityHaste         int `json:"abilityHaste"`
		AbilityPower         int `json:"abilityPower"`
		Armor                int `json:"armor"`
		ArmorPen             int `json:"armorPen"`
		ArmorPenPercent      int `json:"armorPenPercent"`
		AttackDamage         int `json:"attackDamage"`
		AttackSpeed          int `json:"attackSpeed"`
		BonusArmorPenPercent int `json:"bonusArmorPenPercent"`
		BonusMagicPenPercent int `json:"bonusMagicPenPercent"`
		CcReduction          int `json:"ccReduction"`
		CoolDownReduction    int `json:"cooldownReduction"`
		Health               int `json:"health"`
		HealthMax            int `json:"healthMax"`
		HealthRegen          int `json:"healthRegen"`
		Lifesteal            int `json:"lifesteal"`
		MagicPen             int `json:"magicPen"`
		MagicPenPercent      int `json:"magicPenPercent"`
		MagicResist          int `json:"magicResist"`
		MovementSpeed        int `json:"movementSpeed"`
		Omnivamp             int `json:"omnivamp"`
		PhysicalVamp         int `json:"physicalVamp"`
		Power                int `json:"power"`
		PowerMax             int `json:"powerMax"`
		PowerRegen           int `json:"powerRegen"`
		SpellVamp            int `json:"spellVamp"`
	} `json:"championStats"`
	CurrentGold int `json:"currentGold"`
	DamageStats struct {
		MagicDamageDone               int `json:"magicDamageDone"`
		MagicDamageDoneToChampions    int `json:"magicDamageDoneToChampions"`
		MagicDamageTaken              int `json:"magicDamageTaken"`
		PhysicalDamageDone            int `json:"physicalDamageDone"`
		PhysicalDamageDoneToChampions int `json:"physicalDamageDoneToChampions"`
		PhysicalDamageTaken           int `json:"physicalDamageTaken"`
		TotalDamageDone               int `json:"totalDamageDone"`
		TotalDamageDoneToChampions    int `json:"totalDamageDoneToChampions"`
		TotalDamageTaken              int `json:"totalDamageTaken"`
		TrueDamageDone                int `json:"trueDamageDone"`
		TrueDamageDoneToChampions     int `json:"trueDamageDoneToChampions"`
		TrueDamageTaken               int `json:"trueDamageTaken"`
	} `json:"damageStats"`
	GoldPerSecond       int `json:"goldPerSecond"`
	JungleMinionsKilled int `json:"jungleMinionsKilled"`
	Level               int `json:"level"`
	MinionsKilled       int `json:"minionsKilled"`
	ParticipantId       int `json:"participantId"`
	Position            struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
	TimeEnemySpentControlled int `json:"timeEnemySpentControlled"`
	TotalGold                int `json:"totalGold"`
	Xp                       int `json:"xp"`
}

type CurrentGameInfo struct {
	GameID            int64  `json:"gameId"`
	GameStartTime     int64  `json:"gameStartTime"`
	PlatformID        string `json:"platformId"`
	GameMode          string `json:"gameMode"`
	MapID             int    `json:"mapId"`
	GameType          string `json:"gameType"`
	GameQueueConfigID int    `json:"gameQueueConfigId"`
	Observers         struct {
		EncryptionKey string `json:"encryptionKey"`
	} `json:"observers"`
	Participants    []ParticipantGameInfo `json:"participants"`
	GameLength      int                   `json:"gameLength"`
	BannedChampions []struct {
		TeamID     int   `json:"teamId"`
		ChampionID int64 `json:"championId"`
		PickTurn   int64 `json:"pickTurn"`
	} `json:"bannedChampions"`
}

type ParticipantGameInfo struct {
	ProfileIconID            int    `json:"profileIconId"`
	ChampionID               int    `json:"championId"`
	SummonerName             string `json:"summonerName"`
	GameCustomizationObjects []struct {
		Category string `json:"category"`
		Content  string `json:"content"`
	} `json:"gameCustomizationObjects"`
	Bot   bool `json:"bot"`
	Perks struct {
		PerkStyle    int64   `json:"perkStyle"`
		PerkIds      []int64 `json:"perkIds"`
		PerkSubStyle int64   `json:"perkSubStyle"`
	} `json:"perks"`
	Spell2ID   int64  `json:"spell2Id"`
	TeamID     int64  `json:"teamId"`
	Spell1ID   int64  `json:"spell1Id"`
	SummonerID string `json:"summonerId"`
}

type FeaturedGames struct {
	ClientRefreshInterval int64 `json:"clientRefreshInterval"`
	GameList              []struct {
		GameID            int64  `json:"gameId"`
		GameStartTime     int64  `json:"gameStartTime"`
		PlatformID        string `json:"platformId"`
		GameMode          string `json:"gameMode"`
		MapID             int64  `json:"mapId"`
		GameType          string `json:"gameType"`
		GameQueueConfigID int64  `json:"gameQueueConfigId"`
		Observers         struct {
			EncryptionKey string `json:"encryptionKey"`
		} `json:"observers"`
		Participants    []ParticipantFeatured `json:"participants"`
		GameLength      int64                 `json:"gameLength"`
		BannedChampions []struct {
			TeamID     int   `json:"teamId"`
			ChampionID int64 `json:"championId"`
			PickTurn   int64 `json:"pickTurn"`
		} `json:"bannedChampions"`
	} `json:"gameList"`
}

type ParticipantFeatured struct {
	ProfileIconID int    `json:"profileIconId"`
	ChampionID    int    `json:"championId"`
	SummonerName  string `json:"summonerName"`
	Bot           bool   `json:"bot"`
	Spell2ID      int64  `json:"spell2Id"`
	TeamID        int64  `json:"teamId"`
	Spell1ID      int64  `json:"spell1Id"`
}
