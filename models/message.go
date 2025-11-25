package models

type Message struct {
	Id         string `json:"message_id"`
	Type       string `json:"type"`
	Receipient string `json:"from"`
	List       *List  `json:"list_reply"`
	Text       *Text  `json:"text"`
	Location   string `json:"location"`
}

type Receipient struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Text struct {
	Body string `json:"body"`
}

type Button struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type List struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}
