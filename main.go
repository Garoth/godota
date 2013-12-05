package main

import (
    "log"
    "strconv"
    "time"
    "godota/matchhistory"
    "godota/matchdetails"
    "godota/heroes"
    "godota/filter"
)

var (
)

func main() {
    log.SetFlags(0)

    FindPlayerMatches()
    //FindCascadeMatches()
    //FindGamesTogether()
    //ListHeroes()
}

func ListHeroes() {
    in := heroes.List()
    list := <-in

    for _, hero := range list.Heroes {
        log.Printf("%v %v", hero.Id, hero.LocalizedName)
    }
}

func FindPlayerMatches() {
    //var accountId uint64 = 51945535 // Arkanian
    //var accountId uint64 = 51971876 // Kevlarman
    //var accountId uint64 = 75685110 // Nik
    var accountId uint64 = 53071885 // Spen
    //var accountId uint64 = 105771979 // Skeleton Burglar
    //var accountId uint64 = 114426207 // Polychromatic Hyphen
    //var accountId uint64 = 96033201 // Regie
    //var accountId uint64 = 101495620 // Loda
    //var accountId uint64 = 38359315 // Rag
    //var accountId uint64 = 38345312 // Edodes

    var totalKills, totalDeaths, totalAssists uint64 = 0, 0, 0
    /* This is (your team's total gold) / (enemy team's total gold) */
    var totalRomanScore, radiantGoldEarned, direGoldEarned float64 = 0, 0, 0
    maxMatches := 100
    var matchesFound uint64 = 0
    matchStream := matchhistory.ForAccountId(accountId)

    var match matchhistory.Match
    var iAmRadiant bool = true
    for ok, count := true, 0; ok && count < maxMatches; count++ {
        matchesFound++
        radiantGoldEarned, direGoldEarned = 0, 0
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
                if player.PlayerSlot < 64 {
                    iAmRadiant = true
                } else {
                    iAmRadiant = false
                }

                log.Printf("Player %v had %v:%v:%v and is radiant = %v", accountId,
                    player.Kills, player.Deaths, player.Assists, iAmRadiant)
                totalKills += player.Kills
                totalDeaths += player.Deaths
                totalAssists += player.Assists
            }

            if player.PlayerSlot < 64 {
                radiantGoldEarned += float64(player.GoldPerMin /
                    60 * details.Result.Duration)
            } else {
                direGoldEarned += float64(player.GoldPerMin /
                    60 * details.Result.Duration)
            }
        }

        if direGoldEarned > 0 && radiantGoldEarned > 0 {
            if iAmRadiant {
                totalRomanScore += radiantGoldEarned / direGoldEarned
            } else {
                totalRomanScore += direGoldEarned / radiantGoldEarned
            }
        } else {
            log.Println("ONE TEAM EARNED NO GOLD!!")
        }

        log.Printf("Radiant earned %v. Dire earned %v. " +
            "New Roman score total: %.3f", radiantGoldEarned, direGoldEarned,
            totalRomanScore);
        log.Println("----------------------------------------" +
            "-----------------------------------")
    }

    log.Printf("Looked at the last %v matches for %v", matchesFound, accountId);

    if (matchesFound == 0) {
        return
    }

    log.Printf("KDA was %v:%v:%v, which is %.2f",
        totalKills, totalDeaths, totalAssists,
        float64(totalKills + totalAssists) / float64(totalDeaths))
    log.Printf("Roman score is %.3f", (totalRomanScore / float64(matchesFound)))
}

func FindCascadeMatches() {
    in := filter.MatchHistoryByLobby(matchhistory.ForAccountId(51945535), 5)

    var match matchhistory.Match
    for ok := true; ok; {
        match, ok = <-in
        log.Println("Found Cascade TMM Match: " +
            strconv.FormatUint(match.MatchId, 10))
    }
}

func FindGamesTogether() {
    accountIds := []uint64{96033201, 51971876}

    in := filter.ByPlayersInvolved(matchhistory.ForAccountId(51945535),
        accountIds)

    log.Printf("Looking for matches with players %+v", accountIds)
    var match matchhistory.Match
    for ok := true; ok; {
        match, ok = <-in
        log.Printf("Found match: %d", match.MatchId)
    }
}
