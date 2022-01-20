package Service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"stvsljl.com/stvsl/influxdb"
)

func mapPage(c *gin.Context) {
	c.HTML(http.StatusOK, "map.html", nil)
	fmt.Println("IP:", c.ClientIP())
}

func submitInfo(c *gin.Context) {
	var submitInfo influxdb.SubmitInfo
	c.BindJSON(&submitInfo)
	fmt.Println("submitInfo:", submitInfo)
	// 校验节点提交信息，若数据正常写入时序数据库
	var check string = "OK!"
	if submitInfo.Check() {
		go submitInfo.WriteToDB()
	} else {
		check = "error"
	}
	// 向客户端返回校验信息
	c.JSON(http.StatusOK, gin.H{
		"message": check,
	})
}
