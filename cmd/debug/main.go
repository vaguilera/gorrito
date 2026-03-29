package main

import (
	"fmt"
	"os"

	"github.com/vaguilera/gorrito"
)

func main() {
	token := os.Getenv("RIOT_API_TOKEN")
	if token == "" {
		fmt.Println("set RIOT_API_TOKEN to run this debug command")
		return
	}

	region := os.Getenv("RIOT_REGION")
	if region == "" {
		region = "EUW"
	}

	client := gorrito.NewClient(token, &gorrito.Config{
		Region: region,
		Debug:  true,
	})
	if client == nil {
		fmt.Printf("invalid region: %s\n", region)
		return
	}

	account, err := client.AccountByRiotID("HeladoDeLentejas", "#TAG")
	if err != nil {
		fmt.Printf("account lookup failed: %v\n", err)
		return
	}

	fmt.Printf("account: %+v\n", *account)
}
