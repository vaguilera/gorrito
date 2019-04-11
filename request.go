package gorrito

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func parseURL(endpoint string, region string, params map[string]string) string {

	if params != nil {
		for k, v := range params {
			endpoint = strings.Replace(endpoint, "{"+k+"}", v, -1)
		}
	}
	return "https://" + regions[region] + endpoint
}

func (gorrito *Gorrito) requestAPI(url string, data interface{}, params map[string]string, queryParams ...string) int {

	client := &http.Client{}
	client.Timeout = time.Second * 15
	endpoint := parseURL(url, gorrito.region, params)

	if len(queryParams) > 0 {
		endpoint = endpoint + queryParams[0]
	}

	fmt.Println(endpoint)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Fatalf("http.NewRequest() failed with '%s'\n", err)
		return -1
	}
	req.Header.Set("User-Agent", "GorRito Lib by V.Aguilera")
	req.Header.Set("X-Riot-Token", gorrito.token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.Do() failed with '%s'\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
		return -1
	}

	fmt.Println(string(body))

	if resp.StatusCode == 200 {
		if err := json.Unmarshal(body, &data); err != nil {
			log.Fatal("Error unmarshalling response:", err)
			return -1
		}
	}

	return resp.StatusCode
}
