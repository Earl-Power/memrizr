package model

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   string `json:"done"`
	UserID string `json:"userid"`
	User   *User  `json:"user"`
}
