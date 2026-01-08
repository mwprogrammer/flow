package client

// BaseResponse represents the root object of a webhook notification
// from the WhatsApp Business Platform.
type BaseResponse struct {
	Object  string  `json:"object"`
	Entries []Entry `json:"entry"`
}

// Entry contains the changes for a specific WhatsApp Business Account.
type Entry struct {
	ID      string   `json:"id"`
	Changes []Change `json:"changes"`
}

// Change represents a specific update field and its associated data.
type Change struct {
	Field string `json:"field"`
	Value any    `json:"value"`
}

// Recipient contains the profile and contact information of a WhatsApp user.
type Recipient struct {
	Profile            Profile `json:"profile"`
	WhatsAppBusinessID string  `json:"wa_id"`
}

// Profile represents the public-facing information of a WhatsApp user.
type Profile struct {
	Name string `json:"name"`
}

// RecipientMessage contains the common metadata for an incoming message.
type RecipientMessage struct {
	From      string `json:"from"`
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
}

// TextRecipientMessage represents an incoming message that contains text content.
type TextRecipientMessage struct {
	RecipientMessage
	Text TextContent `json:"text"`
}

// TextContent contains the body of a text-based message.
type TextContent struct {
	Body string `json:"body"`
}
