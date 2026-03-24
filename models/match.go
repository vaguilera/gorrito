package models

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
	MapId              int                `json:"mapId"`
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
	TotalUnitsHealed               int    `json:"totalUnitsHealed"`
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
	WardsPlaced                    int    `json:"wardsPlaced"`
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

type MatchesList []string

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
	} `json:"participants"`
}

type MatchTimeLineFrame struct {
	Events []struct {
		RealTimestamp int64  `json:"realTimestamp"`
		Timestamp     int64  `json:"timestamp"`
		Type          string `json:"type"`
	} `json:"events"`
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
