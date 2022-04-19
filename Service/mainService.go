package Service

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"stvsljl.com/stvsl/RSA"
	"stvsljl.com/stvsl/Sql"
)

func Start() {
	RSA.GenerateLocalRsaKey()
	Sql.OpenPool()
	gin.SetMode(gin.ReleaseMode) // 发布模式
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.LoadHTMLGlob("./web/pages/*")
	router.GET("/ping", pingHandler)                                              //连接测试,同时返回服务器RSA密钥信息
	router.POST("/ping", pingpostHandler)                                         //二次连接，连理双向通信
	router.GET("/pages/default/map", mapPageHandler)                              //对外提供默认地图页面
	router.GET("/pages/config/map/", mapConfigPageHandler)                        //对外提供配置地图页面
	router.GET("/favicon.ico", faviconHandler)                                    //对外提供favicon.ico
	router.GET("/redpoint", redpointHandler)                                      //对外提供红色坐标点
	router.GET("/greenpoint", greenpointHandler)                                  //对外提供绿色坐标点
	router.GET("/bluepoint", bluepointHandler)                                    //对外提供蓝色坐标点
	router.POST("/submit/nodesubmit", NodeAuthMiddleware(), submitInfoHandler)    //节点数据上传到时序数据库
	router.POST("/auth/nodeauth", nodeauthHandler)                                //节点获取token
	router.POST("/auth/getauth", authHandler)                                     //客户端获取token
	router.GET("/passwd/getpass", AuthMiddleware(), getPassHandler)               //获取用户密码残片
	router.POST("/passwd/getpass", AuthMiddleware(), setPassHandler)              //设置用户密码残片
	router.GET("/node/info", AuthMiddleware(), nodeInfoSendHandler)               //获取相关节点信息(联合时序数据库混合搜索)
	router.GET("/node/professional", AuthMiddleware(), nodeInfoProfessionHandler) //获取相关专业检测信息(单关系数据库)
	router.GET("/node/alldatainfo", AuthMiddleware(), allnodeDatainfoHandler)     //获取相关节点的传感器数据(时间区间内全部)
	router.GET("/node/allpsubinfo", AuthMiddleware(), allnodePsubinfoHandler)     //获取相关节点的专业机构检测数据数据(时间区间内全部)
	router.POST("/node/psubinfo", AuthMiddleware(), nodePsubPostHandler)          //上传节点的专业机构检测数据数据
	router.GET("/standard", AuthMiddleware(), StandardGetHandler)                 // 获取执行标准数据
	router.GET("/node/main", AuthMiddleware(), nodeConfigGetHandler)              //获取节点配置信息
	router.POST("/node/main", AuthMiddleware(), nodeConfigPostHandler)            //获取节点配置信息
	router.DELETE("/node/main", AuthMiddleware(), nodeConfigDeleteHandler)        //删除节点配置信息
	router.POST("/node/set", AuthMiddleware(), nodeConfigSetHandler)              //更新节点配置信息
	router.GET("/my/account", AuthMiddleware(), myAccountGetHandler)              //获取个人账户信息
	router.POST("/my/account", AuthMiddleware(), myAccountPostHandler)            //更新个人账户信息
	router.DELETE("/my/account", AuthMiddleware(), myAccountDeleteHandler)        //冻结个人账户
	router.GET("/account", AuthMiddleware(), accountGetHandler)                   //获取账户信息
	router.POST("/account", AuthMiddleware(), accountPostHandler)                 //更新账户信息
	router.DELETE("/account", AuthMiddleware(), accountDeleteHandler)             //删除账户
	router.POST("/account/new", AuthMiddleware(), accountNewHandler)              //创建新账户
	router.GET("/update", updateHandler)                                          //更新检测
	router.GET("/update/download", updateDownloadHandler)                         //更新下载
	// 启动SSL
	go router.RunTLS(":10214", "rsa/stvsljl.com.crt", "rsa/stvsljl.com.key")
	go router.Run(":10241")
	control := gin.Default()
	control.GET("/control", commandHandler)
	control.Run(":5210")
}

// 命令解释器
func commandHandler(c *gin.Context) {
	// 验证权限
	var id = c.Query("id")
	var passwd = c.Query("passwd")
	if id == "" || passwd == "" {
		c.JSON(200, gin.H{
			"status": "error",
			"msg":    "参数错误",
		})
		return
	}
	a := Sql.AccountInformations{}
	fragment, err := a.GetPasswdFragment(id)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "error",
			"msg":    "鉴权失败",
		})
		return
	}
	// 计算密码
	sha256 := sha256.New()
	sha256.Write([]byte(passwd))
	hash := sha256.Sum(nil)
	passwd = hex.EncodeToString(hash)
	if fragment != passwd {
		c.JSON(200, gin.H{
			"status": "error",
			"msg":    "鉴权失败",
		})
		return
	}
	// 解析命令
	var command = c.Query("command")
	if command == "" {
		c.JSON(200, gin.H{
			"status": "error",
			"msg":    "参数错误",
		})
		return
	}
	// 分割命令
	var commandList = strings.Split(command, " ")
	// 解析命令
	switch commandList[0] {
	case "sql":
		switch commandList[1] {
		case "query":
			// 执行 commandList[2]的查询语句
			var sql = commandList[2]
			var result, err = Sql.Query(sql)
			if err != nil {
				c.JSON(200, gin.H{
					"status": "error",
					"msg":    "执行失败",
				})
				return
			}
			c.JSON(200, gin.H{
				"status": "success",
				"msg":    "执行成功",
				"result": result,
			})
			return
		case "pool":
			switch commandList[3] {
			case "status":
				// 获取连接池状态
				var result = Sql.PoolStatus()
				c.JSON(200, gin.H{
					"result": result,
				})
				return
			case "set":
				// 重新设置连接池大小
				var size = commandList[4]
				// 转换为整数
				var sizeInt, err = strconv.Atoi(size)
				if err != nil {
					c.JSON(200, gin.H{
						"status": "error",
						"msg":    "参数错误",
					})
					return
				}
				// 设置连接池大小
				err = Sql.SetPoolSize(sizeInt)
				if err != nil {
					c.JSON(200, gin.H{
						"status": "error",
						"msg":    "参数错误，只支持扩容且不超过数据库支持上限，如需减小连接池大小需要重启服务器后进行操作",
					})
					return
				}
				c.JSON(200, gin.H{
					"status": "success",
					"msg":    "设置成功",
				})
				return
			default:
				c.JSON(200, gin.H{
					"status": "error",
					"msg":    "参数错误",
				})
			}
		}
	}
}
