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
	router.GET("/pages/default/map", mapPageHandler)                                 //对外提供地图页面
	router.POST("/submit/nodesubmit", NodeAuthMiddleware(), submitInfoHandler)       //节点数据上传到时序数据库
	router.POST("/auth/nodeauth", nodeauthHandler)                                   //节点获取token
	router.POST("/auth/getauth", authHandler)                                        //客户端获取token
	router.GET("/passwd/getpass", NodeAuthMiddleware(), getPassHandler)              //获取用户密码残片
	router.GET("/node/info", NodeAuthMiddleware(), nodeInfoSendHandler)              //获取相关节点信息(模糊搜索)
	router.GET("/node/info/absolute", NodeAuthMiddleware(), nodeInfoAbsoluteHandler) //获取相关节点信息(严格匹配)
	router.GET("/node/datainfo", NodeAuthMiddleware(), nodeDatainfoHandler)          //获取相关节点的传感器数据
	router.GET("/node/psubinfo", NodeAuthMiddleware(), nodePsubinfoHandler)          //获取相关节点的专业机构检测数据数据
	router.POST("/node/psubinfo", NodeAuthMiddleware(), nodePsubPostHandler)         //上传节点的专业机构检测数据数据
	router.GET("/node/main", NodeAuthMiddleware(), nodeConfigGetHandler)             //获取节点配置信息
	router.POST("/node/main", NodeAuthMiddleware(), nodeConfigPostHandler)           //获取节点配置信息
	router.DELETE("/node/main", NodeAuthMiddleware(), nodeConfigDeleteHandler)       //删除节点配置信息
	router.POST("/node/set", NodeAuthMiddleware(), nodeConfigSetHandler)             //更新节点配置信息
	router.GET("/my/account", NodeAuthMiddleware(), myAccountGetHandler)             //获取个人账户信息
	router.POST("/my/account", NodeAuthMiddleware(), myAccountPostHandler)           //更新个人账户信息
	router.DELETE("/my/account", NodeAuthMiddleware(), myAccountDeleteHandler)       //冻结个人账户
	router.GET("/account", NodeAuthMiddleware(), accountGetHandler)                  //获取账户信息
	router.POST("/account", NodeAuthMiddleware(), accountPostHandler)                //更新账户信息
	router.DELETE("/account", NodeAuthMiddleware(), accountDeleteHandler)            //删除账户
	router.POST("/account/new", NodeAuthMiddleware(), accountNewHandler)             //创建新账户
	router.GET("/update", updateHandler)                                             //更新检测
	router.GET("/update/download", updateDownloadHandler)                            //更新下载
	router.Run(":10214")
}
