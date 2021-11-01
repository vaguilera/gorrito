package gorrito

import (
	"github.com/vaguilera/gorrito/models"
)

func (c *Client) LeagueChallengerByQueue(queue string) (*models.LeagueList, error) {
	litem := models.LeagueList{}
	body, err := c.requestAPI(UriLeagueChallengerByQueue, map[string]string{"queue": queue})
	if err != nil {
		return nil, err
	}
	err = c.unMarshall(body, &litem)
	return &litem, err
}

func (c *Client) LeagueChallengerByQueueRaw(queue string) (*string, error) {

	body, err := c.requestAPI(UriLeagueChallengerByQueue, map[string]string{"queue": queue})
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}

func (c *Client) LeagueGrandMasterByQueue(leagueItem *models.LeagueItem, queue string) (*models.LeagueList, error) {
	litem := models.LeagueList{}
	body, err := c.requestAPI(UriLeagueGrandMasterByQueue, map[string]string{"queue": queue})
	if err != nil {
		return nil, err
	}
	err = c.unMarshall(body, &litem)
	return &litem, err
}

func (c *Client) LeagueGrandMasterByQueueRaw(leagueItem *models.LeagueItem, queue string) (*string, error) {
	body, err := c.requestAPI(UriLeagueGrandMasterByQueue, map[string]string{"queue": queue})
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}

func (c *Client) LeagueMasterByQueue(leagueItem *models.LeagueItem, queue string) (*models.LeagueList, error) {
	litem := models.LeagueList{}
	body, err := c.requestAPI(UriLeagueMasterByQueue, map[string]string{"queue": queue})
	if err != nil {
		return nil, err
	}
	err = c.unMarshall(body, &litem)
	return &litem, err
}

func (c *Client) LeagueMasterByQueueRaw(leagueItem *models.LeagueItem, queue string) (*string, error) {
	body, err := c.requestAPI(UriLeagueMasterByQueue, map[string]string{"queue": queue})
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}

func (c *Client) LeagueLeaguesById(leagueId string) (*models.LeagueList, error) {
	litem := models.LeagueList{}
	body, err := c.requestAPI(UriLeagueLeagues, map[string]string{"leagueId": leagueId})
	if err != nil {
		return nil, err
	}
	err = c.unMarshall(body, &litem)
	return &litem, err
}

func (c *Client) LeagueLeaguesByIdRaw(leagueId string) (*string, error) {
	body, err := c.requestAPI(UriLeagueLeagues, map[string]string{"leagueId": leagueId})
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}

func (c *Client) LeagueLeaguesBySummonerId(summonerId string) (*models.LeagueEntry, error) {
	litem := models.LeagueEntry{}
	body, err := c.requestAPI(UriLeagueBySummonerID, map[string]string{"encryptedSummonerId": summonerId})
	if err != nil {
		return nil, err
	}
	err = c.unMarshall(body, &litem)
	return &litem, err
}

func (c *Client) LeagueLeaguesBySummonerIdRaw(summonerId string) (*string, error) {
	body, err := c.requestAPI(UriLeagueBySummonerID, map[string]string{"encryptedSummonerId": summonerId})
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}

func (c *Client) LeagueLeaguesByQueueTierDivision(queue string, tier string, division string) ([]models.LeagueEntry, error) {
	litem := make([]models.LeagueEntry, 0)
	body, err := c.requestAPI(UriLeagueQueueTierDivision, map[string]string{"queue": queue, "tier": tier, "division": division})
	if err != nil {
		return nil, err
	}
	err = c.unMarshall(body, &litem)
	return litem, err
}

func (c *Client) LeagueLeaguesByQueueTierDivisionRaw(queue string, tier string, division string) (*string, error) {
	body, err := c.requestAPI(UriLeagueQueueTierDivision, map[string]string{"queue": queue, "tier": tier, "division": division})
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}
