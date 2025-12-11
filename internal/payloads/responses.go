package payloads

type BaseResponse struct {
	Object  string  `json:"object"`
	Entries []Entry `json:"entry"`
}

type Entry struct {
	Id      string   `json:"id"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Field string `json:"field"`
	Value any    `json:"value"`
}

type Receipient struct {
	Profile            Profile `json:"profile"`
	WhatsappBusinessId string  `json:"wa_id"`
}

type Profile struct {
	Name string `json:"name"`
}

type ReceipientMessage struct {
	From      string `json:"from"`
	Id        string `json:"id"`
	TimeStamp string `json:"timestamp"`
	Type      string `json:"type"`
}

type TextReceipientMessage struct {
	ReceipientMessage
	Text TextContent `json:"text"`
}

type TextContent struct {
	Body string `json:"body"`
}
