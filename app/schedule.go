package app

import (
	"fmt"
	"regexp"
	"spbu4u-go/spbu_api"
	"strings"
	"time"
)

const (
	RemoteLocation      = "Дистанционная работа с преподавателем,С использованием информационно-коммуникационных технологий"
	RemoteLocationShort = "Дистанционка"
)

const (
	EmojiCalendar = "\U0001F4C5"
	EmojiClock3   = "\U0001F552"
	EmojiSleeping = "\U0001F634"
)

const (
	BotTextScheduleServerIsNotAvailable = "Не удается получить расписание с timetable.spbu.ru." +
		" Возможно, TimeTable временно отключен."
	BotTextScheduleIsNotAllowed = "Просмотр расписания недоступен. Отправь timetable.spbu.ru ссылку на свое расписание."
	BotTextWeekend              = EmojiSleeping + " Выходной"
)

var RegExpSubject = regexp.MustCompile(`^(.*?)(?:, ([^,]*))?$`)

type Schedule interface {
	Parse(parsedChan chan<- string)
}

type ScheduleNotAvailable int

func (scheduleNotAvailable *ScheduleNotAvailable) Parse(parsedChan chan<- string) {
	parsedChan <- BotTextScheduleServerIsNotAvailable
	close(parsedChan)
}

type ScheduleNotAllowed int

func (scheduleNotAllowed *ScheduleNotAllowed) Parse(parsedChan chan<- string) {
	parsedChan <- BotTextScheduleIsNotAllowed
	close(parsedChan)
}

func parseTime(event *spbu_api.Event, timeChan chan<- string) {
	start := time.Time(event.Start)
	end := time.Time(event.End)
	timeChan <- fmt.Sprintf("%s %02d:%02d - %02d:%02d",
		EmojiClock3,
		start.Hour(),
		start.Minute(),
		end.Hour(),
		end.Minute(),
	)
	close(timeChan)
}

func parseSubject(event *spbu_api.Event, subjectChan chan<- string) {
	match := RegExpSubject.FindStringSubmatch(event.Subject)
	if match == nil || len(match) != 3 || match[2] == "" {
		subjectChan <- fmt.Sprintf("<b>%s</b>", event.Subject)
		close(subjectChan)
		return
	}
	subjectTitle, subjectType := match[1], match[2]
	var subjectTypeShort string
	switch subjectType {
	case "лекция":
		subjectTypeShort = "Л"
	case "практическое занятие":
		subjectTypeShort = "ПР"
	case "семинар":
		subjectTypeShort = "С"
	case "сам. работа в присутствии преподавателя":
		subjectTypeShort = "СР"
	case "консультация групповая":
		subjectTypeShort = "КОНС"
	// todo: add more
	default:
		subjectTypeShort = strings.ToUpper(subjectType)
	}
	subjectChan <- fmt.Sprintf("<b>%s - %s</b>", subjectTypeShort, subjectTitle)
	close(subjectChan)
}

func parseGroupLocations(event *spbu_api.Event, locationsChan chan<- string) {
	parsedLocations := ""
	for _, location := range event.EventLocations {
		parsedEducators := ""
		lastIndex := len(location.EducatorIds) - 1
		for i, educator := range location.EducatorIds {
			parsedEducators += educator.Name
			if i != lastIndex {
				parsedEducators += "; "
			}
		}
		locationName := strings.TrimSpace(location.DisplayName)
		if locationName == RemoteLocation {
			locationName = RemoteLocationShort
		}
		if parsedEducators != "" {
			parsedLocations += fmt.Sprintf("%s <i>(%s)</i>\n", locationName, parsedEducators)
		} else {
			parsedLocations += fmt.Sprintf("%s\n", locationName)
		}
	}
	locationsChan <- parsedLocations
	close(locationsChan)
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
			if event.IsCancelled {
				continue
			}
			timeChan := make(chan string, 1)
			go parseTime(&event, timeChan)
			subjectChan := make(chan string, 1)
			go parseSubject(&event, subjectChan)
			locationsChan := make(chan string, 1)
			go parseGroupLocations(&event, locationsChan)
			dayParsed += fmt.Sprintf("%s\n%s\n%s\n", <-timeChan, <-subjectChan, <-locationsChan)
		}
		if dayParsed == "" {
			continue
		}
		parsedCount += 1
		parsedChan <- fmt.Sprintf("%s %s\n\n%s", EmojiCalendar, strings.Title(day.DayString), dayParsed)
	}
	if parsedCount == 0 {
		parsedChan <- BotTextWeekend
	}
	close(parsedChan)
}

func parseEducatorLocations(event *spbu_api.Event, locationsChan chan<- string) {
	eventLocation := strings.TrimSpace(event.LocationsDisplayText)
	if eventLocation == RemoteLocation {
		eventLocation = RemoteLocationShort
	}
	if event.ContingentUnitName != "" {
		locationsChan <- fmt.Sprintf("%s <i>(%s)</i>\n", eventLocation, event.ContingentUnitName)
	} else {
		locationsChan <- fmt.Sprintf("%s\n", eventLocation)
	}
	close(locationsChan)
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
			if event.IsCancelled {
				continue
			}
			timeChan := make(chan string, 1)
			go parseTime(&event, timeChan)
			subjectChan := make(chan string, 1)
			go parseSubject(&event, subjectChan)
			locationsChan := make(chan string, 1)
			go parseEducatorLocations(&event, locationsChan)
			dayParsed += fmt.Sprintf("%s\n%s\n%s\n", <-timeChan, <-subjectChan, <-locationsChan)
		}
		if dayParsed == "" {
			continue
		}
		parsedCount += 1
		parsedChan <- fmt.Sprintf("%s %s\n\n%s", EmojiCalendar, strings.Title(day.DayString), dayParsed)
	}
	if parsedCount == 0 {
		parsedChan <- BotTextWeekend
	}
	close(parsedChan)
}
