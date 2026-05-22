package model

// Message 留言板消息
type Message struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
	Likes     int    `json:"likes"`
}
