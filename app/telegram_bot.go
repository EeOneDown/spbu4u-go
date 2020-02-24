package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"os"
	"spbu4u-go/spbu_api"
	"spbu4u-go/telegram_api"
	"strconv"
	"time"
)

const (
	BotTextDisclaimer = "Это <b>тестовый</b> бот. Для получения доступа свяжитесь с разработчиком."
	BotTextStart      = "Для регистрации отправь мне ссылку на твое расписание на timetable.spbu.ru\n\n" +
		"Например: https://timetable.spbu.ru/HIST/StudentGroupEvents/Primary/248508"
	BotTextRegistering     = "Определяю расписание..."
	BotTextRegisterSuccess = "Твое расписание: <b>%s</b>"
	BotTextMainMenu        = "Главное меню"
)

const (
	BotCommandStart    = "/start"
	BotCommandToday    = "/today"
	BotCommandTomorrow = "/tomorrow"
	BotCommandWeek     = "/week"
	BotCommandWeekNext = "/weeknext"
)

const BotTextSundayScheduleSearching = "Пары в воскресенье?? Ну я гляну, конечно..."

var (
	BotTextSearching = [...]string{
		"Смотрю расписание...",
		"Смотрю расписание на timetable.spbu.ru...",
	}
	BotTextSearchingFun = [...]string{
		"Ищу... Хоть бы выходной...",
		"Поиск расписания активирован...",
		"Призываю Богиню расписания Шедьюлу...",
		"Ахахахаха. Там такое! Не мог не поделиться с другими ботами. Секунду...",
		"Обычно, я рассказываю шутку, но я уже нашёл расписание. Вот оно...",
		"Отправил твое расписание в другой чат. Сейчас извинюсь и отправлю уже тебе...",
	}
	BotTextUnknownMessageReaction = [...]string{
		"Не понимаю.",
		"С такой командой я еще не знаком.",
	}
	BotTextUnknownMessageReactionFun = [...]string{
		"А вот сейчас вообще не понял.",
		"Я бы тебе ответил, да законы робототехники не позволяют.",
		"Я хотел что-то ответить, но забыл что.",
		"Увы, я не чат бот. Давай только по делу.",
		"Когда-то давно, четыре народа жили в мире. Но все изменилось, когда ты начал спамить непонятными сообщениями.",
	}
)

func getSearchingText(from time.Time, to time.Time) string {
	if from.Weekday() == time.Sunday && to.YearDay()-from.YearDay() == 1 {
		return BotTextSundayScheduleSearching
	}
	rand.Seed(time.Now().Unix())
	if chance := rand.Intn(100); chance < 20 {
		return BotTextSearchingFun[rand.Intn(len(BotTextSearchingFun))]
	}
	return BotTextSearching[rand.Intn(len(BotTextSearching))]
}

func getUnknownMessageText() string {
	rand.Seed(time.Now().Unix())
	if chance := rand.Intn(100); chance < 20 {
		return BotTextUnknownMessageReactionFun[rand.Intn(len(BotTextUnknownMessageReactionFun))]
	}
	return BotTextUnknownMessageReaction[rand.Intn(len(BotTextUnknownMessageReaction))]
}

type TelegramBot struct {
	DB  *gorm.DB
	Bot *telegram_api.Bot
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
	bot := &telegram_api.Bot{Token: telegramBotToken}
	telegramBot := TelegramBot{db, bot}
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
	if err := telegramBot.Bot.SetWebHook(&webHookConfig); err != nil {
		log.Fatal(err)
	}
	if _, err := telegramBot.Bot.GetWebHookInfo(); err != nil {
		log.Fatal(err)
	}
}

func (telegramBot *TelegramBot) handleNotAllowed(message *telegram_api.Message) {
	botMessage := telegram_api.BotMessage{
		ChatID:    message.Chat.ID,
		Text:      BotTextDisclaimer,
		ParseMode: telegram_api.ParseModeHTML,
	}
	if _, err := telegramBot.Bot.SendMessage(&botMessage); err != nil {
		log.Println(err)
	}
}

func (telegramBot *TelegramBot) handleMessageStart(message *telegram_api.Message) {
	botMessage := telegram_api.BotMessage{
		ChatID: message.Chat.ID,
		Text:   BotTextStart,
	}
	if _, err := telegramBot.Bot.SendMessage(&botMessage); err != nil {
		log.Println(err)
	}
}

func (telegramBot *TelegramBot) sendMainMenuTo(chat *telegram_api.Chat) {
	botMessage := &telegram_api.BotMessage{
		ChatID: chat.ID,
		Text:   BotTextMainMenu,
		ReplyMarkup: telegram_api.ReplyMarkup{
			Keyboard: [][]telegram_api.KeyboardButton{
				{{Text: BotCommandToday}, {Text: BotCommandTomorrow}, {Text: BotCommandWeek}},
				{{Text: BotCommandStart}, {Text: BotCommandWeekNext}},
			},
			ResizeKeyboard: true,
		},
	}
	if _, err := telegramBot.Bot.SendMessage(botMessage); err != nil {
		log.Println(err)
	}
}

