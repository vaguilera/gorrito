# Gorrito
League of Legends (LOL) API V4 wrapper for GOlang

###  Example
```golang
const token = "MysecretToken"
const region = "EUW"

gorrito := Gorrito{region: region, token: token}

var positions []LeaguePosition
status := gorrito.leaguePositions(&positions, "RANKED_SOLO_5x5", "GOLD", "I", "TOP", 0)
fmt.Println(status, positions)
```
Look inside every file to see the methods and structures you need to make the calls.
All the endpoints are supported except the tournament ones.
