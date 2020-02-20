package app

import (
	"fmt"
	"spbu4u-go/spbu_api"
	"strings"
)

type Schedule interface {
	Parse(parsedChan chan<- string)
}

type ScheduleNotAllowed int

func (scheduleNotAllowed *ScheduleNotAllowed) Parse(parsedChan chan<- string) {
	parsedChan <- "I can't get schedule from the timetable.spbu.ru. Perhaps the TimeTable server is down."
	close(parsedChan)
}

type NotRegistered int

func (notRegistered *NotRegistered) Parse(parsedChan chan<- string) {
	parsedChan <- "Schedule is not allowed for you. Please register via the timetable.spbu.ru schedule link."
	close(parsedChan)
}

type GroupEvents spbu_api.GroupEvents

func (groupEvents *GroupEvents) Parse(parsedChan chan<- string) {
	parsedCount := 0
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
		parsedCount += 1
		parsedChan <- fmt.Sprintf("%s\n\n%s", strings.Title(day.DayString), dayParsed)
	}
	if parsedCount == 0 {
		parsedChan <- "Nothing to display."
	}
	close(parsedChan)
}

type EducatorEvents spbu_api.EducatorEvents

func (educatorEvents *EducatorEvents) Parse(parsedChan chan<- string) {
	parsedCount := 0
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
		parsedCount += 1
		parsedChan <- fmt.Sprintf("%s\n\n%s", strings.Title(day.DayString), dayParsed)
	}
	if parsedCount == 0 {
		parsedChan <- "Nothing to display."
	}
	close(parsedChan)
}
