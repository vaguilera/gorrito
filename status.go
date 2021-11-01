package gorrito

import "github.com/vaguilera/gorrito/models"

func (c *Client) LolStatus() (*models.ShardStatus, error) {
	shard := models.ShardStatus{}
	body, err := c.requestAPI(UriLolStatus, nil)
	if err != nil {
		return nil, err
	}
	err = c.unMarshall(body, &shard)
	return &shard, err
}

func (c *Client) LolStatusRaw() (*string, error) {
	body, err := c.requestAPI(UriLolStatus, nil)
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}
