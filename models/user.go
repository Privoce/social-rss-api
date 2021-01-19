package models

// UserInfo 用户基础信息
type UserInfo struct {
	NickName  string `json:"nick_name"`
	Account   string `json:"account"`
	AvatarURL string `json:"avatar_url"`
}

func NewUser(name, account, avatar string) *UserInfo {
	return &UserInfo{NickName: name, Account: account, AvatarURL: avatar}
}
