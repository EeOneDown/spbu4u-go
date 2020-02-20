package telegram_api

const (
	BaseUrl           = "https://api.telegram.org/bot%s"
	SetWebHookUrl     = BaseUrl + "/setWebhook"
	GetWebHookInfoUrl = BaseUrl + "/getWebhookInfo"
	SendMessage       = BaseUrl + "/sendMessage"
	EditMessageText   = BaseUrl + "/editMessageText"
)

const (
	ParseModeHTML = "HTML"
	//ParseModeMarkdown = "Markdown"
)
