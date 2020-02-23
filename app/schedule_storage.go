package app

import (
	"github.com/jinzhu/gorm"
	"spbu4u-go/spbu_api"
	"time"
)

const (
	ScheduleStorageTypeGroup uint8 = iota + 1
	ScheduleStorageTypeEducator
)

var ScheduleStorageTypeMapper = map[string]uint8{
	"StudentGroupEvents": ScheduleStorageTypeGroup,
	"EducatorEvents":     ScheduleStorageTypeEducator,
	"WeekEducatorEvents": ScheduleStorageTypeEducator,
}

type ScheduleStorage struct {
	gorm.Model
	TimeTableId int64  `gorm:"not null;unique_index:idx_ttid_type"`
	Type        uint8  `gorm:"not null;unique_index:idx_ttid_type"`
	Name        string `gorm:"not null"`
	Users       []User
}

func (scheduleStorage *ScheduleStorage) GetSchedule(from time.Time, to time.Time) (Schedule, error) {
	switch scheduleStorage.Type {
	case ScheduleStorageTypeGroup:
		groupEvents, err := spbu_api.GetGroupScheduleFor(scheduleStorage.TimeTableId, from, to)
		if err != nil {
			var scheduleNotAllowed *ScheduleNotAllowed
			var schedule Schedule = scheduleNotAllowed
			return schedule, err
		}
		var schedule Schedule = (*GroupEvents)(groupEvents)
		return schedule, nil
	case ScheduleStorageTypeEducator:
		educatorEvents, err := spbu_api.GetEducatorScheduleFor(scheduleStorage.TimeTableId, from, to)
		if err != nil {
			var scheduleNotAllowed *ScheduleNotAllowed
			var schedule Schedule = scheduleNotAllowed
			return schedule, err
		}
		var schedule Schedule = (*EducatorEvents)(educatorEvents)
		return schedule, nil
	default:
		var notRegistered *NotRegistered
		var schedule Schedule = notRegistered
		return schedule, nil
	}
}
