package main

import (
	"fmt"
	"spbu4u-go/constants"
	"spbu4u-go/spbu_api"
	"spbu4u-go/spbu_api/types"
	"time"
)

type Schedule interface {
	Parse() ([]string, error)
}

func (scheduleStorage *ScheduleStorage) GetScheduleFor(from time.Time, to time.Time) (Schedule, error) {
	if scheduleStorage.ID == 0 {
		var notRegistered *NotRegistered
		var schedule Schedule = notRegistered
		return schedule, nil
	} else if scheduleStorage.Type == constants.GROUP {
		groupEvents, err := spbu_api.GetGroupScheduleFor(scheduleStorage.TimeTableId, from, to)
		if err != nil {
			var scheduleNotAllowed *ScheduleNotAllowed
			var schedule Schedule = scheduleNotAllowed
			return schedule, err
		}
		var schedule Schedule = (*GroupEvents)(groupEvents)
		return schedule, nil
	} else {
		educatorEvents, err := spbu_api.GetEducatorScheduleFor(scheduleStorage.TimeTableId, from, to)
		if err != nil {
			var scheduleNotAllowed *ScheduleNotAllowed
			var schedule Schedule = scheduleNotAllowed
			return schedule, err
		}
		var schedule Schedule = (*EducatorEvents)(educatorEvents)
		return schedule, nil
	}
}

type ScheduleNotAllowed int

func (scheduleNotAllowed *ScheduleNotAllowed) Parse() ([]string, error) {
	return []string{"I can't get schedule from the timetable.spbu.ru. Perhaps the TimeTable server is down."}, nil
}

type NotRegistered int

func (notRegistered *NotRegistered) Parse() ([]string, error) {
	return []string{"Schedule is not allowed for you. Please register via the timetable.spbu.ru schedule link."}, nil
}

type GroupEvents types.GroupEvents

func (groupEvents *GroupEvents) Parse() ([]string, error) {
	var parsed []string
	for _, day := range groupEvents.Days {
		if len(day.DayStudyEvents) == 0 {
			continue
		}
		dayParsed := ""
		for _, event := range day.DayStudyEvents {
			dayParsed += fmt.Sprintf("%s\n%s\n%s (%s)\n\n",
				event.TimeIntervalString,
				event.Subject,
				event.LocationsDisplayText,
				event.EducatorsDisplayText,
			)
		}
		if dayParsed == "" {
			continue
		}
		dayParsed = fmt.Sprintf("%s\n\n%s", day.DayString, dayParsed)
		parsed = append(parsed, day.DayString)
	}
	if len(parsed) == 0 {
		return []string{"Nothing to display."}, nil
	}
	return parsed, nil
}

type EducatorEvents types.EducatorEvents

func (educatorEvents *EducatorEvents) Parse() ([]string, error) {
	return []string{"Working on it."}, nil
}
