package methods

import "go-telegram-bot-api/types"

type SendMessageMethod struct {
	ChatId                int64                 `json:"chat_id"`
	MessageThreadId       int                   `json:"message_thread_id,omitempty"`
	Text                  string                `json:"text"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	Entities              []types.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool                  `json:"disable_notification,omitempty"`
	ProtectContent        bool                  `json:"protect_content,omitempty"`
	ReplyToMessageId      int64                 `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{}           `json:"reply_markup,omitempty"`
}
