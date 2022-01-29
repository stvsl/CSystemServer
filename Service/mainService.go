package Service

import (
	"github.com/gin-gonic/gin"
	"stvsljl.com/stvsl/RSA"
)

func Start() {
	RSA.GenerateLocalRsaKey()
	// gin.SetMode(gin.ReleaseMode) // 发布模式
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.LoadHTMLGlob("./web/pages/*")
	router.GET("/pages/default/map", mapPageHandler)                           //对外提供地图页面
	router.POST("/submit/nodesubmit", NodeAuthMiddleware(), submitInfoHandler) //节点数据上传到时序数据库
	router.POST("/auth/nodeauth", nodeauthHandler)                             //节点获取token
	router.POST("/auth/getauth", authHandler)                                  //客户端获取token
	router.GET("/passwd/getpass", NodeAuthMiddleware(), getPassHandler)        //获取密码残片
	router.GET("/node/info", NodeAuthMiddleware(), nodeInfoSendHandler)        //获取相关节点信息
	router.Run(":10214")
}
