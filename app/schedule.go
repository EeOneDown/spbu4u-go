package app

import (
	"fmt"
	"spbu4u-go/spbu_api"
	"strings"
)

type Schedule interface {
	Parse() ([]string, error)
}

type ScheduleNotAllowed int

func (scheduleNotAllowed *ScheduleNotAllowed) Parse() ([]string, error) {
	return []string{"I can't get schedule from the timetable.spbu.ru. Perhaps the TimeTable server is down."}, nil
}

type NotRegistered int

func (notRegistered *NotRegistered) Parse() ([]string, error) {
	return []string{"Schedule is not allowed for you. Please register via the timetable.spbu.ru schedule link."}, nil
}

type GroupEvents spbu_api.GroupEvents

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
		dayParsed = fmt.Sprintf("%s\n\n%s", strings.Title(day.DayString), dayParsed)
		parsed = append(parsed, dayParsed)
	}
	if len(parsed) == 0 {
		return []string{"Nothing to display."}, nil
	}
	return parsed, nil
}

type EducatorEvents spbu_api.EducatorEvents

func (educatorEvents *EducatorEvents) Parse() ([]string, error) {
	var parsed []string
	for _, day := range educatorEvents.EducatorEventsDays {
		if len(day.DayStudyEvents) == 0 {
			continue
		}
		dayParsed := ""
		for _, event := range day.DayStudyEvents {
			dayParsed += fmt.Sprintf("%s\n%s\n%s\n\n",
				event.TimeIntervalString,
				event.Subject,
				event.LocationsDisplayText,
			)
		}
		if dayParsed == "" {
			continue
		}
		dayParsed = fmt.Sprintf("%s\n\n%s", strings.Title(day.DayString), dayParsed)
		parsed = append(parsed, dayParsed)
	}
	if len(parsed) == 0 {
		return []string{"Nothing to display."}, nil
	}
	return parsed, nil
}
