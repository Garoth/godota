package main

import (
    "log"
    "strconv"
    "time"
    "godota/matchhistory"
    "godota/matchdetails"
)

var (
)

func main() {
    FindPlayerMatches()
}

func FindPlayerMatches() {
    var accountId uint64 = 51945535 // Arkanian
    //var accountId uint64 = 51971876 // Kevlarman
    //var accountId uint64 = 75685110 // Nik
    //var accountId uint64 = 53071885 // Spen
    //var accountId uint64 = 105771979 // Skeleton Burglar
    //var accountId uint64 = 114426207 // Polychromatic Hyphen

    var totalKills, totalDeaths uint64 = 0, 0
    maxMatches := 100
    matchStream := matchhistory.MatchFeed(accountId)

    var match matchhistory.Match
    for ok, count := true, 0; ok && count < maxMatches; count++ {
        match, ok = <-matchStream
        if !ok {
            break
        }

        t := time.Unix(int64(match.StartTime), 0)
        log.Printf("Match %v on %v", match.MatchId, t.String())
        details := <-matchdetails.ForMatch(match.MatchId)
        details = details
        for _, player := range details.Result.Players {
            if player.AccountId == accountId {
                log.Printf("Player %v had %v:%v", accountId,
                    player.Kills, player.Deaths)
                totalKills += player.Kills
                totalDeaths += player.Deaths
            }
        }
    }

    log.Printf("----------------------")
    log.Printf("KDR was %v:%v, which is %.2f",
        totalKills, totalDeaths, float64(totalKills) / float64(totalDeaths))
}

func FindCascadeMatches() {
    matchStream := matchhistory.MatchFeed(51945535)

    var match matchhistory.Match
    for ok:= true; ok; {
        match, ok = <-matchStream
        if match.LobbyType == 5 {
            log.Println("Found Cascade TMM Match: " +
                strconv.FormatUint(match.MatchId, 10))
        }
    }
}
