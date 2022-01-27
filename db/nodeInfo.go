package Sql

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"time"
)

/******************************************************************************
 * 节点信息表
 * @Author: stvsl
 * @Date: 2022.1.22
 ****************************************************************************/

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

// 实现NodeInfo接口的Insert方法
// 添加新节点
func (nodeinfo NodeInformation) Insert() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	db.Create(nodeinfo)
	return nil
}

// 实现接口的Update方法
// 根据json信息修改指定的节点
func (nodeinfo NodeInformation) Update(jsonstr string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	// 根据json信息序列化为对象
	err = json.Unmarshal([]byte(jsonstr), &nodeinfo)
	if err != nil {
		log.Panicln("json解析失败：", err)
		return errors.New("json解析失败")
	}
	db.Save(nodeinfo)
	return nil
}

// 实现NodeInfo接口的Delete方法
// 删除指定的节点
func (nodeinfo NodeInformation) Delete(id string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	db.Where("ID = ?", id).Delete(nodeinfo)
	return nil
}

// 实现NodeInfo接口的ToJson方法
// 根据id查询节点
func (nodeinfo NodeInformation) Get(id string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 查询
	db.Where("ID = ?", id).First(&nodeinfo)
	// 判断是否查询到
	if nodeinfo.ID == "" {
		return "", errors.New("查询不到该节点信息")
	}
	// 转换为json格式
	json, err := json.Marshal(nodeinfo)
	if err != nil {
		log.Panicln("转换json格式失败：", err)
		return "", errors.New("转换json格式失败")
	}
	return string(json), nil
}

// 实现NodeInfo接口的ReadByLocate方法
// 读取符合地理位置的节点信息并返回对应json格式信息
func (nodeinfo NodeInformation) ReadByLocate(locate string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 保存结果集合
	var result []NodeInformation
	// 查询
	db.Where("LOCATE = ?", locate).Find(&result)
	// 判断是否查询到
	if len(result) == 0 {
		return "", errors.New("查询不到该节点信息")
	}
	// 转换为json格式
	json, err := json.Marshal(result)
	if err != nil {
		log.Panicln("转换json格式失败：", err)
		return "", errors.New("转换json格式失败")
	}
	return string(json), nil
}

// 实现NodeInfo接口的ReadByID方法
// 读取符合ID的节点信息并返回对应json格式信息
func (nodeinfo NodeInformation) ReadByID(id string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 保存结果到nodeinfo
	db.Where("ID = ?", id).First(&nodeinfo)
	// 判断是否查询到
	if nodeinfo.ID == "" {
		return "", errors.New("查询不到该节点信息")
	}
	// 转换为json格式
	json, err := json.Marshal(nodeinfo)
	if err != nil {
		log.Panicln("转换json格式失败：", err)
		return "", errors.New("转换json格式失败")
	}
	return string(json), nil
}

// 实现NodeInfo接口的GetIP方法
// 根据ID获取节点的IP地址
func (nodeinfo NodeInformation) GetIP(id string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 保存结果到nodeinfo
	db.Where("ID = ?", id).First(&nodeinfo)
	// 判断是否查询到
	if nodeinfo.ID == "" {
		return "", errors.New("查询不到该节点信息")
	}
	return nodeinfo.IP, nil
}

// 实现NodeInfo接口的GetRSAPublic方法
// 根据ID获取节点的RSA公钥
func (nodeinfo NodeInformation) GetRSAPublic(id string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 保存结果到nodeinfo
	db.Where("ID = ?", id).First(&nodeinfo)
	// 判断是否查询到
	if nodeinfo.ID == "" {
		return "", errors.New("查询不到该节点信息")
	}
	return nodeinfo.RsaPublic, nil
}

// 实现NodeInfo接口的GetRSAPrivate方法
// 根据ID获取节点的RSA私钥
func (nodeinfo NodeInformation) GetRSAPrivate(id string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 保存结果到nodeinfo
	db.Where("ID = ?", id).First(&nodeinfo)
	// 判断是否查询到
	if nodeinfo.ID == "" {
		return "", errors.New("查询不到该节点信息")
	}
	return nodeinfo.RsaPrivate, nil
}

// 实现NodeInfo接口的GetByID方法
// 根据ID获取节点信息对象
func (nodeinfo NodeInformation) GetByID(id string) (*NodeInformation, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	// 保存结果到nodeinfo
	db.Where("ID = ?", id).First(&nodeinfo)
	return &nodeinfo, nil
}

// 实现NodeInfo接口的GetByType方法
// 根据类型获取节点信息对象
func (nodeinfo NodeInformation) GetByType(type1 string) (*[]NodeInformation, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	// 保存结果集合
	var nodeinfolist []NodeInformation
	// 查询数据库
	db.Where("Type = ?", type1).Find(&nodeinfolist)
	return &nodeinfolist, nil
}

// 实现NodeInfo接口的GetCheck方法
// 根据节点id获取其校验信息
func (nodeinfo NodeInformation) GetCheck(id string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 获取数据库中指定id的DataConfig字段数据
	var dataConfig string
	db.Where("ID = ?", id).First(&nodeinfo)
	dataConfig = nodeinfo.DataConfig
	// 判断是否查询到
	if dataConfig == "" {
		return "", errors.New("查询不到该节点信息")
	}
	return dataConfig, nil
}

// 实现NodeInfo接口的GetFeatures方法
// 根据ID获取节点的特征信息
func (nodeinfo NodeInformation) GetFeatures(id string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 保存结果到nodeinfo
	db.Where("ID = ?", id).First(&nodeinfo)
	// 判断是否查询到
	if nodeinfo.ID == "" {
		return "", errors.New("查询不到该节点信息")
	}
	str := nodeinfo.RsaPublic
	// 获取str的SHA256值
	sha256 := sha256.New()
	sha256.Write([]byte(str))
	hash := sha256.Sum(nil)
	// 转换成16进制
	hex := hex.EncodeToString(hash)
	return hex, nil
}
