package telegram_api

type WebHookConfig struct {
	Url            string   `json:"url"`
	MaxConnections int      `json:"max_connections"`
	AllowedUpdates []string `json:"allowed_updates"`
}

type Update struct {
	UpdateID int64    `json:"update_id"`
	Message  *Message `json:"message"`
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

type Message struct {
	MessageID      int64    `json:"message_id"`
	Chat           *Chat    `json:"chat"`
	Text           string   `json:"text"`
	ReplyToMessage *Message `json:"reply_to_message"`
}

type Chat struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

type InlineKeyboardButton struct {
	Text                         string    `json:"text"`
	Url                          string    `json:"url,omitempty"`
	LoginUrl                     *LoginUrl `json:"login_url,omitempty"`
	CallbackData                 string    `json:"callback_data,omitempty"`
	SwitchInlineQuery            string    `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string    `json:"switch_inline_query_current_chat,omitempty"`
	//CallbackGame                 CallbackGame `json:"callback_game,omitempty"`
	Pay bool `json:"pay,omitempty"`
}

type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"`
}

type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  bool                    `json:"request_contact,omitempty"`
	RequestLocation bool                    `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`
}

type ReplyMarkup struct {
	InlineKeyboard  [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
	Keyboard        [][]KeyboardButton       `json:"keyboard,omitempty"`
	ResizeKeyboard  bool                     `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool                     `json:"one_time_keyboard,omitempty"`
	Selective       bool                     `json:"selective,omitempty"`
	RemoveKeyboard  bool                     `json:"remove_keyboard,omitempty"`
	ForceReply      bool                     `json:"force_reply,omitempty"`
}

type BotMessage struct {
	ChatID                int64        `json:"chat_id"`
	Text                  string       `json:"text"`
	ParseMode             string       `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool         `json:"disable_notification,omitempty"`
	ReplyToMessageID      int64        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           *ReplyMarkup `json:"reply_markup,omitempty"`
}

type BotEditedMessage struct {
	ChatID                int64  `json:"chat_id,omitempty"`
	MessageID             int64  `json:"message_id,omitempty"`
	InlineMessageID       string `json:"inline_message_id,omitempty"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
}
