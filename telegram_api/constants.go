package telegram_api

const (
	BaseUrl           = "https://api.telegram.org/bot%s"
	SetWebHookUrl     = BaseUrl + "/setWebhook"
	GetWebHookInfoUrl = BaseUrl + "/getWebhookInfo"
	DeleteWebHookUrl  = BaseUrl + "/deleteWebhook"
)
