package client

// BasePayload contains the common fields required for every request
// sent to the WhatsApp Business Platform.
type BasePayload struct {
	Product       string `json:"messaging_product"`
	RecipientType string `json:"recipient_type"`
	To            string `json:"to"`
	Type          string `json:"type"`
}

// MarkAsReadPayload represents the payload used to manually mark
// a received message as read.
type MarkAsReadPayload struct {
	BasePayload
	MessageID string `json:"message_id"`
	Status    string `json:"status"`
}

// DisplayTypingPayload represents the payload used to trigger
// a typing indicator in the user's chat window.
type DisplayTypingPayload struct {
	MarkAsReadPayload
	Indicator TypingIndicator `json:"typing_indicator"`
}

// TypingIndicator defines the visual feedback shown to the user.
type TypingIndicator struct {
	Type string `json:"type"`
}

// ReplyWithTextPayload represents the payload for sending a
// standard text message response.
type ReplyWithTextPayload struct {
	BasePayload
	TextMessage *TextMessage `json:"text"`
}

// TextMessage contains the body text and metadata for a text-based message.
type TextMessage struct {
	PreviewURL bool   `json:"preview_url"`
	Body       string `json:"body"`
}
