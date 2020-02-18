package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"spbu4u-go/spbu_api"
	"spbu4u-go/telegram_api"
	"strconv"
	"time"
)

type TelegramBot struct {
	DB    *gorm.DB
	Token string
}

func InitTelegramBot(db *gorm.DB) *TelegramBot {
	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegramBotToken == "" {
		log.Fatal("$TELEGRAM_BOT_TOKEN must be set")
	}
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		log.Fatal("$DOMAIN must be set")
	}

	telegramBot := TelegramBot{db, telegramBotToken}
	if os.Getenv("SKIP_TELEGRAM_WEB_HOOK_SET") == "" {
		telegramBot.setWebHook(domain)
	}

	return &telegramBot
}

func (telegramBot *TelegramBot) setWebHook(domain string) {
	url := fmt.Sprintf("https://%s:443/tg/updates", domain)
	log.Println(url)
	webHookConfig := telegram_api.WebHookConfig{
		Url:            url,
		MaxConnections: 40,
		AllowedUpdates: []string{"message"},
	}
	if err := telegram_api.SetWebHookFor(telegramBot.Token, &webHookConfig); err != nil {
		log.Fatal(err)
	}
	if _, err := telegram_api.GetWebHookInfoFor(telegramBot.Token); err != nil {
		log.Fatal(err)
	}
}

func (telegramBot *TelegramBot) handleMessageStart(message *telegram_api.Message) {
	log.Println(message.Chat)
	botMessage := telegram_api.BotMessage{
		ChatID: message.Chat.ID,
		Text: "Send me your schedule link from the timetable.spbu.ru\n" +
			"e.g. https://timetable.spbu.ru/HIST/StudentGroupEvents/Primary/248508",
	}
	if _, err := telegram_api.SendMessageFrom(telegramBot.Token, &botMessage); err != nil {
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
	scheduleType := ScheduleStorageTypeMapper[typeStr]

	//get schedule storage name
	var scheduleStorageName string
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	if scheduleType == ScheduleStorageTypeGroup {
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
	telegramBot.DB.Where(ScheduleStorage{
		TimeTableId: scheduleId,
		Type:        scheduleType,
	}).Assign(ScheduleStorage{
		Name: scheduleStorageName,
	}).FirstOrCreate(&scheduleStorage)

	// update or create user
	var user User
	telegramBot.DB.Where(User{
		TelegramChatID: message.Chat.ID,
	}).Assign(User{
		ScheduleStorageID: scheduleStorage.ID,
	}).FirstOrCreate(&user)

	botMessage := telegram_api.BotMessage{
		ChatID: message.Chat.ID,
		Text:   fmt.Sprintf("Your schedule storage is %s", scheduleStorageName),
	}
	if _, err := telegram_api.SendMessageFrom(telegramBot.Token, &botMessage); err != nil {
		log.Println(err)
	}
}

func (telegramBot *TelegramBot) handleMessageToday(message *telegram_api.Message) {
	var scheduleStorage ScheduleStorage
	telegramBot.DB.Joins(DBQueryGetStorageFor, message.Chat.ID).Find(&scheduleStorage)
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	schedule, err := scheduleStorage.GetSchedule(today, tomorrow)
	if err != nil {
		log.Println(err)
		return
	}
	parsed, err := schedule.Parse()
	if err != nil {
		log.Println(err)
		return
	}
	for _, scheduleText := range parsed {
		botMessage := telegram_api.BotMessage{
			ChatID: message.Chat.ID,
			Text:   scheduleText,
		}
		if _, err := telegram_api.SendMessageFrom(telegramBot.Token, &botMessage); err != nil {
			log.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}

func (telegramBot *TelegramBot) handleMessageTomorrow(message *telegram_api.Message) {
	var scheduleStorage ScheduleStorage
	telegramBot.DB.Joins(DBQueryGetStorageFor, message.Chat.ID).Find(&scheduleStorage)
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	dayAfterTomorrow := today.AddDate(0, 0, 2)
	schedule, err := scheduleStorage.GetSchedule(tomorrow, dayAfterTomorrow)
	if err != nil {
		log.Println(err)
		return
	}
	parsed, err := schedule.Parse()
	if err != nil {
		log.Println(err)
		return
	}
	for _, scheduleText := range parsed {
		botMessage := telegram_api.BotMessage{
			ChatID: message.Chat.ID,
			Text:   scheduleText,
		}
		if _, err := telegram_api.SendMessageFrom(telegramBot.Token, &botMessage); err != nil {
			log.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}

func (telegramBot *TelegramBot) handleMessage(message *telegram_api.Message) {
	log.Println(fmt.Sprintf("HANDLE MESSAGE STARTED: %s", message.Text))
	if message.Text == "/start" {
		telegramBot.handleMessageStart(message)
	} else if match := RegExpScheduleLink.FindStringSubmatch(message.Text); match != nil && len(match) == 3 {
		telegramBot.handleMessageRegisterUrl(message, match...)
	} else if message.Text == "/today" {
		telegramBot.handleMessageToday(message)
	} else if message.Text == "/tomorrow" {
		telegramBot.handleMessageTomorrow(message)
	} else {
		log.Println(message)
	}
}

func (telegramBot *TelegramBot) handleUpdate(update *telegram_api.Update) {
	if update.Message != nil {
		telegramBot.handleMessage(update.Message)
	} else {
		log.Println(update)
	}
}
