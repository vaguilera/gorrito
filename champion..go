package gorrito

import "github.com/vaguilera/gorrito/models"

func (c *Client) ChampionRotations() (*models.Champions, error) {
	champ := models.Champions{}
	body, err := c.requestAPI(UriChampionRotations, nil)
	if err != nil {
		return nil, err
	}
	err = c.unMarshall(body, &champ)
	return &champ, err
}

func (c *Client) ChampionRotationsRaw() (*string, error) {
	body, err := c.requestAPI(UriChampionRotations, nil)
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}
