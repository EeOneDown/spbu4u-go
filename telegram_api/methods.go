package telegram_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func SetWebHookFor(token string, webHookConfig *WebHookConfig) error {
	data, err := json.Marshal(webHookConfig)
	if err != nil {
		return err
	}
	r := bytes.NewReader(data)
	_, err = http.Post(fmt.Sprintf(SetWebHookUrl, token), "application/json", r)
	if err != nil {
		return err
	}
	return nil
}

func GetWebHookInfoFor(token string) (*WebHookInfo, error) {
	type WebHookInfoResponse struct {
		Ok     bool         `json:"ok"`
		Result *WebHookInfo `json:"result"`
	}
	resp, err := http.Get(fmt.Sprintf(GetWebHookInfoUrl, token))
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

func SendMessageFrom(token string, message *BotMessage) (*Message, error) {
	type MessageResponse struct {
		Ok     bool     `json:"ok"`
		Result *Message `json:"result"`
	}
	data, err := json.Marshal(message)
	if err != nil {
		return &Message{}, err
	}
	log.Println(data)
	r := bytes.NewReader(data)
	resp, err := http.Post(fmt.Sprintf(SendMessage, token), "application/json", r)
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
	log.Println(messageResponse.Ok, messageResponse.Result)
	return messageResponse.Result, nil
}
