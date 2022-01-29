package Service

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"stvsljl.com/stvsl/Sql"
)

// stvsljl.com/stvsl/node/info?Locate=北京&Belong&Type=0
func nodeInfoSendHandler(c *gin.Context) {
	// 获取query参数
	Locate := c.Query("Locate")
	Belong := c.Query("Belong")
	Type := c.Query("Type")
	// 获取token中的id信息
	id := c.GetString("id")
	//查询数据库对象
	n := Sql.NodeInformation{}
	// 查询数据库
	str, err := n.Get(id, Locate, Belong, Type)
	if err != nil {
		// 判断err内容是否是数据库连接失败
		if strings.Contains(err.Error(), "数据库连接失败") {
			CX101(c)
			return
		}

		CX301(c)
		return
	}
	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
		"data":    str,
	})
}
