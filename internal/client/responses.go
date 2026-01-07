package client

// BaseResponse represents the root object of a webhook notification
// from the WhatsApp Business Platform.
type BaseResponse struct {
	// Object describes the type of webhook (usually "whatsapp_business_account").
	Object  string  `json:"object"`
	// Entries contains the specific updates for this notification.
	Entries []Entry `json:"entry"`
}

// Entry contains the changes for a specific WhatsApp Business Account.
type Entry struct {
	// ID is the WhatsApp Business Account ID.
	ID      string   `json:"id"`
	// Changes is a list of updates occurring in this entry.
	Changes []Change `json:"changes"`
}

// Change represents a specific update field and its associated data.
type Change struct {
	// Field identifies the type of change (e.g., "messages").
	Field string `json:"field"`
	// Value contains the actual data for the change.
	Value any    `json:"value"`
}

// Recipient contains the profile and contact information of a WhatsApp user.
type Recipient struct {
	// Profile contains the user's public profile details.
	Profile            Profile `json:"profile"`
	// WhatsAppBusinessID is the user's unique WhatsApp ID (phone number).
	WhatsAppBusinessID string  `json:"wa_id"`
}

// Profile represents the public-facing information of a WhatsApp user.
type Profile struct {
	// Name is the display name the user has configured in WhatsApp.
	Name string `json:"name"`
}

// RecipientMessage contains the common metadata for an incoming message.
type RecipientMessage struct {
	// From is the phone number of the user who sent the message.
	From      string `json:"from"`
	// ID is the unique message identifier.
	ID        string `json:"id"`
	// Timestamp is the time the message was sent by the user.
	Timestamp string `json:"timestamp"`
	// Type defines the kind of message (e.g., "text", "image").
	Type      string `json:"type"`
}

// TextRecipientMessage represents an incoming message that contains text content.
type TextRecipientMessage struct {
	RecipientMessage
	// Text contains the actual body of the message.
	Text TextContent `json:"text"`
}

// TextContent contains the body of a text-based message.
type TextContent struct {
	// Body is the text content of the received message.
	Body string `json:"body"`
}