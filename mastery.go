package gorrito

import (
	"strconv"

	"github.com/vaguilera/gorrito/models"
)

func (c *Client) ChampionMasteryBySummonerID(summonerId string) ([]models.Mastery, error) {
	masteries := make([]models.Mastery, 0)
	body, err := c.requestAPI(UriChampionMasteryBySummonerID, map[string]string{"encryptedSummonerId": summonerId})
	if err != nil {
		return nil, err
	}

	err = c.unMarshall(body, masteries)
	return masteries, err
}

func (c *Client) ChampionMasteryByBySummonerIDChampion(summonerId string, champion int) (*models.Mastery, error) {
	mastery := models.Mastery{}

	body, err := c.requestAPI(UriChampionMasteryByBySummonerIDChampion, map[string]string{"encryptedSummonerId": summonerId, "championId": strconv.Itoa(champion)})
	if err != nil {
		return nil, err
	}
	err = c.unMarshall(body, &mastery)
	return &mastery, err
}

func (c *Client) ChampionMasteryScoresBySummonerID(summonerId string) (int, error) {
	var score int
	body, err := c.requestAPI(UriChampionMasteryScoresBySummonerID, map[string]string{"encryptedSummonerId": summonerId})
	if err != nil {
		return 0, err
	}

	err = c.unMarshall(body, &score)
	return score, err
}

func (c *Client) ChampionMasteryBySummonerIDRaw(summonerId string) (*string, error) {
	body, err := c.requestAPI(UriChampionMasteryBySummonerID, map[string]string{"encryptedSummonerId": summonerId})
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}

func (c *Client) ChampionMasteryByBySummonerIDChampionRaw(summonerId string, champion int) (*string, error) {
	body, err := c.requestAPI(UriChampionMasteryByBySummonerIDChampion, map[string]string{"encryptedSummonerId": summonerId, "championId": strconv.Itoa(champion)})
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}

func (c *Client) ChampionMasteryScoresBySummonerIDRaw(summonerId string) (*string, error) {
	body, err := c.requestAPI(UriChampionMasteryScoresBySummonerID, map[string]string{"encryptedSummonerId": summonerId})
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}
