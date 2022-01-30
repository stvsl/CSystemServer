package Service

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"stvsljl.com/stvsl/Sql"
	"stvsljl.com/stvsl/influxdb"
)

// stvsljl.com/stvsl/node/info?Locate=北京&Belong&Type=0
func nodeInfoAbsoluteHandler(c *gin.Context) {
	// 获取query参数
	Locate := c.Query("Locate")
	Belong := c.Query("Belong")
	Type := c.Query("Type")
	// 获取token中的id信息
	id := c.GetString("id")
	//查询数据库对象
	n := Sql.NodeInformation{}
	// 查询数据库
	str, err := n.GetAbsolute(id, Locate, Belong, Type)
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

func nodeDatainfoHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	// 从query中获取开始时间和结束时间
	startTime := c.Query("StartTime")
	endTime := c.Query("EndTime")
	// 转换为influxdb时间
	// 在account中查询其下的所有节点id
	// 查询数据库对象
	n := Sql.AccountInformations{}
	// 查询其机构代码
	s, err := n.GetOrganization(id)
	if err != nil {
		CX101(c)
		return
	}
	// 查询数据库
	var ids []string
	n2 := Sql.NodeInformation{}
	ids, err = n2.GetIDsByBelong(s)
	if err != nil {
		CX101(c)
		return
	}
	//查询时序数据库
	str, err := influxdb.Query(ids, startTime, endTime)
	if err != nil {
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

func nodePsubinfoHandler(c *gin.Context) {
	// 从token获取id
	id := c.GetString("id")
	// 从query中获取开始时间和结束时间
	startTime := c.Query("StartTime")
	endTime := c.Query("EndTime")
	// 查询数据库
	p := Sql.Professiondata{}
	// 查询数据库对象
	n := Sql.AccountInformations{}
	// 查询其机构代码
	s, err := n.GetOrganization(id)
	if err != nil {
		CX101(c)
		return
	}
	// 查询数据库
	var ids []string
	n2 := Sql.NodeInformation{}
	ids, err = n2.GetIDsByBelong(s)
	if err != nil {
		CX101(c)
		return
	}
	// 查询数据库
	data, x, err := p.GetByTime(ids, startTime, endTime)
	if err != nil {
		CX302(c)
		return
	}
	// 转换为json
	str, _ := json.Marshal(data)
	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
		"data":    string(str),
		"count":   x,
	})
}

func nodePsubPostHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	// 检查id权限
	a := Sql.AccountInformations{}
	typei, err := a.GetType(id)
	if err != nil {
		CX101(c)
		return
	}
	if typei != 1 {
		CX303(c)
		return
	}
	// 获取json数据
	var data Sql.Professiondata
	err = c.BindJSON(&data)
	if err != nil {
		CX301(c)
		return
	}
	// 写入数据库
	err = data.Add()
	if err != nil {
		CX301(c)
		return
	}
	// 返回操作结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
	})
}

func nodeConfigGetHandler(c *gin.Context) {
	// 获取query参数
	id := c.Query("nodeid")
	// 查询数据库对象
	n := Sql.NodeInformation{}
	// 查询数据库
	node, err := n.GetByID(id)
	if err != nil {
		// 判断err内容是否是数据库连接失败
		if strings.Contains(err.Error(), "数据库连接失败") {
			CX101(c)
			return
		}

		CX301(c)
		return
	}
	// node转换为json
	str, _ := json.Marshal(node)
	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
		"config":  str,
	})
}

func nodeConfigPostHandler(c *gin.Context) {
	obj := Sql.NodeInformation{}
	// 绑定json数据
	err := c.BindJSON(&obj)
	if err != nil {
		CX302(c)
		return
	}
	// 更新数据库
	err = obj.Insert()
	if err != nil {
		CX401(c)
		return
	}
	// 返回操作结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
	})
}

func nodeConfigDeleteHandler(c *gin.Context) {
	obj := Sql.NodeInformation{}
	err := obj.Delete(c.Query("nodeid"))
	if err != nil {
		CX401(c)
		return
	}
	// 返回操作结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
	})
}

func nodeConfigSetHandler(c *gin.Context) {
	n := Sql.NodeInformation{}
	// 绑定json数据
	err := c.BindJSON(&n)
	if err != nil {
		CX302(c)
		return
	}
	// 更新数据库
	err = n.Update()
	if err != nil {
		CX401(c)
		return
	}
	// 返回操作结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
	})
}
