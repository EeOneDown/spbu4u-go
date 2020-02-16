package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"spbu4u-go/constants"
	"spbu4u-go/spbu_api"
	"spbu4u-go/telegram_api"
	"strconv"
	"time"
)

type TelegramBot struct {
	DB    *gorm.DB
	Token string
}

type Server struct {
	DB          *gorm.DB
	TelegramBot *TelegramBot
}

func (telegramBot *TelegramBot) setWebHook(domain string) {
	webHookConfig := telegram_api.WebHookConfig{
		Url:            fmt.Sprintf("https://%s:443/tg/updates", domain),
		MaxConnections: 40,
		AllowedUpdates: []string{"message"},
	}
	if err := telegram_api.SetWebHookFor(telegramBot.Token, &webHookConfig); err != nil {
		log.Fatal(err)
	}
}

func (telegramBot *TelegramBot) deleteWebHook() {
	if err := telegram_api.DeleteWebHookFor(telegramBot.Token); err != nil {
		log.Fatal(err)
	}
}

func (telegramBot *TelegramBot) handleMessageStart(message *telegram_api.Message) {
	botMessage := telegram_api.BotMessage{
		ChatID: message.Chat.ChatID,
		Text: "Send me the schedule link from the timetable.spbu.ru\n" +
			"e.g. https://timetable.spbu.ru/HIST/StudentGroupEvents/Primary/248508",
	}
	if err := telegram_api.SendMessageFrom(telegramBot.Token, &botMessage); err != nil {
		log.Println(err)
	}
}

func (telegramBot *TelegramBot) handleMessageRegisterUrl(message *telegram_api.Message, match ...string) {
	typeStr := match[1]
	scheduleId, err := strconv.ParseInt(match[2], 10, 64)
	if err != nil {
		log.Println(match)
		return
	}
	scheduleType := constants.ScheduleTypeMapper[typeStr]

	//get schedule storage name
	var scheduleStorageName string
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	if scheduleType == constants.GROUP {
		res, err := spbu_api.GetGroupScheduleFor(scheduleId, today, tomorrow)
		if err != nil {
			return
		}
		scheduleStorageName = res.StudentGroupDisplayName
	} else {
		res, err := spbu_api.GetEducatorScheduleFor(scheduleId, today, tomorrow)
		if err != nil {
			return
		}
		scheduleStorageName = res.EducatorLongDisplayText
	}

	var scheduleStorage ScheduleStorage
	telegramBot.DB.FirstOrCreate(&scheduleStorage, ScheduleStorage{
		TimeTableId: scheduleId,
		Type:        scheduleType,
		Name:        scheduleStorageName,
	})

	// update or create user
	var user User
	telegramBot.DB.FirstOrCreate(&user, User{
		TelegramChatID:    message.Chat.ChatID,
		ScheduleStorageID: scheduleStorage.ID,
	})

	botMessage := telegram_api.BotMessage{
		ChatID: message.Chat.ChatID,
		Text:   fmt.Sprintf("Your schedule storage is %s", scheduleStorageName),
	}
	if err := telegram_api.SendMessageFrom(telegramBot.Token, &botMessage); err != nil {
		log.Println(err)
	}
}

func (telegramBot *TelegramBot) handleMessage(message *telegram_api.Message) {
	if message.Text == "/start" {
		go telegramBot.handleMessageStart(message)
	} else if match := constants.ScheduleLink.FindStringSubmatch(message.Text); match != nil && len(match) == 3 {
		go telegramBot.handleMessageRegisterUrl(message, match...)
	} else {
		log.Println(message)
	}
}

func (telegramBot *TelegramBot) handleUpdate(update *telegram_api.Update) {
	if update.Message != nil {
		go telegramBot.handleMessage(update.Message)
	} else {
		log.Println(update)
	}
}

func (server *Server) telegramUpdateWebHook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// unmarshal
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var update telegram_api.Update
		if err := json.Unmarshal(data, &update); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		go server.TelegramBot.handleUpdate(&update)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// init helpers
func initDB() *gorm.DB {
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
func initTelegramBot(db *gorm.DB) *TelegramBot {
	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegramBotToken == "" {
		log.Fatal("$TELEGRAM_BOT_TOKEN must be set")
	}
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		log.Fatal("$DOMAIN must be set")
	}

	telegramBot := TelegramBot{db, telegramBotToken}
	telegramBot.setWebHook(domain)

	return &telegramBot
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	db := initDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	telegramBot := initTelegramBot(db)
	defer telegramBot.deleteWebHook()
	server := &Server{db, telegramBot}

	http.HandleFunc("/tg/updates", server.telegramUpdateWebHook)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
