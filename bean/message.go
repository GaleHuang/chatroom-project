package bean

type Message struct {
	Time       int64  `json:"time"`
	SenderName string `json:"sender_name"`
	Content    string `json:"content"`
}
