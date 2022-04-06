package Sql

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

/******************************************************************************
 * 节点信息表
 * @Author: stvsl
 * @Date: 2022.1.22
 ****************************************************************************/

// NodeInformations表结构体
type NodeInformations struct {
	IP          string    `gorm:"column:IP" json:"ip"`                    // IP地址
	ID          string    `gorm:"column:ID" json:"id"`                    // 节点ID信息
	LOCATE      string    `gorm:"column:LOCATE" json:"locate"`            // 节点地理位置
	LO          float32   `gorm:"column:LO" json:"lo"`                    // 径度
	LI          float32   `gorm:"column:LI" json:"li"`                    // 纬度
	TYPE        int       `gorm:"column:TYPE" json:"type"`                // 节点类型
	BELONG      string    `gorm:"column:BELONG" json:"belong"`            // 所属机构信息
	PRINCIPAL   string    `gorm:"column:PRINCIPAL" json:"principal"`      // 负责人信息
	INSTALLER   string    `gorm:"column:INSTALLER" json:"installer"`      // 安装人信息
	MAINTAINER  string    `gorm:"column:MAINTAINER" json:"maintainer"`    // 维护人信息
	DATACONFIG  string    `gorm:"column:DATA_CONFIG" json:"dataconfig"`   // 数据来源配置
	AESKEY      string    `gorm:"column:AES_KEY" json:"aeskey"`           // AES密钥
	SELFINFO    string    `gorm:"column:SELF_INFO" json:"selfinfo"`       // 自检信息
	INSTALLDATE time.Time `gorm:"column:INSTALL_DATE" json:"installdate"` // 安装日期
	LASTUPLOAD  time.Time `gorm:"column:LAST_UPLOAD" json:"lastupload"`   // 上次更新日期
	SELFDATE    time.Time `gorm:"column:SELF_DATE" json:"selfdate"`       // 上次自检日期
	REMARK      string    `gorm:"column:REMARK" json:"remark"`            // 备注信息
}

// 实现NodeInfo接口的Insert方法
// 添加新节点
func (nodeinfo NodeInformations) Insert() error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	db.Create(nodeinfo)
	return nil
}

// 实现接口的Update方法
// 根据json信息修改指定的节点
func (nodeinfo NodeInformations) Update() error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	if err != nil {
		log.Panicln("json解析失败：", err)
		return errors.New("json解析失败")
	}
	db.Save(nodeinfo)
	return nil
}

// 实现NodeInfo接口的Delete方法
// 删除指定的节点
func (nodeinfo NodeInformations) Delete(id string) error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	db.Where("ID = ?", id).Delete(nodeinfo)
	return nil
}

// 实现NodeInfo接口的Get方法
// 根据id,地址信息，所属信息，类型聚合查找
func (nodeinfo NodeInformations) GetByBelong(belong string) ([]NodeInformations, error) {
	db, i, err := GetDB()
	if err != nil {
		return nil, errors.New("数据库连接失败")
	}
	defer Release(i)
	// 结果数组
	var result []NodeInformations
	// 在node_informations表中查询BELONG的值为id的所有记录
	db.Where("BELONG = ?", belong).Find(&result)
	return result, nil
}

// 实现NodeInfo接口的ReadByLocate方法
// 读取符合地理位置的节点信息并返回对应json格式信息
func (nodeinfo NodeInformations) GetByLocate(locate string) ([]NodeInformations, error) {
	db, i, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer Release(i)
	// 保存结果集合
	var result []NodeInformations
	// 模糊查询
	db.Where("LOCATE LIKE ?", "%"+locate+"%").Find(&result)
	// 判断是否查询到
	if len(result) == 0 {
		return nil, errors.New("查询不到该节点信息")
	}
	return result, nil
}

// 实现NodeInfo接口的GetIP方法
// 根据ID获取节点的IP地址
func (nodeinfo NodeInformations) GetIP(id string) (string, error) {
	db, i, err := GetDB()
	if err != nil {
		return "", err
	}
	defer Release(i)
	// 保存结果到nodeinfo
	db.Where("ID = ?", id).First(&nodeinfo)
	// 判断是否查询到
	if nodeinfo.ID == "" {
		return "", errors.New("查询不到该节点信息")
	}
	return nodeinfo.IP, nil
}

