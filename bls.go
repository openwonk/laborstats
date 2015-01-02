package bls

import (
        // "bytes"
        "fmt"
        "io/ioutil"
        "net/http"
)

// example series "LAUCN040010000000005"
func SingleSeries(series string) string {
        url := "http://api.bls.gov/publicAPI/v2/timeseries/data/"
        // var jsonStr = []byte(`{"seriesid": ['CUUR0000SA0','SUUR0000SA0'],"startyear":"2011", "endyear":"2014"}`)
        req, err := http.NewRequest("POST", url+series, nil)
        req.Header.Set("Content-Type", "application/json")

        client := &http.Client{}
        resp, err := client.Do(req)
        check(err)
        defer resp.Body.Close()

        fmt.Println("response Status:", resp.Status)
        fmt.Println("response Headers:", resp.Header)
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println("response Body:", string(body))
        
        return string(body)
}

func check (err error) {
        if err != nil {
                panic(err)
        }               
        
}
