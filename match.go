package gorrito

import (
	"bytes"
	"strconv"

	"github.com/vaguilera/gorrito/models"
)

func (c *Client) MatchByID(matchID string) (*models.Match, error) {
	match := models.Match{}
	body, err := c.requestAPI(UriMatchByID, map[string]string{"matchId": matchID})
	if err != nil {
		return nil, err
	}
	err = c.unMarshall(body, &match)
	return &match, err
}

func (c *Client) MatchByIDRaw(matchID string) (*string, error) {
	body, err := c.requestAPI(UriMatchByID, map[string]string{"matchId": matchID})
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}

func getMatchesByPuuidParameters(options *matchesByPuuidOptions) string {
	var parameters bytes.Buffer

	parameters.WriteString("?")

	if options.startTime != nil {
		parameters.WriteString("&startTime=" + strconv.FormatInt(*options.startTime, 10))
	}
	if options.endTime != nil {
		parameters.WriteString("&endTime=" + strconv.FormatInt(*options.endTime, 10))
	}
	if options.queue != nil {
		parameters.WriteString("&queue=" + strconv.Itoa(*options.queue))
	}
	if options.ttype != nil {
		parameters.WriteString("&type=" + *options.ttype)
	}
	if options.start != nil {
		parameters.WriteString("&start=" + strconv.Itoa(*options.start))
	}
	if options.count != nil {
		parameters.WriteString("&count=" + strconv.Itoa(*options.count))
	}

	return parameters.String()
}

func (c *Client) MatchesByPuuid(puuid string, options *matchesByPuuidOptions) (models.MatchesList, error) {
	matches := make(models.MatchesList, 0)
	parameters := getMatchesByPuuidParameters(options)
	body, err := c.requestAPI(UriMatchByPuuid, map[string]string{"puuid": puuid}, parameters)
	if err != nil {
		return nil, err
	}

	err = c.unMarshall(body, &matches)
	return matches, err
}

func (c *Client) MatchesByPuuidRaw(puuid string, options *matchesByPuuidOptions) (*string, error) {
	parameters := getMatchesByPuuidParameters(options)
	body, err := c.requestAPI(UriMatchByPuuid, map[string]string{"puuid": puuid}, parameters)
	if err != nil {
		return nil, err
	}
	resBody := string(body)
	return &resBody, nil
}

func (c *Client) MatchTimeline(matchID string) (*models.MatchTimeLine, error) {
	timeline := models.MatchTimeLine{}
	body, err := c.requestAPI(UriMatchTimeline, map[string]string{"matchId": matchID})
	if err != nil {
		return nil, err

	}

	err = c.unMarshall(body, &timeline)
	return &timeline, err
}

func (c *Client) MatchTimelineRaw(matchID string) (*string, error) {
	body, err := c.requestAPI(UriMatchTimeline, map[string]string{"matchId": matchID})
	if err != nil {
		return nil, err

	}
	resBody := string(body)
	return &resBody, nil
}
