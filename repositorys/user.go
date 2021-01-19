package repositorys

import "api.privoce.com/rss/models"

var (
	userDataBase []*models.UserInfo
)

func init() {
	userDataBase = append(userDataBase, models.NewUser("ding", "ding@privoce.com", "https://iconfont.alicdn.com/t/1570850261949.png@100h_100w.jpg"))
	userDataBase = append(userDataBase, models.NewUser("hansu", "hansu@privoce.com", "https://iconfont.alicdn.com/t/1557889264039.JPG@100h_100w.jpg"))
	userDataBase = append(userDataBase, models.NewUser("tom", "tom@privoce.com", "https://iconfont.alicdn.com/t/1557889264039.JPG@100h_100w.jpg"))
}

// FindOneUserByName 根据name查询用户
func FindOneUserByName(name string) *models.UserInfo {
	for _, u := range userDataBase {
		if u.NickName == name {
			return u
		}
	}
	return nil
}
