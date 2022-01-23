package Service

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.LoadHTMLGlob("./web/pages/*")
	router.GET("/pages/default/map", mapPageHandler)     //对外提供地图页面
	router.POST("/submit/nodesubmit", submitInfoHandler) //节点数据上传到时序数据库
	router.POST("/auth/getauth", authHandler)            //获取token
	router.Run(":10214")
}
