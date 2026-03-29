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
	if isRegionalEndpoint(endpoint) {
		return "https://" + regionsV5[c.regionV5] + endpoint
	}
	return "https://" + regions[c.region] + endpoint

}

func isRegionalEndpoint(endpoint string) bool {
	return strings.HasPrefix(endpoint, "/lol/match/v5/") ||
		strings.HasPrefix(endpoint, "/riot/account/v1/")
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

	if c.debug {
		fmt.Printf("[gorrito][request] %s %s\n", req.Method, req.URL.String())
		fmt.Printf("[gorrito][request][headers] %v\n", req.Header)
	}

	resp, err := client.Do(req)
	if err != nil {
		if c.debug {
			fmt.Printf("[gorrito][error] %v\n", err)
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Printf("[gorrito][response] status=%s\n", resp.Status)
		fmt.Printf("[gorrito][response][headers] %v\n", resp.Header)
		fmt.Printf("[gorrito][response][body] %s\n", string(body))
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
