// Author: SDing <deen.job@qq.com>
// Date: 2021/1/19 - 7:47 下午 - UTC/GMT+08:00

package handler

import (
	"api.privoce.com/rss/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	us = new(services.UserService)
	ss = new(services.SocialService)
)

func RSSList(c *gin.Context) {
	name := c.Param("name")
	userInfo := us.GetUserByName(name)
	if userInfo == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "user info not find.",
		})
		return
	}
	ins := ss.GetUserLookIns(userInfo.NickName)
	if len(ins) == 0 || ins == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":      400,
			"msg":       "user look in no data.",
			"user_info": userInfo,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "successful.",
		"look_in":   ins,
		"user_info": userInfo,
	})
	return
}
