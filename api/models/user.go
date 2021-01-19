package models

// UserInfo 用户基础信息
type UserInfo struct {
	NickName  string `json:"nick_name"`
	Account   string `json:"account"`
	AvatarURL string `json:"Avatar_URL"`
}
