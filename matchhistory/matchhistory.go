package matchhistory

import (
    "log"
    "strconv"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "godota/globals"
)

var (
    API_URL = "http://api.steampowered.com/IDOTA2Match_" + globals.API_REALM +
        "/GetMatchHistory/v001?key=" + globals.API_KEY
)

type Match struct {
    MatchId         uint64      `json:"match_id"`
    MatchSeqNum     uint64      `json:"match_seq_num"`
    StartTime       uint64      `json:"start_time"`
    LobbyType       uint64      `json:"lobby_type"`
    Players []struct {
        AccountId   uint64      `json:"account_id"`
        PlayerSlot  uint64      `json:"player_slot"`
        HeroId      uint64      `json:"hero_id"`
    } `json:"players"`
}

type MatchHistory struct {
    Result struct {
        Status              uint64      `json:"status"`
        NumResults          uint64      `json:"num_results"`
        TotalResults        uint64      `json:"total_results"`
        RemainingResults    uint64      `json:"results_remaining"`
        Matches             []Match     `json:"matches"`
    } `json:"result"`
}

func ForAccountId(AccountId uint64) chan Match {
    matchStream := make(chan Match)

    go func() {
        var startMatchId uint64 = 99999999999999

        for {
            url := API_URL + "&account_id=" +
                strconv.FormatUint(AccountId, 10) +
                "&start_at_match_id=" +
                strconv.FormatUint(startMatchId - 1, 10)
            resp, err := http.Get(url)
            if err != nil {
                log.Fatalln(err)
            }

            body, _ := ioutil.ReadAll(resp.Body)
            resp.Body.Close()
            target := MatchHistory{}
            json.Unmarshal(body, &target)

            if target.Result.NumResults <= 0 {
                close(matchStream)
                break
            } else {
                for _, match := range target.Result.Matches {
                    startMatchId = match.MatchId
                    matchStream <- match
                }
            }
        }
    }()

    return matchStream
}
