package constants

const (
	GROUP    uint8 = 0
	EDUCATOR uint8 = 1
)

var (
	ScheduleTypeMapper = map[string]uint8{
		"StudentGroupEvents": GROUP,
		"EducatorEvents":     EDUCATOR,
		"WeekEducatorEvents": EDUCATOR,
	}
)
