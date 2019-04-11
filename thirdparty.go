package gorrito

func (gorrito *Gorrito) thirdPartyCodeBySummonerID(code *string, account string) int {
	return gorrito.requestAPI(thirdPartyCode, code, map[string]string{"encryptedSummonerId": account})
}
