package main

import (
	"spbu4u-go/constants"
	"spbu4u-go/spbu_api"
	"spbu4u-go/spbu_api/types"
	"time"
)

type Schedule interface {
	Parse() ([]string, error)
}

func (scheduleStorage *ScheduleStorage) GetScheduleFor(from time.Time, to time.Time) (Schedule, error) {
	if scheduleStorage.ID == 0 {
		var notRegistered *NotRegistered
		var schedule Schedule = notRegistered
		return schedule, nil
	} else if scheduleStorage.Type == constants.GROUP {
		groupEvents, err := spbu_api.GetGroupScheduleFor(scheduleStorage.TimeTableId, from, to)
		if err != nil {
			var scheduleNotAllowed *ScheduleNotAllowed
			var schedule Schedule = scheduleNotAllowed
			return schedule, err
		}
		var schedule Schedule = (*GroupEvents)(groupEvents)
		return schedule, nil
	} else {
		educatorEvents, err := spbu_api.GetEducatorScheduleFor(scheduleStorage.TimeTableId, from, to)
		if err != nil {
			var scheduleNotAllowed *ScheduleNotAllowed
			var schedule Schedule = scheduleNotAllowed
			return schedule, err
		}
		var schedule Schedule = (*EducatorEvents)(educatorEvents)
		return schedule, nil
	}
}

type ScheduleNotAllowed int

func (scheduleNotAllowed *ScheduleNotAllowed) Parse() ([]string, error) {
	return []string{"I can't get schedule from the timetable.spbu.ru. Perhaps the TimeTable server is down."}, nil
}

type NotRegistered int

func (notRegistered *NotRegistered) Parse() ([]string, error) {
	return []string{"Schedule is not allowed for you. Please register via the timetable.spbu.api schedule link."}, nil
}

type GroupEvents types.GroupEvents

func (groupEvents *GroupEvents) Parse() ([]string, error) {
	return []string{"Working on it."}, nil
}

type EducatorEvents types.EducatorEvents

func (educatorEvents *EducatorEvents) Parse() ([]string, error) {
	return []string{"Working on it."}, nil
}