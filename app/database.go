package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

const DBQueryGetStorageFor = "JOIN users ON users.schedule_storage_id = schedule_storages.id AND users.telegram_chat_id = ?"

func InitDB() *gorm.DB {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("$DATABASE_URL must be set")
	}
	db, err := gorm.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&User{}, &ScheduleStorage{})
	return db
}
