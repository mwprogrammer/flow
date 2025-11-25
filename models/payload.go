package models

type Payload struct {
	Product            string              `json:"messaging_product"`
	MessageId          int                 `json:"message_id"`
	Status             string              `json:"status"`
	ReceipientType     string              `json:"receipient_type"`
	To                 string              `json:"to"`
	Type               string              `json:"type"`
	PreviewUrl         bool                `json:"preview_url"`
	TextMessage        *TextMessage        `json:"text"`
	InteractiveMessage *InteractiveMessage `json:"interactive"`
	ImageMessage       *ImageMessage       `json:"image"`
	VideoMessage       *VideoMessage       `json:"video"`
	AudioMessage       *AudioMessage       `json:"audio"`
	DocumentMessage    *DocumentMessage    `json:"document"`
	LocationMessage    *LocationMessage    `json:"location"`
}

type TextMessage struct {
	Body string `json:"body"`
}

type InteractiveMessage struct {
	Type   string `json:"type"`
	Header Header `json:"header"`
	Body   Body   `json:"body"`
	Footer Footer `json:"footer"`
	Action Action `json:"action"`
}

type Header struct {
	Text string `json:"text"`
}

type Body struct {
	Text string `json:"text"`
}

type Footer struct {
	Text string `json:"text"`
}

type Action struct {
	Button   string           `json:"button"`
	Buttons  []ButtonMessage  `json:"buttons"`
	Sections []SectionMessage `json:"sections"`
}

type ButtonMessage struct {
	Type  string `json:"type"`
	Reply Button `json:"reply"`
}

type SectionMessage struct {
	Title string       `json:"title"`
	Rows  []SectionRow `json:"rows"`
}

type SectionRow struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ImageMessage struct {
	Id      int    `json:"id"`
	Caption string `json:"caption"`
	Link    string `json:"link"`
}

type VideoMessage struct {
	Id      int    `json:"id"`
	Caption string `json:"caption"`
	Link    string `json:"link"`
}

type AudioMessage struct {
	Id      int    `json:"id"`
	Caption string `json:"caption"`
	Link    string `json:"link"`
}

type DocumentMessage struct {
	Id   int    `json:"id"`
	Name string `json:"filename"`
	Link string `json:"link"`
}

type LocationMessage struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
	Address   string `json:"address"`
}
