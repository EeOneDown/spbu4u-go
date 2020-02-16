package tests

import (
	"spbu4u-go/constants"
	"strconv"
	"testing"
)

func TestFindScheduleFromGroupUrl(t *testing.T) {
	t.Parallel()
	goodUrls := []string{
		"https://timetable.spbu.ru/BIOL/StudentGroupEvents/Online/247986",
		"timetable.spbu.ru/BIOL/StudentGroupEvents/Attestation/247986",
		"https://timetable.spbu.ru/BIOL/StudentGroupEvents/Final/247986",
		"https://timetable.spbu.ru/BIOL/StudentGroupEvents/Semester/247986/1",
		"timetable.spbu.ru/BIOL/StudentGroupEvents/Primary/247986",
		"timetable.spbu.ru/BIOL/StudentGroupEvents/Primary/247986/2020-02-17",
		"https://timetable.spbu.ru/StudentGroupEvents/Semester?publicDivisionAlias=BIOL&studentGroupId=247986",
	}
	const goodTypeStr = "StudentGroupEvents"
	const goodScheduleId = 247986
	const goodScheduleType = constants.GROUP
	for _, goodUrl := range goodUrls {
		match := constants.ScheduleLink.FindStringSubmatch(goodUrl)
		if match == nil || len(match) != 3 {
			t.Fail()
			continue
		}
		typeStr := match[1]
		if typeStr != goodTypeStr {
			t.Fail()
			continue
		}
		scheduleId, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil || scheduleId != goodScheduleId {
			t.Fail()
			continue
		}
		scheduleType := constants.ScheduleTypeMapper[typeStr]
		if scheduleType != goodScheduleType {
			t.Fail()
			continue
		}
	}
}

func TestFindScheduleFromEducatorUrl(t *testing.T) {
	t.Parallel()
	goodUrls := []string{
		"timetable.spbu.ru/EducatorEvents/1420",
		"https://timetable.spbu.ru/EducatorEvents/1420/1",
	}
	const goodTypeStr = "EducatorEvents"
	const goodScheduleId = 1420
	const goodScheduleType = constants.EDUCATOR
	for _, goodUrl := range goodUrls {
		match := constants.ScheduleLink.FindStringSubmatch(goodUrl)
		if match == nil || len(match) != 3 {
			t.Fail()
			continue
		}
		typeStr := match[1]
		if typeStr != goodTypeStr {
			t.Fail()
			continue
		}
		scheduleId, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil || scheduleId != goodScheduleId {
			t.Fail()
			continue
		}
		scheduleType := constants.ScheduleTypeMapper[typeStr]
		if scheduleType != goodScheduleType {
			t.Fail()
			continue
		}
	}
}

func TestFindScheduleFromWeekEducatorUrl(t *testing.T) {
	t.Parallel()
	goodUrls := []string{
		"timetable.spbu.ru/WeekEducatorEvents/1420",
		"https://timetable.spbu.ru/WeekEducatorEvents/1420/2020-02-17",
	}
	const goodTypeStr = "WeekEducatorEvents"
	const goodScheduleId = 1420
	const goodScheduleType = constants.EDUCATOR
	for _, goodUrl := range goodUrls {
		match := constants.ScheduleLink.FindStringSubmatch(goodUrl)
		if match == nil || len(match) != 3 {
			t.Fail()
			continue
		}
		typeStr := match[1]
		if typeStr != goodTypeStr {
			t.Fail()
			continue
		}
		scheduleId, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil || scheduleId != goodScheduleId {
			t.Fail()
			continue
		}
		scheduleType := constants.ScheduleTypeMapper[typeStr]
		if scheduleType != goodScheduleType {
			t.Fail()
			continue
		}
	}
}
