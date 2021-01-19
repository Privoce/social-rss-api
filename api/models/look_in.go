package models

// LookIn 用户在看消息体
type LookIn struct {
	User    *UserInfo `json:"user_info"`
	Time    string    `json:"time"`
	ItemURL string    `json:"URL"`
}
