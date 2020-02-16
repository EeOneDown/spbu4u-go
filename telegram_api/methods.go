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

func GetWebHookInfoFor(token string) error {
	resp, err := http.Get(fmt.Sprintf(GetWebHookInfoUrl, token))
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var webHookInfo WebHookInfo
	if err := json.Unmarshal(body, &webHookInfo); err != nil {
		return err
	}
	log.Println(webHookInfo)
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
