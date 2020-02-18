package main

import (
	"log"
	"spbu4u-go/app"
)

func main() {
	db := app.InitDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	telegramBot := app.InitTelegramBot(db)

	log.Fatal(app.InitServerAndListen(db, telegramBot))
}
