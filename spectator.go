package gorrito

import "github.com/vaguilera/gorrito/models"

func (c *Client) SpectatorActiveGames(summonerId string) (*models.CurrentGameInfo, error) {
	currentGameInfo := models.CurrentGameInfo{}
	body, err := c.requestAPI(UriSpectatorActiveGames, map[string]string{"encryptedSummonerId": summonerId})
	if err != nil {
		return nil, err
	}

	err = c.unMarshall(body, &currentGameInfo)
	return &currentGameInfo, err
}

func (c *Client) SpectatorFeaturedGames() (*models.FeaturedGames, error) {
	featuredGames := models.FeaturedGames{}
	body, err := c.requestAPI(UriSpectatorFeaturedGames, nil)
	if err != nil {
		return nil, err
	}
	err = c.unMarshall(body, &featuredGames)
	return &featuredGames, err
}

func (c *Client) SpectatorActiveGamesRaw(summonerId string) (*string, error) {
	body, err := c.requestAPI(UriSpectatorActiveGames, map[string]string{"encryptedSummonerId": summonerId})
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil

}

func (c *Client) SpectatorFeaturedGamesRaw() (*string, error) {
	body, err := c.requestAPI(UriSpectatorFeaturedGames, nil)
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}
