package main

import (
    "log"
    "strconv"
    //"godota/teams"
    "godota/matchhistory"
)

var (
)

func main() {
    matchStream := matchhistory.MatchFeed(51945535)

    var ok = true
    var match matchhistory.Match
    for ok {
        match, ok = <-matchStream
        if match.LobbyType == 5 {
            log.Println("Found Cascade TMM Match: " +
                strconv.FormatUint(match.MatchId, 10))
        }
    }
}
