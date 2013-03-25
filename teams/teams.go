package teams

import (
    "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "godota/globals"
)

var (
    API_URL = "http://api.steampowered.com/IDOTA2Match_" + globals.API_REALM +
        "/GetTeamInfoByTeamID/v001?key=" + globals.API_KEY +
        "&teams_requested=1" +
        "&start_at_team_id=4448"
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
            Status int `json:"status"`
            Teams []struct {
                TeamId        uint64   `json:"team_id"`
                Name          string   `json:"name"`
                Tag           string   `json:"tag"`
                TimeCreated   uint64   `json:"time_created"`
                Rating        uint64   `json:"rating"`
                Logo          uint64   `json:"logo"`
                LogoSponsor   uint64   `json:"logo_sponsor"`
                CountryCode   string   `json:"country_code"`
                Url           string   `json:"url"`
                GamesPlayed   uint64   `json:"games_played_with_current_roster"`
                Player0Id     uint64   `json:"player_0_account_id"`
                Player1Id     uint64   `json:"player_1_account_id"`
                Player2Id     uint64   `json:"player_2_account_id"`
                Player3Id     uint64   `json:"player_3_account_id"`
                Player4Id     uint64   `json:"player_4_account_id"`
                AdminId       uint64   `json:"admin_account_id"`
            } `json:"teams"`
        } `json:"result"`
    }{}

    json.Unmarshal(body, &target)

    log.Printf("%+v", target)
}
