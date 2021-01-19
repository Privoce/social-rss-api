package models

// LookIn 用户在看消息体
type LookIn struct {
	NickName string `json:"nick_name"`
	Time     string `json:"time"`
	ItemURL  string `json:"URL"`
}

func NewLookIn(name, time, url string) *LookIn {
	return &LookIn{NickName: name, Time: time, ItemURL: url}
}
