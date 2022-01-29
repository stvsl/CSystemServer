package Service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"stvsljl.com/stvsl/Sql"
)

func getPassHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	//查询数据库对象
	n := Sql.AccountInformations{}
	// 查询数据库
	str, err := n.GetFragment(id)
	if err != nil {
		CX101(c)
		return
	}
	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
		"feature": str,
	})
}
