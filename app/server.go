package app

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"spbu4u-go/telegram_api"
)

type Server struct {
	DB          *gorm.DB
	TelegramBot *TelegramBot
}

func InitServerAndListen(db *gorm.DB, telegramBot *TelegramBot) error {
	// load required ENVs
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		log.Fatal("$DOMAIN must be set")
	}
	tgSecretPath := os.Getenv("TG_SECRET_PATH")
	if tgSecretPath == "" {
		log.Fatal("$TG_SECRET_PATH must be set")
	}

	// init server
	server := Server{
		DB:          db,
		TelegramBot: telegramBot,
	}

	// setup TG bot
	telegramBot.setWebHook(domain, tgSecretPath)

	// register endpoints
	http.HandleFunc(tgSecretPath, server.telegramUpdateWebHook)
	http.HandleFunc(tgSecretPath+"/getWebHookInfo", server.getTelegramWebHookInfo)
	http.HandleFunc(tgSecretPath+"/setWebHook", server.setTelegramWebHook)
	return http.ListenAndServe(":"+port, nil)
}

func (server *Server) getTelegramWebHookInfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		webHookInfo, err := server.TelegramBot.Bot.GetWebHookInfo()
		if err != nil {
			log.Println(err)
		}
		webHookInfoJson, _ := json.Marshal(webHookInfo)
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(webHookInfoJson); err != nil {
			log.Println(err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (server *Server) setTelegramWebHook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		server.TelegramBot.setWebHook(os.Getenv("DOMAIN"), os.Getenv("TG_SECRET_PATH"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (server *Server) telegramUpdateWebHook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// unmarshal
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var update telegram_api.Update
		if err := json.Unmarshal(data, &update); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		go server.TelegramBot.handleUpdate(&update)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