// 实现NodeInfo接口的GetByID方法
// 根据ID获取节点信息对象
func (nodeinfo NodeInformations) GetByID(id string) (*NodeInformations, error) {
	db, i, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer Release(i)
	// 保存结果到nodeinfo
	db.Where("ID = ?", id).First(&nodeinfo)
	return &nodeinfo, nil
}

// 实现NodeInfo接口的GetByType方法
// 根据类型获取节点信息对象
func (nodeinfo NodeInformations) GetByType(type1 string) (*[]NodeInformations, error) {
	db, i, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer Release(i)
	// 保存结果集合
	var nodeinfolist []NodeInformations
	// 查询数据库
	db.Where("Type = ?", type1).Find(&nodeinfolist)
	return &nodeinfolist, nil
}

// 实现NodeInfo接口的GetCheck方法
// 根据节点id获取其校验信息
func (nodeinfo NodeInformations) GetCheck(id string) (string, error) {
	db, i, err := GetDB()
	if err != nil {
		return "", err
	}
	defer Release(i)
	// 获取数据库中指定id的DataConfig字段数据
	var dataConfig string
	db.Where("ID = ?", id).First(&nodeinfo)
	dataConfig = nodeinfo.DATACONFIG
	// 判断是否查询到
	if dataConfig == "" {
		return "", errors.New("查询不到该节点信息")
	}
	return dataConfig, nil
}

// 实现NodeInfo接口的GetFeatures方法
// 根据ID获取节点的特征信息
func (nodeinfo NodeInformations) GetFeatures(id string) (string, error) {
	// db,i, err := GetDB()
	// if err != nil {
	// 	return "", err
	// }
	// // 保存结果到nodeinfo
	// db.Where("ID = ?", id).First(&nodeinfo)
	// // 判断是否查询到
	// if nodeinfo.ID == "" {
	// 	return "", errors.New("查询不到该节点信息")
	// }
	// str := nodeinfo.RsaPublic
	// // 获取str的SHA256值
	// sha256 := sha256.New()
	// sha256.Write([]byte(str))
	// hash := sha256.Sum(nil)
	// // 转换成16进制
	// hex := hex.EncodeToString(hash)
	return "", nil
}

// 实现NodeInfo接口的GetIDsByBelong方法
// 根据所属信息获取所有匹配的节点的ID
func (nodeinfo NodeInformations) GetIDsByBelong(belong string) ([]string, error) {
	db, i, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer Release(i)
	// 保存查询结果
	var ids []string
	// 查询数据库
	db.Where("Belong = ?", belong).Pluck("ID", &ids)
	return ids, nil
}
func (nodeinfo NodeInformations) VirtualMake() {
	// 创建虚拟节点
	// 创建虚拟节点的数据库
	db, i, err := GetDB()
	if err != nil {
		return
	}
	defer Release(i)
	// 创建虚拟节点的数据库
	// 创建8000个虚拟节点
	tmpp := 0.0
	for i := 0; i < 500; i++ {
		// 创建虚拟节点
		// i 转为7位字符串
		id := strconv.Itoa(i)
		//判断id是否为7位
		for j := len(id); j < 7; j++ {
			id = "0" + id
		}
		// 如果i是5的余数,重新生成i
		if i%5 == 0 {
			tmpp = rand.Float64()
			fmt.Println(tmpp)
		}
		nodeinfo.ID = "CX" + id
		nodeinfo.IP = "127.0.0.1"
		nodeinfo.TYPE = 0
		nodeinfo.LOCATE = "北京"
		nodeinfo.LO = float32(116.404 + tmpp + rand.Float64()/10 - rand.Float64()/10)
		nodeinfo.LI = float32(39.915 + tmpp + rand.Float64()/10 - rand.Float64()/10)
		nodeinfo.BELONG = "CCOM0000001"
		nodeinfo.PRINCIPAL = "张三"
		nodeinfo.INSTALLER = "李四"
		nodeinfo.MAINTAINER = "王五"
		nodeinfo.DATACONFIG = "11111111111111111"
		nodeinfo.AESKEY = "sdffsgsdfg"
		nodeinfo.SELFINFO = "模拟自检"
		nodeinfo.INSTALLDATE = time.Now()
		nodeinfo.LASTUPLOAD = time.Now()
		nodeinfo.SELFDATE = time.Now()
		nodeinfo.REMARK = "模拟数据"
		// 将虚拟节点信息插入数据库
		db.Create(&nodeinfo)
	}
}
