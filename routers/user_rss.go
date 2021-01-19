package routers

import (
	"api.privoce.com/rss/handler"
	"github.com/gin-gonic/gin"
)

// mappingApi mapping social rss api
func mappingApi(r *gin.Engine) {
	r.GET("/:name", handler.RSSList)
	r.POST("/:name", handler.AddUser)

}
