package gorrito

import (
	"github.com/vaguilera/gorrito/models"
)

func (c *Client) performAccountQuery(url string, params map[string]string) (*models.Account, error) {
	account := models.Account{}
	body, err := c.requestAPI(url, params)
	if err != nil {
		return nil, err
	}

	err = c.unMarshall(body, &account)
	return &account, err
}

func (c *Client) performActiveShardQuery(url string, params map[string]string) (*models.ActiveShard, error) {
	activeShard := models.ActiveShard{}
	body, err := c.requestAPI(url, params)
	if err != nil {
		return nil, err
	}

	err = c.unMarshall(body, &activeShard)
	return &activeShard, err
}

func (c *Client) performAccountQueryRaw(url string, params map[string]string) (*string, error) {
	body, err := c.requestAPI(url, params)
	if err != nil {
		return nil, err
	}

	resBody := string(body)
	return &resBody, nil
}

func (c *Client) AccountByPuuid(puuid string) (*models.Account, error) {
	return c.performAccountQuery(UriAccountByPuuid, map[string]string{"puuid": puuid})
}

func (c *Client) AccountByRiotID(gameName string, tagLine string) (*models.Account, error) {
	return c.performAccountQuery(UriAccountByRiotID, map[string]string{"gameName": gameName, "tagLine": tagLine})
}

func (c *Client) AccountActiveShardByPuuid(game string, puuid string) (*models.ActiveShard, error) {
	return c.performActiveShardQuery(UriAccountActiveShardByPuuid, map[string]string{"game": game, "puuid": puuid})
}

func (c *Client) AccountByPuuidRaw(puuid string) (*string, error) {
	return c.performAccountQueryRaw(UriAccountByPuuid, map[string]string{"puuid": puuid})
}

func (c *Client) AccountByRiotIDRaw(gameName string, tagLine string) (*string, error) {
	return c.performAccountQueryRaw(UriAccountByRiotID, map[string]string{"gameName": gameName, "tagLine": tagLine})
}

func (c *Client) AccountActiveShardByPuuidRaw(game string, puuid string) (*string, error) {
	return c.performAccountQueryRaw(UriAccountActiveShardByPuuid, map[string]string{"game": game, "puuid": puuid})
}
