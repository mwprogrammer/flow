package payloads

type BasePayload struct {
	Product        string `json:"messaging_product"`
	ReceipientType string `json:"receipient_type"`
	To             string `json:"to"`
	Type           string `json:"type"`
}

type MarkAsReadPayload struct {
	BasePayload
	MessageId string `json:"message_id"`
	Status    string `json:"status"`
}

type DisplayTypingPayload struct {
	MarkAsReadPayload
	Indicator TypingIndicator `json:"typing_indicator"`
}

type TypingIndicator struct {
	Type string `json:"type"`
}

type ReplyWithTextPayload struct {
	BasePayload
	TextMessage *TextMessage `json:"text"`
}

type TextMessage struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}
