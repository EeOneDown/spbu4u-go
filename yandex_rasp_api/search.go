package yandex_rasp_api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type SearchSchedulePointPointParams struct {
	From           Station
	To             Station
	Format         string
	Lang           string
	Date           time.Time
	TransportTypes string
	System         string
	ShowSystems    string
	Offset         int
	Limit          int
	AddDaysMask    bool
	ResultTimezone string
	Transfers      bool
}

func SearchSchedulePointPoint(params *SearchSchedulePointPointParams) (*SearchResponse, error) {
	var searchResponse SearchResponse

	request, err := http.NewRequest("GET", SchedulePointPoint, nil)
	if err != nil {
		return &searchResponse, err
	}
	request.Header.Add("Authorization", os.Getenv("YANDEX_API_KEY"))

	q := request.URL.Query()
	q.Add("from", params.From.Code)
	q.Add("to", params.To.Code)
	q.Add("date", params.Date.Format("2006-01-02"))
	request.URL.RawQuery = q.Encode()

	resp, err := netClient.Do(request)
	if err != nil {
		return &searchResponse, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &searchResponse, err
	}
	if err := json.Unmarshal(body, &searchResponse); err != nil {
		return &searchResponse, err
	}
	return &searchResponse, err
}
