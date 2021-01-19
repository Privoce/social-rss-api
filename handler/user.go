// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2021/1/19 - 8:42 下午 - UTC/GMT+08:00

package handler

import (
	"api.privoce.com/rss/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(c *gin.Context) {
	var userinfo models.UserInfo
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
}