func (telegramBot *TelegramBot) handleMessageRegisterUrl(message *telegram_api.Message, match ...string) {
	botMessageChan := make(chan *telegram_api.Message, 1)
	go func() {
		botMessage := &telegram_api.BotMessage{
			ChatID: message.Chat.ID,
			Text:   BotTextRegistering,
		}
		if err := telegramBot.Bot.SendMessageToEdit(botMessage, botMessageChan); err != nil {
			log.Println(err)
		}
	}()
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
	if scheduleType == ScheduleStorageTypeGroup {
		res, err := spbu_api.GetGroupScheduleFor(scheduleId, today, today)
		if err != nil {
			return
		}
		scheduleStorageName = res.StudentGroupDisplayName
	} else {
		res, err := spbu_api.GetEducatorScheduleFor(scheduleId, today, today)
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

	botEditedMessage := telegram_api.BotEditedMessage{
		ChatID:    message.Chat.ID,
		MessageID: (<-botMessageChan).MessageID,
		Text:      fmt.Sprintf(BotTextRegisterSuccess, scheduleStorageName),
		ParseMode: telegram_api.ParseModeHTML,
	}
	if _, err := telegramBot.Bot.EditMessage(&botEditedMessage); err != nil {
		log.Println(err)
	}
}

func (telegramBot *TelegramBot) sendScheduleTo(chat *telegram_api.Chat, from time.Time, to time.Time) {
	botMessageChan := make(chan *telegram_api.Message, 1)
	go func() {
		botMessage := &telegram_api.BotMessage{
			ChatID: chat.ID,
			Text:   getSearchingText(from, to),
		}
		if err := telegramBot.Bot.SendMessageToEdit(botMessage, botMessageChan); err != nil {
			log.Println(err)
		}
	}()
	var scheduleStorage ScheduleStorage
	telegramBot.DB.Joins(DBQueryGetStorageFor, chat.ID).Find(&scheduleStorage)
	schedule, err := scheduleStorage.GetSchedule(from, to)
	if err != nil {
		log.Println(err)
		return
	}
	parsedChan := make(chan string)
	go schedule.Parse(parsedChan)

	botEditedMessage := telegram_api.BotEditedMessage{
		ChatID:    chat.ID,
		MessageID: (<-botMessageChan).MessageID,
		Text:      <-parsedChan,
		ParseMode: telegram_api.ParseModeHTML,
	}
	if _, err := telegramBot.Bot.EditMessage(&botEditedMessage); err != nil {
		log.Println(err)
	}
	for parsed := range parsedChan {
		botMessage := telegram_api.BotMessage{
			ChatID:    chat.ID,
			Text:      parsed,
			ParseMode: telegram_api.ParseModeHTML,
		}
		if _, err := telegramBot.Bot.SendMessage(&botMessage); err != nil {
			log.Println(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (telegramBot *TelegramBot) handleMessageToday(message *telegram_api.Message) {
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	telegramBot.sendScheduleTo(message.Chat, today, tomorrow)
}

func (telegramBot *TelegramBot) handleMessageTomorrow(message *telegram_api.Message) {
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	dayAfterTomorrow := today.AddDate(0, 0, 2)
	telegramBot.sendScheduleTo(message.Chat, tomorrow, dayAfterTomorrow)
}

func (telegramBot *TelegramBot) handleMessageWeek(message *telegram_api.Message) {
	today := time.Now()
	monday := today.AddDate(0, 0, 1-int(today.Weekday()))
	sunday := monday.AddDate(0, 0, 6)
	telegramBot.sendScheduleTo(message.Chat, monday, sunday)
}

func (telegramBot *TelegramBot) handleMessageWeekNext(message *telegram_api.Message) {
	today := time.Now()
	monday := today.AddDate(0, 0, 8-int(today.Weekday()))
	sunday := monday.AddDate(0, 0, 6)
	telegramBot.sendScheduleTo(message.Chat, monday, sunday)
}

func (telegramBot *TelegramBot) handleMessageUnknown(message *telegram_api.Message) {
	botMessage := telegram_api.BotMessage{
		ChatID:    message.Chat.ID,
		Text:      getUnknownMessageText(),
		ParseMode: telegram_api.ParseModeHTML,
	}
	if _, err := telegramBot.Bot.SendMessage(&botMessage); err != nil {
		log.Println(err)
	}
}

func (telegramBot *TelegramBot) handleMessage(message *telegram_api.Message) {
	// todo: remove after release
	if match := RegExpAllowedTgID.FindStringSubmatch(strconv.FormatInt(message.Chat.ID, 10)); match == nil {
		telegramBot.handleNotAllowed(message)
	} else if message.Text == BotCommandStart {
		telegramBot.handleMessageStart(message)
	} else if match := RegExpScheduleLink.FindStringSubmatch(message.Text); match != nil && len(match) == 3 {
		telegramBot.handleMessageRegisterUrl(message, match...)
	} else if message.Text == BotCommandToday {
		telegramBot.handleMessageToday(message)
	} else if message.Text == BotCommandTomorrow {
		telegramBot.handleMessageTomorrow(message)
	} else if message.Text == BotCommandWeek {
		telegramBot.handleMessageWeek(message)
	} else if message.Text == BotCommandWeekNext {
		telegramBot.handleMessageWeekNext(message)
	} else {
		telegramBot.handleMessageUnknown(message)
	}
}

func (telegramBot *TelegramBot) handleUpdate(update *telegram_api.Update) {
	if update.Message != nil {
		telegramBot.handleMessage(update.Message)
	} else {
		log.Println(update)
	}
}
