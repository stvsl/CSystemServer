package Service

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

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

//
func nodeInfoSendHandler(c *gin.Context) {
	// 获取token中的id信息
	id := c.GetString("id")
	// 从query中获取开始时间和结束时间
	startTime := c.Query("StartTime")
	endTime := c.Query("EndTime")
	//查询数据库对象
	n := Sql.NodeInformation{}
	var nodelist *[]Sql.NodeInformation

	// 查询账户知否是最高管理员														* 逻辑模式判断，最高管理员可以查看所有节点的数据，按照地理位置模糊查找，获取节点列表
	isAdmin, err := n.IsHAdmin(id)
	if err != nil {
		CX101(c)
		return
	}
	if isAdmin {
		// 获取地理位置
		Locate := c.Query("Locate")
		// 查询数据库中Location为Locate的节点
		nodelist, err = n.GetByLocate(Locate)
		if err != nil {
			CX101(c)
			return
		}
	} else {
		// 查询数据库
		nodelist, err = n.Get(id)
		if err != nil {
			// 判断err内容是否是数据库连接失败
			if strings.Contains(err.Error(), "数据库连接失败") {
				CX101(c)
				return
			}
			CX301(c)
			return
		}
	}

	// 获取nodelist中的所有节点id
	var ids []string
	for _, v := range *nodelist {
		ids = append(ids, v.ID)
	}

	// 查询时序数据库相关结点 														* 时序数据库联合查询，获取节点的时序数据
	influxdata, err := influxdb.Query(ids, startTime, endTime)
	if err != nil {
		CX301(c)
		return
	}

	// 获取nodelist中的所有节点的企业ID												* 企业信息联合查询，获取节点的企业信息
	var orgs []string
	for _, v := range *nodelist {
		orgs = append(orgs, v.Belong)
	}

	// 融数据结构体																	* 融合数据结构体，将节点的时序数据和mariadb数据库信息融合
	type datapoints struct {
		// 结点数据
		IP          string    `json:"ip"`            // IP地址
		ID          string    `json:"id"`            // 节点ID信息
		Locate      string    `json:"locate"`        // 节点地理位置
		Type        int       `json:"type"`          // 节点类型
		Belong      string    `json:"belong"`        // 所属信息
		Principal   string    `json:"principal"`     // 负责人信息
		Installer   string    `json:"installer"`     // 安装人信息
		Maintainer  string    `json:"maintainer"`    // 维护人信息
		DataConfig  string    `json:"dataconfig"`    // 数据来源配置
		AESKey      string    `json:"aeskeyfeature"` // AES密钥指纹
		SelfInfo    string    `json:"selfinfo"`      // 自检信息
		LastUpload  time.Time `json:"lastupload"`    // 上次更新日期
		SelfDate    time.Time `json:"selfdate"`      // 上次自检日期
		Remark      string    `json:"remark"`        // 备注信息
		InstallDate time.Time `json:"installdate"`   // 安装日期
		Lo          float32   `json:"lo"`            // 经度
		Li          float32   `json:"li"`            // 纬度
		// 时序数据
		GasConcentration    float64 `json:"GasConcentration"`    // 气体浓度
		Temperature         float64 `json:"Temperature"`         // 温度
		PH                  float64 `json:"PH"`                  // PH
		Density             float64 `json:"Density"`             // 浊度
		Conductivity        float64 `json:"Conductivity"`        // 电导率
		OxygenConcentration float64 `json:"OxygenConcentration"` // 含氧量
		MetalConcentration  float64 `json:"MetalConcentration"`  // 重金属
		SC                  float64 `json:"SC"`                  // 溶解性固体
		FSC                 float64 `json:"FSC"`                 // 悬浮性固体
		TN                  float64 `json:"TN"`                  // 总氮
		TP                  float64 `json:"TP"`                  // 总磷
		TOC                 float64 `json:"TOC"`                 // 总有机碳
		BOD                 float64 `json:"BOD"`                 // 生物需氧量
		COD                 float64 `json:"COD"`                 // 化学需氧量
		BC                  int64   `json:"BC"`                  // 细菌总数
		SLC                 int64   `json:"SLC"`                 // 大肠杆菌数
	}
	// 查询企业数据库信息

	var dataf []datapoints
	// dataf 转换为json
	json, _ := json.Marshal(&dataf)
	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
		"data":    string(json),
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
	// 查询时序数据库
	str, err := influxdb.Query(ids, startTime, endTime)
	if err != nil {
		CX301(c)
		return
	}
	// 查询关系数据库

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
