package telegram_api

type Chat struct {
	ChatID    int64  `json:"chat_id"`
	Type      string `json:"type"`
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Message struct {
	MessageID      int64    `json:"message_id"`
	Chat           *Chat    `json:"chat"`
	Text           string   `json:"text"`
	ReplyToMessage *Message `json:"reply_to_message"`
}

type Update struct {
	UpdateID int64    `json:"update_id"`
	Message  *Message `json:"message"`
}

type WebHookConfig struct {
	Url            string   `json:"url"`
	MaxConnections int      `json:"max_connections"`
	AllowedUpdates []string `json:"allowed_updates"`
}

type WebHookInfo struct {
	Url                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int64    `json:"pending_update_count"`
	LastErrorDate        int64    `json:"last_error_date"`
	LastErrorMessage     string   `json:"last_error_message"`
	MaxConnections       int64    `json:"max_connections"`
	AllowedUpdates       []string `json:"allowed_updates"`
}

type BotMessage struct {
	ChatID                int64  `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool   `json:"disable_notification,omitempty"`
	ReplyToMessageID      int64  `json:"reply_to_message_id,omitempty"`
}
