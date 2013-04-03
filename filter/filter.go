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
