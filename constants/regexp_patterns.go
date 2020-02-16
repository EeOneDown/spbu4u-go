package constants

import "regexp"

var (
	ScheduleLink = regexp.MustCompile(`^(?:https?://)?timetable\.spbu\.ru/(?:[[:alpha:]]+/)?(StudentGroupEvents|(?:Week)?EducatorEvents)(?:/[[:alpha:]]+(?:[?&=a-zA-Z]+studentGroupId)?)?[/=]([[:digit:]]+)(?:/.*)?$`)
)
