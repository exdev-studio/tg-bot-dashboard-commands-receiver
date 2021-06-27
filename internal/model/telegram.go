package telegram

type Chat struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

type Message struct {
	Id   int64  `json:"message_id"`
	Text string `json:"text"`
	Date int    `json:"date"`
	Chat Chat   `json:"chat"`
}

type Update struct {
	Id      int64   `json:"update_id"`
	Message Message `json:"message"`
}
