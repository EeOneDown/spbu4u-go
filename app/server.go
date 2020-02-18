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
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	server := Server{
		DB:          db,
		TelegramBot: telegramBot,
	}
	http.HandleFunc("/tg/updates", server.telegramUpdateWebHook)
	// http.HandleFunc("/getTelegramWebHookInfo", server.getTelegramWebHookInfo)
	// http.HandleFunc("/setTelegramWebHook", server.setTelegramWebHook)
	return http.ListenAndServe(":"+port, nil)
}

func (server *Server) getTelegramWebHookInfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		webHookInfo, err := telegram_api.GetWebHookInfoFor(server.TelegramBot.Token)
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
		server.TelegramBot.setWebHook(os.Getenv("DOMAIN"))
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
