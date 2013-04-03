package filter

import (
    "godota/matchhistory"
)

func MatchHistoryByLobby(in chan matchhistory.Match,
        lobby uint64) chan matchhistory.Match {

    out := make(chan matchhistory.Match)

    go func() {
        var match matchhistory.Match

        for ok := true; ok; {
            match, ok = <-in
            if match.LobbyType == lobby {
                out <- match
            }
        }

        close(out)
    }()

    return out
}

func ByPlayersInvolved(in chan matchhistory.Match,
        accountIds[]uint64) chan matchhistory.Match {

    out := make(chan matchhistory.Match)

    go func() {
        var match matchhistory.Match

        for ok := true; ok; {
            match, ok = <-in

            found := false

            for _, id := range accountIds {
                found = false

                for _, player := range match.Players {
                    if player.AccountId == id {
                        found = true
                        break
                    }
                }

                if found == false {
                    break
                }
            }

            if found {
                out <- match
            }
        }

        close(out)
    }()

    return out
}
