package Service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/******************************************************************************
 * 服务错误码信息
 ******************************************************************************/

//请求获取token时的格式异常
func CX001(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX001",
		"message": "未允许的请求",
	})
}

// 未携带token信息或携带位置异常
func CX002(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    "CX002",
		"message": "请求异常,请确认是否已登陆",
	})
}

// 请求的帐号异常或者第二校验信息（如密码残片或特征值）异常
func CX003(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX003",
		"message": "未允许的请求",
	})
}

//token字段无法解析
func CX006(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": "CX006",
		"msg":  "格式异常",
	})
}

//token无效（过期）
func CX007(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": "CX007",
		"msg":  "无效凭据",
	})
}

//请求的数据无效
func CX301(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX301",
		"message": "数据无效",
	})
}

// 数据解析失败
func CX302(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX302",
		"message": "数据解析失败",
	})
}

//数据库异常
func CX401(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX401",
		"message": "未知错误",
	})
}

// 时序数据库异常
func CX411(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX411",
		"message": "未知错误",
	})
}
