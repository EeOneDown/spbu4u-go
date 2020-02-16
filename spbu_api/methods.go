package spbu_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"spbu4u-go/spbu_api/types"
	"time"
)

// Parsers
func parseGroupEvents(resp *http.Response, groupEvents *types.GroupEvents) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &groupEvents); err != nil {
		return err
	}
	return nil
}

func parseEducatorEvents(resp *http.Response, educatorEvents *types.EducatorEvents) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &educatorEvents); err != nil {
		return err
	}
	return nil
}

// methods
func GetGroupScheduleFor(id int64, from time.Time, to time.Time) (*types.GroupEvents, error) {
	var groupEvents types.GroupEvents
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

func GetEducatorScheduleFor(id int64, from time.Time, to time.Time) (*types.EducatorEvents, error) {
	var educatorEvents types.EducatorEvents
	fromYear, fromMonth, fromDay := from.Date()
	toYear, toMonth, toDay := to.Date()
	url := fmt.Sprintf(EducatorCustomUrl, id, fromYear, fromMonth, fromDay, toYear, toMonth, toDay)
	resp, err := http.Get(url)
	if err != nil {
		return &educatorEvents, err
	}
	err = parseEducatorEvents(resp, &educatorEvents)
	return &educatorEvents, err
}
