package spbu_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//
// T Y P E S
//
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

type EventsDay struct {
	Day            ZLessTime
	DayString      string
	DayStudyEvents []Event
}

type Event struct {
	StudyEventsTimeTableKindCode     int
	Start                            ZLessTime
	End                              ZLessTime
	Subject                          string
	TimeIntervalString               string
	DateWithTimeIntervalString       string
	DisplayDateAndTimeIntervalString string
	LocationsDisplayText             string
	EducatorsDisplayText             string
	HasEducators                     bool
	IsCancelled                      bool
	ContingentUnitName               string
	DivisionAndCourse                string
	IsAssigned                       bool
	TimeWasChanged                   bool
	LocationsWereChanged             bool
	EducatorsWereReassigned          bool
	ElectiveDisciplinesCount         int
	IsElective                       bool
	HasTheSameTimeAsPreviousItem     bool
	IsStudy                          bool
	AllDay                           bool
	WithinTheSameDay                 bool
	EventLocations                   []EventLocation
	EducatorIds                      []EducatorId
}

type EventLocation struct {
	IsEmpty                  bool
	DisplayName              string
	HasGeographicCoordinates bool
	Latitude                 float32
	Longitude                float32
	LatitudeValue            string
	LongitudeValue           string
	EducatorsDisplayText     string
	HasEducators             bool
	EducatorIds              []EducatorId
}

type EducatorId struct {
	Id   int64  `json:"Item1"`
	Name string `json:"Item2"`
}

//
// H E L P E R S
//
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

//
// M E T H O D S
//
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
