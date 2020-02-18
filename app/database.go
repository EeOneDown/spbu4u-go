package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func InitDB() *gorm.DB {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("$DATABASE_URL must be set")
	}
	db, err := gorm.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.DropTableIfExists(&User{}, &ScheduleStorage{}) // todo: remove
	db.AutoMigrate(&User{}, &ScheduleStorage{})
	return db
}
