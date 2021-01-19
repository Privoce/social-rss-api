package repositorys

import "api.privoce.com/rss/models"

var (
	lookInDataBase []*models.LookIn
)

func init() {
	lookInDataBase = append(lookInDataBase, models.NewLookIn("ding", "2021-01-19 18:54:45", "https://github.com/higker"))
	lookInDataBase = append(lookInDataBase, models.NewLookIn("hansu", "2021-01-19 18:54:45", "https://github.com/suhan1996"))
	lookInDataBase = append(lookInDataBase, models.NewLookIn("tom", "2021-01-19 18:54:45", "https://google.com/"))
	lookInDataBase = append(lookInDataBase, models.NewLookIn("ding", "2021-01-19 18:54:45", "https://google.com/"))
	lookInDataBase = append(lookInDataBase, models.NewLookIn("ding", "2021-01-19 18:54:45", "https://google.com/"))
	lookInDataBase = append(lookInDataBase, models.NewLookIn("hansu", "2021-01-19 18:54:45", "https://google.com/"))
	lookInDataBase = append(lookInDataBase, models.NewLookIn("tom", "2021-01-19 18:54:45", "https://google.com/"))
}

// FindLookInsByUserName 根据用户名获取在看消息
func FindLookInsByUserName(name string) []*models.LookIn {
	var lookIns []*models.LookIn
	for _, v := range lookInDataBase {
		if v.NickName == name {
			lookIns = append(lookIns, v)
		}
	}
	return lookIns
}
