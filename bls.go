package bls

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	RegistrationKey, StartYear, EndYear  string
	Catalog, Calculations, AnnualAverage bool
	Series                               []string
}

func (r *Request) SingleSeries(series int) string {
	url := "http://api.bls.gov/publicAPI/v2/timeseries/data/"
	req, err := http.NewRequest("POST", url+r.Series[series], nil)
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

func (r *Request) MultipleSeries() {
	// "LAUCN040010000000005", "LAUCN040010000000006"
	s := "\"" + strings.Join(r.Series, "\",\"") + "\""
	payload := `{"seriesid":[` + s + `],
					"startyear":"2010", 
					"endyear":"2012",
					"catalog":false,
					"calculations":true,
					"annualaverage":true,
					"registrationKey":"acaf8533b1024e3c827414ff9e01e8f7"}`
	url := "http://api.bls.gov/publicAPI/v2/timeseries/data/"

	jsonStr := []byte(payload)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func check(err error) {
	if err != nil {
		panic(err)
	}

}
