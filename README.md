# Gorrito
League of Legends (LOL) API V5 wrapper for GOlang

###  Example
```golang
const token = "MysecretToken"
const region = "EUW"

client := gorrito.NewClient(token, &gorrito.Config{Region: region})
queue, err := client.LeagueLeaguesByQueueTierDivision("RANKED_SOLO_5x5", "GOLD", "II")
fmt.Println(queue, err)
```

### Match filters with constants
```golang
options := client.NewMatchesByPuuidOptions().
    Queue(gorrito.QueueRankedSolo).
    Type(gorrito.MatchTypeRanked).
    Count(20)

matches, err := client.MatchesByPuuid("my-puuid", options)
fmt.Println(matches, err)
```

### Rate limit configuration
You can throttle outgoing requests per client instance:

```golang
client := gorrito.NewClient(token, &gorrito.Config{
    Region:             "EUW",
    RateLimitPerSecond: 20,
    RateLimitPerMinute: 100,
})
```

Both limits are evaluated at the same time. If one is reached, Gorrito waits before sending the next request.
Set either value to `0` to disable that specific limit.

### Debug mode
You can enable request/response debugging from the client config:

```golang
client := gorrito.NewClient(token, &gorrito.Config{
    Region: "EUW",
    Debug:  true,
})
```

When `Debug` is `true`, Gorrito prints:
- HTTP method and URL
- request headers
- response status
- response headers
- response body

There is also a runnable example in `cmd/debug/main.go` that calls:
`client.AccountByRiotID("HeladoDeLentejas", "TAG")`.

V5 endpoints supported  
Included package 'models' contains all the structs for the returned data

## Compatibility notes (reviewed on 2026-03-24)

This repository focuses only on League of Legends endpoints.

### Current compatibility status
- ✅ `match-v5` is supported.
- ✅ `account-v1` is supported (`by-puuid`, `by-riot-id`, and `active-shards`).
- ✅ `league-v4`, `champion-mastery-v4`, `champion-v3`, `lol-status-v4`, and current `summoner-v4` lookups are supported (`by-puuid` and `by-summoner-id`).
- ✅ `spectator-v5` is supported.
- ✅ Regional routing includes SEA shards (`PH`, `SG`, `TH`, `TW`, `VN`) plus `sea.api.riotgames.com`.
- ✅ Models have been split per endpoint domain and several JSON tag mismatches were corrected to improve unmarshalling accuracy.
