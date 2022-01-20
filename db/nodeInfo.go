package Sql

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn string = "root:stvsl2060@tcp(127.0.0.1:3306)/CSystemServer?charset=utf8mb4&parseTime=True&loc=Local"

type NodeInformation struct {
	IP         string    `gorm:"column:IP" json:"ip"`                  // IP地址
	ID         string    `gorm:"column:ID" json:"id"`                  // 节点ID信息
	LOCATE     string    `gorm:"column:LOCATE" json:"locate"`          // 节点地理位置
	TYPE       int       `gorm:"column:TYPE" json:"type"`              // 节点类型
	BELONG     string    `gorm:"column:BELONG" json:"belong"`          // 所属信息
	PRINCIPAL  string    `gorm:"column:PRINCIPAL" json:"principal"`    // 负责人信息
	INSTALLER  string    `gorm:"column:INSTALLER" json:"installer"`    // 安装人信息
	MAINTAINER string    `gorm:"column:MAINTAINER" json:"maintainer"`  // 维护人信息
	DATACONFIG string    `gorm:"column:DATA_CONFIG" json:"dataconfig"` // 数据来源配置
	RSAPRIVATE string    `gorm:"column:RSA_PRIVATE" json:"rsapublic"`  // RSA私钥
	RSAPUBLIC  string    `gorm:"column:RSA_PUBLIC" json:"rsaprivate"`  // RSA公钥
	SELFINFO   string    `gorm:"column:SELF_INFO" json:"selfinfo"`     // 自检信息
	LASTUPLOAD time.Time `gorm:"column:LAST_UPLOAD" json:"lastupload"` // 上次更新日期
	SELFDATE   time.Time `gorm:"column:SELF_DATE" json:"selfdate"`     // 上次自检日期
	REMARK     string    `gorm:"column:REMARK" json:"remark"`          // 备注信息
}

// 建立数据库表
func CreateNodeInformationTable() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 建立数据库表
	db.AutoMigrate(&NodeInformation{})
}

// NodeInformation转换成json
func (NodeInformationList *NodeInformation) ToJson() (string, error) {
	jsonstr, err := json.Marshal(NodeInformationList)
	return string(jsonstr), err
}

// 读取符合地理位置的节点信息并返回对应json格式信息
func ReadNodeInformationByLocate(locate string) string {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 读取符合地理位置的节点信息(模糊搜索)
	var NodeInformationList []NodeInformation
	db.Where("locate LIKE ?", "%"+locate+"%").Find(&NodeInformationList)
	// 将NodeInformation数组转换成json并返回
	jsonstr, err := json.Marshal(NodeInformationList)
	if err != nil {
		panic("转换json格式失败")
	}
	return string(jsonstr)
}

// 读取ID为id的某个节点信息并返回对应json格式信息
func ReadNodeInformationByID(id string) string {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 读取ID为id的某个节点信息
	var NodeInformationList []NodeInformation
	db.Where("id = ?", id).Find(&NodeInformationList)
	// 将NodeInformation数组转换成json并返回
	jsonstr, err := json.Marshal(NodeInformationList)
	if err != nil {
		panic("转换json格式失败")
	}
	return string(jsonstr)
}

// 根据IP查找节点信息并返回对应json格式信息
func ReadNodeInformationByIP(ip string) string {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 读取IP为ip的某个节点信息
	var NodeInformationList []NodeInformation
	db.Where("ip = ?", ip).Find(&NodeInformationList)
	// 将NodeInformation数组转换成json并返回
	jsonstr, err := json.Marshal(NodeInformationList)
	if err != nil {
		panic("转换json格式失败")
	}
	return string(jsonstr)
}

// 新节点信息写入数据库
func InsertNodeInformation(NodeInfo *NodeInformation) string {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 校验节点是否已经存在
	var NodeInformationList []NodeInformation
	db.Where("id = ?", NodeInfo.ID).Find(&NodeInformationList)
	if len(NodeInformationList) == 0 {
		// 插入新节点信息
		db.Create(NodeInfo)
		fmt.Print("新节点信息写入数据库成功")
		return "写入成功"
	} else {
		fmt.Print("新节点信息已存在")
		return "节点已存在"
	}
}

// 更新ID为id的节点信息
func UpdateNodeInformationByID(id string, NodeInformation *NodeInformation) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 更新数据库
	db.Model(&NodeInformation).Where("ID = ?", id).Updates(NodeInformation)
}

// 删除ID为id的节点信息
func DeleteNodeInformationByID(id string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 删除数据库
	db.Where("ID = ?", id).Delete(&NodeInformation{})
}

// 获取指定ID的RSA私钥
func GetRSAKeyByID(id string) string {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 读取ID为id的某个节点RsaPrivate的值
	var NodeInformationList []NodeInformation
	db.Where("ID = ?", id).Find(&NodeInformationList)
	return NodeInformationList[0].RSAPRIVATE
}

// 获取指定ID的RSA公钥
func GetRSAPublicKeyByID(id string) string {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 读取ID为id的某个节点RsaPublic的值
	var NodeInformationList []NodeInformation
	db.Where("ID = ?", id).Find(&NodeInformationList)
	return NodeInformationList[0].RSAPUBLIC
}

// 根据ID返回对象
func GetNodeInformationByID(id string) NodeInformation {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 读取ID为id的某个节点信息
	var NodeInformationList []NodeInformation
	db.Where("ID = ?", id).Find(&NodeInformationList)
	return NodeInformationList[0]
}

// 查询指定类型的节点的ID
func GetNodeIDByType(type_ string) []string {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	// 读取ID为id的某个节点信息
	var NodeInformationList []NodeInformation
	db.Where("type = ?", type_).Find(&NodeInformationList)
	var nodeID []string
	for _, v := range NodeInformationList {
		nodeID = append(nodeID, v.ID)
	}
	return nodeID
}

// 数据库优化
func Optimize() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	db.Exec("optimize table node_information")
}
