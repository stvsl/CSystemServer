package Service

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"stvsljl.com/stvsl/AES"
	"stvsljl.com/stvsl/Sql"
)

func setPassHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	// 信息绑定
	var account Sql.AccountInformations
	account.ID = id
	if err := c.ShouldBindJSON(&account); err != nil {
		CX302(c)
		return
	}
	// 写入
	account.Update()
}

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

func myAccountGetHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	//查询数据库对象
	n := Sql.AccountInformations{}
	// 查询数据库
	str, err := n.Get(id)
	if err != nil {
		CX101(c)
		return
	}
	// 获取token中的aes信息
	aes := c.GetString("aes")
	// AES加密查询结果
	enstr, _ := AES.AesEncrypt([]byte(str), []byte(aes))
	// base64
	str64 := base64.StdEncoding.EncodeToString(enstr)
	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
		"data":    str64,
	})
}

func myAccountPostHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	// 绑定json信息
	var account Sql.AccountInformations
	if err := c.ShouldBindJSON(&account); err != nil {
		CX302(c)
		return
	}
	// 设置id
	account.ID = id
	// 更新数据库
	err := account.Update()
	if err != nil {
		CX303(c)
		return
	}
	// 返回操作结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
	})
}

func myAccountDeleteHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	// 查询数据库对象
	n := Sql.AccountInformations{}
	n.ID = id
	n.STATUS = -1
	err := n.Update()
	if err != nil {
		CX101(c)
		return
	}
}

func accountGetHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	// 查询数据库对象
	n := Sql.AccountInformations{}
	// 校验type信息
	typi, err := n.GetType(id)
	if err != nil {
		CX101(c)
		return
	}
	if typi < 2 {
		CX303(c)
		return
	}
	//获取id信息
	id = c.Query("id")
	// 查询数据库
	str, err := n.Get(id)
	if err != nil {
		CX101(c)
		return
	}
	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
		"data":    str,
	})
}

func accountPostHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	// 绑定json信息
	var account Sql.AccountInformations
	if err := c.ShouldBindJSON(&account); err != nil {
		CX302(c)
		return
	}
	// 校验type信息
	typi, err := account.GetType(id)
	if err != nil {
		CX101(c)
		return
	}
	if typi < 2 {
		CX303(c)
		return
	}
	// 更新
	err = account.Update()
	if err != nil {
		CX101(c)
		return
	}
	// 返回操作结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
	})
}

func accountDeleteHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	// 查询数据库对象
	n := Sql.AccountInformations{}
	// 校验type信息
	typi, err := n.GetType(id)
	if err != nil {
		CX101(c)
		return
	}
	if typi < 2 {
		CX303(c)
		return
	}
	// 获取query中的id信息
	id = c.Query("id")
	// 从数据库中删除此条信息
	err = n.Delete(id)
	if err != nil {
		CX101(c)
		return
	}
	// 返回操作结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
	})
}

func accountNewHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	// 查询数据库对象
	n := Sql.AccountInformations{}
	// 校验type信息
	typi, err := n.GetType(id)
	if err != nil {
		CX101(c)
		return
	}
	if typi < 2 {
		CX303(c)
		return
	}
	// 绑定json信息
	if err := c.ShouldBindJSON(&n); err != nil {
		CX302(c)
		return
	}
	// 插入数据库
	err = n.Add()
	if err != nil {
		CX101(c)
		return
	}
	// 返回操作结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
	})
}
