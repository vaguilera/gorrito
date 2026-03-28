package gorrito

import "testing"

func TestIsRegionalEndpoint(t *testing.T) {
	tests := []struct {
		endpoint string
		want     bool
	}{
		{endpoint: UriMatchByID, want: true},
		{endpoint: UriAccountByPuuid, want: true},
		{endpoint: UriSpectatorFeaturedGames, want: false},
		{endpoint: UriSummonerByPuuid, want: false},
	}

	for _, tc := range tests {
		got := isRegionalEndpoint(tc.endpoint)
		if got != tc.want {
			t.Fatalf("isRegionalEndpoint(%q) = %t, want %t", tc.endpoint, got, tc.want)
		}
	}
}

func TestParseURLRouting(t *testing.T) {
	client := NewClient("token", &Config{Region: "EUW"})

	matchURL := client.parseURL(UriMatchByID, map[string]string{"matchId": "EUW1_1"})
	if matchURL != "https://europe.api.riotgames.com/lol/match/v5/matches/EUW1_1" {
		t.Fatalf("unexpected match URL: %s", matchURL)
	}

	accountURL := client.parseURL(UriAccountByPuuid, map[string]string{"puuid": "abc"})
	if accountURL != "https://europe.api.riotgames.com/riot/account/v1/accounts/by-puuid/abc" {
		t.Fatalf("unexpected account URL: %s", accountURL)
	}

	spectatorURL := client.parseURL(UriSpectatorFeaturedGames, nil)
	if spectatorURL != "https://euw1.api.riotgames.com/lol/spectator/v5/featured-games" {
		t.Fatalf("unexpected spectator URL: %s", spectatorURL)
	}
}
