package app

import (
	"os"
	"regexp"
)

var RegExpAllowedTgID = regexp.MustCompile(os.Getenv("ALLOWED_TG_ID"))

var (
	RegExpScheduleLink = regexp.MustCompile(`^(?:https?://)?timetable\.spbu\.ru/(?:[[:alpha:]]+/)?(StudentGroupEvents|(?:Week)?EducatorEvents)(?:/[[:alpha:]]+(?:[?&=a-zA-Z]+studentGroupId)?)?[/=]([[:digit:]]+)(?:/.*)?$`)
	RegExpSubject      = regexp.MustCompile(`^(.*?)(?:, ([^,]*))?$`)
)
