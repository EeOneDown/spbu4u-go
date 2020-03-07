package app

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	TelegramChatID    int64 `gorm:"unique_index;not null"`
	ScheduleStorageID uint  `gorm:"not null"`
	ScheduleStorage   ScheduleStorage
	// save schedules (m2m)
}
