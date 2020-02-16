package main

import "github.com/jinzhu/gorm"

type ScheduleStorage struct {
	gorm.Model
	TimeTableId int64  `gorm:"not null;unique_index:idx_ttid_type"`
	Type        uint8  `gorm:"not null;unique_index:idx_ttid_type"`
	Name        string `gorm:"not null"`
	Users       []User
}

type User struct {
	gorm.Model
	TelegramChatID    int64 `gorm:"unique_index;not null"`
	ScheduleStorageID uint  `gorm:"not null"`
	// save schedules (m2m)
}
