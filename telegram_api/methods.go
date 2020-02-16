package telegram_api

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func DeleteWebHookFor(token string) error {
	_, err := http.Get(fmt.Sprintf(DeleteWebHookUrl, token))
	if err != nil {
		return err
	}
	return nil
}

func SendMessageFrom(token string, message *BotMessage) error {
	data, err := json.Marshal(message)
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
