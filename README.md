# Gorrito
League of Legends (LOL) API V5 wrapper for GOlang

###  Example
```golang
const token = "MysecretToken"
const region = "EUW"

gorrito := NewClient{token: token}
client := gorrito.NewClient(token, &gorrito.Config{Region: region})

queue, err := client.LeagueLeaguesByQueueTierDivision("RANKED_SOLO_5x5", "GOLD", "II")
fmt.Println(queue, err)
```

V5 endpoints supported  
Included package 'models' contains all the structs for the returned data



