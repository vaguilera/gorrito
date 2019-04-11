package gorrito

import "net/url"

type Summoner struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"Name"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
	AccountID     string `json:"accountId"`
	ID            string `json:"id"`
	RevisionDate  int64  `json:"revisionDate"`
}

// SummonerByAccountID ... Return summoner from encrypted account ID
func (gorrito *Gorrito) SummonerByAccountID(summoner *Summoner, account string) int {

	return gorrito.requestAPI(summonerByAccountID, summoner, map[string]string{"encryptedAccountId": account})
}

// SummonerByName ... Return summoner from encrypted account ID
func (gorrito *Gorrito) SummonerByName(summoner *Summoner, account string) int {

	return gorrito.requestAPI(summonerByName, summoner, map[string]string{"summonerName": url.QueryEscape(account)})
}

// SummonerByPuuid ... Return summoner from encrypted account ID
func (gorrito *Gorrito) SummonerByPuuid(summoner *Summoner, account string) int {

	return gorrito.requestAPI(summonerByPuuid, summoner, map[string]string{"encryptedPUUID": account})
}

// SummonerBySummonerID ... Return summoner from encrypted account ID
func (gorrito *Gorrito) SummonerBySummonerID(summoner *Summoner, account string) int {

	return gorrito.requestAPI(summonerBySummonerID, summoner, map[string]string{"encryptedSummonerId": account})

}
