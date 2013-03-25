package matchhistory

import (
    "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "godota/globals"
)

var (
    API_URL = "http://api.steampowered.com/IDOTA2Match_" + globals.API_REALM +
        "/GetMatchHistory/v001?key=" + globals.API_KEY +
        "&account_id=51945535"
)

func Test() {
    log.Println("URL is " + API_URL)

    resp, err := http.Get(API_URL)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    target := struct {
        Result struct {
            Status              uint64      `json:"status"`
            NumResults          uint64      `json:"num_results"`
            TotalResults        uint64      `json:"total_results"`
            RemainingResults    uint64      `json:"results_remaining"`
            Matches []struct {
                MatchId         uint64      `json:"match_id"`
                MatchSeqNum     uint64      `json:"match_seq_num"`
                StartTime       uint64      `json:"start_time"`
                LobbyType       uint64      `json:"lobby_type"`
                Players []struct {
                    AccountId   uint64      `json:"account_id"`
                    PlayerSlot  uint64      `json:"player_slot"`
                    HeroId      uint64      `json:"hero_id"`
                } `json:"players"`
            } `json:"matches"`
        } `json:"result"`
    }{}

    json.Unmarshal(body, &target)

    log.Printf("%+v", target)
}
