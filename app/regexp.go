package app

import "regexp"

var (
	RegExpScheduleLink = regexp.MustCompile(`^(?:https?://)?timetable\.spbu\.ru/(?:[[:alpha:]]+/)?(StudentGroupEvents|(?:Week)?EducatorEvents)(?:/[[:alpha:]]+(?:[?&=a-zA-Z]+studentGroupId)?)?[/=]([[:digit:]]+)(?:/.*)?$`)
	RegExpSubject      = regexp.MustCompile(`^(.*?)(?:, ([^,]*))?$`)
)
