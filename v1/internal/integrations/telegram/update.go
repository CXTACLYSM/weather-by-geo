package telegram

type Update struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}

type Message struct {
	MessageID int       `json:"message_id"`
	From      *User     `json:"from,omitempty"`
	Chat      *Chat     `json:"chat"`
	Text      string    `json:"text,omitempty"`
	Location  *Location `json:"location,omitempty"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username,omitempty"`
}

type Chat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
