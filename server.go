// Copyright (c) 2020 Privoce Project
// Author: SDing <deen.job@qq.com>
// Date: 2020/8/23 - 9:10 PM - UTC/GMT+08:00

package main

import (
	"api.privoce.com/rss/configs"
	"api.privoce.com/rss/routers"
	"log"
)

func main() {
	log.Println("Server running... Port:", configs.Port)
	routers.Start()
}
