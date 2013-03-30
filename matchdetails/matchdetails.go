package matchdetails

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
        "/GetMatchDetails/v001?key=" + globals.API_KEY +
        "&account_id=51945535"
)

type MatchDetails struct {
    Result struct {
        Players []struct {
            AccountId      uint64   `json:"account_id"`
            PlayerSlot     uint64   `json:"player_slot"`
            HeroId         uint64   `json:"hero_id"`
            Item0          uint64   `json:"item_0"`
            Item1          uint64   `json:"item_1"`
            Item2          uint64   `json:"item_2"`
            Item3          uint64   `json:"item_3"`
            Item4          uint64   `json:"item_4"`
            Item5          uint64   `json:"item_5"`
            Kills          uint64   `json:"kills"`
            Deaths         uint64   `json:"deaths"`
            Assists        uint64   `json:"assists"`
            LeaverStatus   uint64   `json:"leaver_status"`
            Gold           uint64   `json:"gold"`
            LastHits       uint64   `json:"last_hits"`
            Denies         uint64   `json:"denies"`
            GoldPerMin     uint64   `json:"gold_per_min"`
            XpPerMin       uint64   `json:"xp_per_min"`
            GoldSpent      uint64   `json:"gold_spent"`
            HeroDamage     uint64   `json:"hero_damage"`
            TowerDamage    uint64   `json:"tower_damage"`
            HeroHealing    uint64   `json:"hero_healing"`
            Level          uint64   `json:"level"`

            Abilities []struct {
                Ability   uint64   `json:"ability"`
                Time      uint64   `json:"time"`
                Level     uint64   `json:"level"`
            } `json:"ability_upgrades"`
        } `json:"players"`

        Season                  uint64   `json:"season"`
        RadiantWin              uint64   `json:"radiant_win"`
        Duration                uint64   `json:"duration"`
        StartTime               uint64   `json:"start_time"`
        MatchId                 uint64   `json:"match_id"`
        MatchSeqNum             uint64   `json:"match_seq_num"`
        TowerStatusRadiant      uint64   `json:"tower_status_radiant"`
        TowerStatusDire         uint64   `json:"tower_status_dire"`
        BarracksStatusRadiant   uint64   `json:"barracks_status_radiant"`
        BarracksStatusDire      uint64   `json:"barracks_status_dire"`
        Cluster                 uint64   `json:"cluster"`
        FirstBloodTime          uint64   `json:"first_blood_time"`
        LobbyType               uint64   `json:"lobby_type"`
        HumanPlayers            uint64   `json:"human_players"`
        LeagueId                uint64   `json:"leagueid"`
        PositiveVotes           uint64   `json:"positive_votes"`
        NegativeVotes           uint64   `json:"negative_votes"`
        GameMode                uint64   `json:"game_mode"`
    } `json:"result"`
}

func ForMatch(matchId uint64) chan MatchDetails {
    matchStream := make(chan MatchDetails)

    go func() {
        url := API_URL + "&match_id=" + strconv.FormatUint(matchId, 10)
        resp, err := http.Get(url)
        if err != nil {
            log.Fatalln(err)
        }

        body, _ := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        target := MatchDetails{}
        json.Unmarshal(body, &target)
        matchStream <- target
    }()

    return matchStream
}
