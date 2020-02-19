package app

import (
	"fmt"
	"spbu4u-go/spbu_api"
	"strings"
)

type Schedule interface {
	Parse() []string
}

type ScheduleNotAllowed int

func (scheduleNotAllowed *ScheduleNotAllowed) Parse() []string {
	return []string{"I can't get schedule from the timetable.spbu.ru. Perhaps the TimeTable server is down."}
}

type NotRegistered int

func (notRegistered *NotRegistered) Parse() []string {
	return []string{"Schedule is not allowed for you. Please register via the timetable.spbu.ru schedule link."}
}

type GroupEvents spbu_api.GroupEvents

func (groupEvents *GroupEvents) Parse() []string {
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
		return []string{"Nothing to display."}
	}
	return parsed
}

type EducatorEvents spbu_api.EducatorEvents

func (educatorEvents *EducatorEvents) Parse() []string {
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
		return []string{"Nothing to display."}
	}
	return parsed
}
