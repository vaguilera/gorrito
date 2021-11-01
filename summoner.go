package gorrito

import (
	"net/url"

	"github.com/vaguilera/gorrito/models"
)

func (c *Client) performSummonerQuery(url string, params map[string]string) (*models.Summoner, error) {
	sum := models.Summoner{}
	body, err := c.requestAPI(url, params)
	if err != nil {
		return nil, err
	}

	err = c.unMarshall(body, &sum)
	return &sum, err
}

func (c *Client) performSummonerQueryRaw(url string, params map[string]string) (*string, error) {
	body, err := c.requestAPI(url, params)
	if err != nil {
		return nil, err
	}

	resBody := string(body)
	return &resBody, nil
}

func (c *Client) SummonerByAccountID(account string) (*models.Summoner, error) {
	return c.performSummonerQuery(UriSummonerByAccountID, map[string]string{"encryptedAccountId": account})
}

func (c *Client) SummonerByName(name string) (*models.Summoner, error) {
	return c.performSummonerQuery(UriSummonerByName, map[string]string{"summonerName": url.QueryEscape(name)})
}

func (c *Client) SummonerByPuuid(puuid string) (*models.Summoner, error) {
	return c.performSummonerQuery(UriSummonerByPuuid, map[string]string{"encryptedPUUID": puuid})
}

func (c *Client) SummonerBySummonerID(summonerId string) (*models.Summoner, error) {
	return c.performSummonerQuery(UriSummonerBySummonerID, map[string]string{"encryptedSummonerId": summonerId})
}

// Raw Versions
func (c *Client) SummonerByAccountIDRaw(account string) (*string, error) {
	return c.performSummonerQueryRaw(UriSummonerByAccountID, map[string]string{"encryptedAccountId": account})
}

func (c *Client) SummonerByNameRaw(name string) (*string, error) {
	return c.performSummonerQueryRaw(UriSummonerByName, map[string]string{"summonerName": url.QueryEscape(name)})
}

func (c *Client) SummonerByPuuidRaw(puuid string) (*string, error) {
	return c.performSummonerQueryRaw(UriSummonerByPuuid, map[string]string{"encryptedPUUID": puuid})
}

func (c *Client) SummonerBySummonerIDRaw(summonerId string) (*string, error) {
	return c.performSummonerQueryRaw(UriSummonerBySummonerID, map[string]string{"encryptedSummonerId": summonerId})
}
