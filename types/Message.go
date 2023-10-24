package types

type Message struct {
	MessageID            int64            `json:"message_id"`
	From                 *User            `json:"from"`
	Date                 int64            `json:"date"`
	Chat                 *Chat            `json:"chat"`
	ForwardFrom          *User            `json:"forward_from"`
	ForwardFromChat      *Chat            `json:"forward_from_chat"`
	ForwardFromMessageID int64            `json:"forward_from_message_id"`
	ForwardSignature     string           `json:"forward_signature"`
	ForwardSenderName    string           `json:"forward_sender_name"`
	ForwardDate          string           `json:"forward_date"`
	IsTopicMessage       bool             `json:"is_topic_message"`
	IsAutomaticForward   bool             `json:"is_automatic_forward"`
	ReplyToMessage       *Message         `json:"reply_to_message,omitempty"`
	ViaBot               *User            `json:"via_bot"`
	EditDate             int64            `json:"edit_date"`
	HasProtectedContent  bool             `json:"has_protected_content"`
	MediaGroupId         string           `json:"media_group_id"`
	AuthorSignature      string           `json:"author_signature"`
	Text                 string           `json:"text"`
	Entities             []*MessageEntity `json:"entities"`
}
