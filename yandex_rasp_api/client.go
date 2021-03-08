package yandex_rasp_api

import (
	"net/http"
	"time"
)

const (
	BaseUrl = "https://api.rasp.yandex.net/v3.0"
	//Copyright = BaseUrl + "/copyright/"
	SchedulePointPoint = BaseUrl + "/search/"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}
