package gorrito

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type ApiError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%d:%s", e.StatusCode, e.Status)
}

func (c *Client) parseURL(endpoint string, params map[string]string) string {

	if params != nil {
		for k, v := range params {
			endpoint = strings.Replace(endpoint, "{"+k+"}", v, -1)
		}
	}
	if strings.Contains(endpoint, "/v5/") {
		return "https://" + regionsV5[c.regionV5] + endpoint
	}
	return "https://" + regions[c.region] + endpoint

}

func (c *Client) unMarshall(body []byte, st interface{}) error {
	err := json.Unmarshal(body, &st)

	return err
}

func (c *Client) doRequest(endpoint string) ([]byte, error) {

	client := &http.Client{}
	client.Timeout = time.Second * 15

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "GorRito Lib by V.Aguilera")
	req.Header.Set("X-Riot-Token", c.token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, &ApiError{StatusCode: resp.StatusCode, Status: resp.Status, Body: string(body)}
	}

	return body, nil
}

func (c *Client) requestAPI(url string, params map[string]string, queryParams ...string) ([]byte, error) {
	endpoint := c.parseURL(url, params)

	if len(queryParams) > 0 {
		endpoint = endpoint + queryParams[0]
	}

	return c.doRequest(endpoint)
}
