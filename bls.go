package bls

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	url = "http://api.bls.gov/publicAPI/v2/timeseries/data/"
)

type Series struct {
	RegistrationKey, StartYear, EndYear  string
	Catalog, Calculations, AnnualAverage bool
	Series                               []string
}

func (r *Series) Request() string {
	// "LAUCN040010000000005", "LAUCN040010000000006"
	payload := `{
		"seriesid":[` + "\"" + strings.Join(r.Series, "\",\"") + "\"" + `],
		"startyear":"` + r.StartYear + `","endyear":"` + r.EndYear + `",
		"catalog":` + strconv.FormatBool(r.Catalog) + `,
		"calculations":` + strconv.FormatBool(r.Calculations) + `,
		"annualaverage":` + strconv.FormatBool(r.AnnualAverage) + `,
		"registrationKey":"` + r.RegistrationKey + `"}`

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
