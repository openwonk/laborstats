package bls

import (
	"bytes"
	// "fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Request struct {
	RegistrationKey, StartYear, EndYear  string
	Catalog, Calculations, AnnualAverage bool
	Series                               []string
}

func (r *Request) SingleSeries() string {
	url := "http://api.bls.gov/publicAPI/v2/timeseries/data/"
	req, err := http.NewRequest("POST", url+r.Series[0], nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func (r *Request) MultipleSeries() string {
	// "LAUCN040010000000005", "LAUCN040010000000006"
	payload := `{
		"seriesid":[` + "\"" + strings.Join(r.Series, "\",\"") + "\"" + `],
		"startyear":"` + r.StartYear + `","endyear":"` + r.EndYear + `",
		"catalog":` + strconv.FormatBool(r.Catalog) + `,
		"calculations":` + strconv.FormatBool(r.Calculations) + `,
		"annualaverage":` + strconv.FormatBool(r.AnnualAverage) + `,
		"registrationKey":"` + r.RegistrationKey + `"}`

	url := "http://api.bls.gov/publicAPI/v2/timeseries/data/"

	jsonStr := []byte(payload)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func check(err error) {
	if err != nil {
		panic(err)
	}

}
