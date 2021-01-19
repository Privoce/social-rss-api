// Author: SDing <deen.job@qq.com>
// Date: 2021/1/19 - 6:19 下午 - UTC/GMT+08:00

package services

import (
	"api.privoce.com/rss/models"
	"api.privoce.com/rss/repositorys"
)

type UserService struct{}

func (s UserService) GetUserByName(name string) *models.UserInfo {
	if len(name) == 0 {
		return nil
	}
	return repositorys.FindOneUserByName(name)
}
