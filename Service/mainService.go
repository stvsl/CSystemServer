package Service

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.LoadHTMLGlob("./web/pages/*")
	router.GET("/pages/default", mapPage)
	// 接受到节点提交信息后转换成json
	router.POST("/submit/nodesubmit", submitInfo)
	router.Run(":10214")
}
