package heroes

import (
    "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "godota/globals"
)

var (
    API_URL = "http://api.steampowered.com/IEconDOTA2_" + globals.API_REALM +
        "/GetHeroes/v0001?key=" + globals.API_KEY +
        "&language=en_us"
)

type Heroes struct {
    Heroes []struct{
        Name string `json:"name"`
        Id uint64 `json:"id"`
        LocalizedName string `json:"localized_name"`
    } `json:"heroes"`
    Count uint64 `json:"count"`
}

func List() chan Heroes {
    out := make(chan Heroes)

    go func() {
        resp, err := http.Get(API_URL)
        if err != nil {
            log.Fatalln(err)
        }

        body, _ := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        log.Println(string(body))
        target := struct{
            Result Heroes `json:"result"`
        }{}
        json.Unmarshal(body, &target)
        out <- target.Result
        close(out)
    }()

    return out
}
