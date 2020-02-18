package app

import "regexp"

const (
	ScheduleStorageTypeGroup    uint8 = 0
	ScheduleStorageTypeEducator uint8 = 1
	DBQueryGetStorageFor              = "JOIN users ON users.schedule_storage_id = schedule_storages.id AND users.telegram_chat_id = ?"
)

var (
	ScheduleStorageTypeMapper = map[string]uint8{
		"StudentGroupEvents": ScheduleStorageTypeGroup,
		"EducatorEvents":     ScheduleStorageTypeEducator,
		"WeekEducatorEvents": ScheduleStorageTypeEducator,
	}
	RegExpScheduleLink = regexp.MustCompile(`^(?:https?://)?timetable\.spbu\.ru/(?:[[:alpha:]]+/)?(StudentGroupEvents|(?:Week)?EducatorEvents)(?:/[[:alpha:]]+(?:[?&=a-zA-Z]+studentGroupId)?)?[/=]([[:digit:]]+)(?:/.*)?$`)
)
