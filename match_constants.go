package gorrito

// Match type constants accepted by the match-v5 endpoint query parameter "type".
const (
	MatchTypeRanked   = "ranked"
	MatchTypeNormal   = "normal"
	MatchTypeTourney  = "tourney"
	MatchTypeTutorial = "tutorial"
)

// Queue ID constants from Riot game constants used by match-v5 filters.
const (
	QueueCustom                       = 0
	QueueNormalDraftPick              = 400
	QueueRankedSolo                   = 420
	QueueNormalBlindPick              = 430
	QueueRankedFlex                   = 440
	QueueAram                         = 450
	QueueBlindPick                    = 430
	QueueQuickplay                    = 490
	QueueClash                        = 700
	QueueCoopVsAIIntro                = 830
	QueueCoopVsAIBeginner             = 840
	QueueCoopVsAIIntermediate         = 850
	QueueURF                          = 900
	QueueARURF                        = 900
	QueueOneForAll                    = 1020
	QueueNexusBlitz                   = 1300
	QueueUltimateSpellbook            = 1400
	QueueArena                        = 1700
	QueueAramClash                    = 720
	QueueSummonersRiftClash           = 700
	QueueTeamfightTacticsNormal       = 1090
	QueueTeamfightTacticsRanked       = 1100
	QueueTeamfightTacticsTutorial     = 1110
	QueueTeamfightTacticsHyperRoll    = 1130
	QueueTeamfightTacticsDoubleUp     = 1160
	QueueTeamfightTacticsSoulBrawl    = 1180
	QueueTeamfightTacticsChonccs      = 1210
	QueueTeamfightTacticsTreasure     = 1220
	QueueTeamfightTacticsPenguParty   = 1230
)
