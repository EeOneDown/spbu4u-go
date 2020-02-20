package telegram_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bot struct {
	Token string
}

func (bot *Bot) SetWebHook(webHookConfig *WebHookConfig) error {
	data, err := json.Marshal(webHookConfig)
	if err != nil {
		return err
	}
	r := bytes.NewReader(data)
	_, err = http.Post(fmt.Sprintf(SetWebHookUrl, bot.Token), "application/json", r)
	if err != nil {
		return err
	}
	return nil
}

func (bot *Bot) GetWebHookInfo() (*WebHookInfo, error) {
	type WebHookInfoResponse struct {
		Ok     bool         `json:"ok"`
		Result *WebHookInfo `json:"result"`
	}
	resp, err := http.Get(fmt.Sprintf(GetWebHookInfoUrl, bot.Token))
	if err != nil {
		return &WebHookInfo{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &WebHookInfo{}, err
	}
	var webHookInfoResponse WebHookInfoResponse
	if err := json.Unmarshal(body, &webHookInfoResponse); err != nil {
		return &WebHookInfo{}, err
	}
	log.Println(webHookInfoResponse.Ok, *webHookInfoResponse.Result)
	return webHookInfoResponse.Result, nil
}

func (bot *Bot) SendMessage(message *BotMessage) (*Message, error) {
	type MessageResponse struct {
		Ok     bool     `json:"ok"`
		Result *Message `json:"result"`
	}
	data, err := json.Marshal(message)
	if err != nil {
		return &Message{}, err
	}
	r := bytes.NewReader(data)
	resp, err := http.Post(fmt.Sprintf(SendMessage, bot.Token), "application/json", r)
	if err != nil {
		return &Message{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Message{}, err
	}
	var messageResponse MessageResponse
	if err := json.Unmarshal(body, &messageResponse); err != nil {
		return &Message{}, err
	}
	return messageResponse.Result, nil
}

func (bot *Bot) SendMessageToEdit(message *BotMessage, botMessageChan chan<- *Message) error {
	type MessageResponse struct {
		Ok     bool     `json:"ok"`
		Result *Message `json:"result"`
	}
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	r := bytes.NewReader(data)
	resp, err := http.Post(fmt.Sprintf(SendMessage, bot.Token), "application/json", r)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var messageResponse MessageResponse
	if err := json.Unmarshal(body, &messageResponse); err != nil {
		return err
	}
	botMessageChan <- messageResponse.Result
	close(botMessageChan)
	return nil
}

func (bot *Bot) EditMessage(message *BotEditedMessage) (*Message, error) {
	type MessageResponse struct {
		Ok     bool     `json:"ok"`
		Result *Message `json:"result"`
	}
	data, err := json.Marshal(message)
	if err != nil {
		return &Message{}, err
	}
	r := bytes.NewReader(data)
	resp, err := http.Post(fmt.Sprintf(EditMessageText, bot.Token), "application/json", r)
	if err != nil {
		return &Message{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Message{}, err
	}
	var messageResponse MessageResponse
	if err := json.Unmarshal(body, &messageResponse); err != nil {
		return &Message{}, err
	}
	return messageResponse.Result, nil
}
