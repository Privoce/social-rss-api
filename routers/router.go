package routers

import (
	"api.privoce.com/rss/configs"
	"github.com/gin-gonic/gin"
)

var (
	root *gin.Engine
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	root = gin.New()
	root.Use(gin.Logger(), gin.Recovery())
}

//Start 运行Web
func Start() {
	mappingApi(root)
	_ = root.Run(configs.Port)
}
