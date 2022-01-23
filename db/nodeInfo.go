// @Title 数据库
// @Description 实现了所有的数据库组件功能
// @Author stvsl
// @Date 2022.1.23
package Sql

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/* 定义接口 */
type NodeInfo interface {
	Insert() // 插入节点信息
	// Update()            // 更新ID为id的节点信息
	// ToJson()            // 将对象转换为json格式
	// ReadByLocate()      // 读取符合地理位置的节点信息并返回对应json格式信息
	// ReadByID()          // 读取ID为id的某个节点信息并返回对应json格式信息
	// ReadByIP()          // 读取IP为ip的某个节点信息并返回对应json格式信息
	// Delete()            // 删除节点信息
	// DeleteByID()        // 删除ID为id的节点信息
	// GetRSAPublicByID()  // 根据ID获取节点的RSA公钥
	// GetRSAPrivateByID() // 根据ID获取节点的RSA私钥
	// GetObjectByID()     // 根据ID获取节点信息对象
	// GetObjectByType()   // 根据类型获取节点信息对象
}

// NodeInformation表结构体
type NodeInformation struct {
	IP         string    `gorm:"column:IP" json:"ip"`                  // IP地址
	ID         string    `gorm:"column:ID" json:"id"`                  // 节点ID信息
	Locate     string    `gorm:"column:LOCATE" json:"locate"`          // 节点地理位置
	Type       int       `gorm:"column:TYPE" json:"type"`              // 节点类型
	Belong     string    `gorm:"column:BELONG" json:"belong"`          // 所属信息
	Principal  string    `gorm:"column:PRINCIPAL" json:"principal"`    // 负责人信息
	Installer  string    `gorm:"column:INSTALLER" json:"installer"`    // 安装人信息
	Maintainer string    `gorm:"column:MAINTAINER" json:"maintainer"`  // 维护人信息
	DataConfig string    `gorm:"column:DATA_CONFIG" json:"dataconfig"` // 数据来源配置
	RsaPrivate string    `gorm:"column:RSA_PRIVATE" json:"rsaprivate"` // RSA私钥
	RsaPublic  string    `gorm:"column:RSA_PUBLIC" json:"rsapublic"`   // RSA公钥
	SelfInfo   string    `gorm:"column:SELF_INFO" json:"selfinfo"`     // 自检信息
	LastUpload time.Time `gorm:"column:LAST_UPLOAD" json:"lastupload"` // 上次更新日期
	SelfDate   time.Time `gorm:"column:SELF_DATE" json:"selfdate"`     // 上次自检日期
	Remark     string    `gorm:"column:REMARK" json:"remark"`          // 备注信息
}

// 实现nodeInfo接口
func (nodeinfo *NodeInformation) Insert() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	db.Create(nodeinfo)
}
