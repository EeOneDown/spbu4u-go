package spbu_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type EducatorEvents struct {
	EducatorMasterId                 int64
	EducatorDisplayText              string
	EducatorLongDisplayText          string
	PreviousWeekMonday               JsonDate
	NextWeekMonday                   JsonDate
	IsPreviousWeekReferenceAvailable bool
	IsNextWeekReferenceAvailable     bool
	IsCurrentWeekReferenceAvailable  bool
	WeekDisplayText                  string
	WeekMonday                       JsonDate
	EducatorEventsDays               []EventsDay
}

func parseEducatorEvents(resp *http.Response, educatorEvents *EducatorEvents) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &educatorEvents); err != nil {
		return err
	}
	return nil
}

func GetEducatorScheduleFor(id int64, from time.Time, to time.Time) (*EducatorEvents, error) {
	var educatorEvents EducatorEvents
	fromYear, fromMonth, fromDay := from.Date()
	toYear, toMonth, toDay := to.Date()
	url := fmt.Sprintf(EducatorCustomUrl, id, fromYear, fromMonth, fromDay, toYear, toMonth, toDay)
	resp, err := netClient.Get(url)
	if err != nil {
		return &educatorEvents, err
	}
	err = parseEducatorEvents(resp, &educatorEvents)
	return &educatorEvents, err
}
