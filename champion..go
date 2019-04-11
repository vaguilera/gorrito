package gorrito

type Champions struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

func (gorrito *Gorrito) championRotations(champions *Champions) int {

	return gorrito.requestAPI(championRotations, champions, nil)
}
