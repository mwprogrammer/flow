package client

// BasePayload contains the common fields required for every request
// sent to the WhatsApp Business Platform.
type BasePayload struct {
	// Product is the messaging service being used (usually "whatsapp").
	Product       string `json:"messaging_product"`
	// RecipientType defines the type of recipient (usually "individual").
	RecipientType string `json:"recipient_type"`
	// To is the WhatsApp ID (phone number) of the message recipient.
	To            string `json:"to"`
	// Type defines the type of message being sent (e.g., "text").
	Type          string `json:"type"`
}

// MarkAsReadPayload represents the payload used to manually mark 
// a received message as read.
type MarkAsReadPayload struct {
	BasePayload
	// MessageID is the unique identifier of the message to be marked as read.
	MessageID string `json:"message_id"`
	// Status must be set to "read" for this operation.
	Status    string `json:"status"`
}

// DisplayTypingPayload represents the payload used to trigger 
// a typing indicator in the user's chat window.
type DisplayTypingPayload struct {
	MarkAsReadPayload
	// Indicator contains the configuration for the typing animation.
	Indicator TypingIndicator `json:"typing_indicator"`
}

// TypingIndicator defines the visual feedback shown to the user.
type TypingIndicator struct {
	// Type is the kind of indicator (e.g., "typing").
	Type string `json:"type"`
}

// ReplyWithTextPayload represents the payload for sending a 
// standard text message response.
type ReplyWithTextPayload struct {
	BasePayload
	// TextMessage contains the specific content of the text message.
	TextMessage *TextMessage `json:"text"`
}

// TextMessage contains the body text and metadata for a text-based message.
type TextMessage struct {
	// PreviewURL determines if the message should include a URL preview.
	PreviewURL bool   `json:"preview_url"`
	// Body is the actual text content of the message.
	Body       string `json:"body"`
}