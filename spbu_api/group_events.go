package spbu_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type GroupEvents struct {
	StudentGroupId                   int64
	StudentGroupDisplayName          string
	TimeTableDisplayName             string
	PreviousWeekMonday               JsonDate
	NextWeekMonday                   JsonDate
	IsPreviousWeekReferenceAvailable bool
	IsNextWeekReferenceAvailable     bool
	IsCurrentWeekReferenceAvailable  bool
	WeekDisplayText                  string
	WeekMonday                       JsonDate
	Days                             []EventsDay
}

func parseGroupEvents(resp *http.Response, groupEvents *GroupEvents) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &groupEvents); err != nil {
		return err
	}
	return nil
}

func GetGroupScheduleFor(id int64, from time.Time, to time.Time) (*GroupEvents, error) {
	var groupEvents GroupEvents
	fromYear, fromMonth, fromDay := from.Date()
	toYear, toMonth, toDay := to.Date()
	url := fmt.Sprintf(GroupCustomUrl, id, fromYear, fromMonth, fromDay, toYear, toMonth, toDay)
	resp, err := http.Get(url)
	if err != nil {
		return &groupEvents, err
	}
	err = parseGroupEvents(resp, &groupEvents)
	return &groupEvents, err
}
