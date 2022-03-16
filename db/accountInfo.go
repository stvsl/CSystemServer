package Sql

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
)

/******************************************************************************
 * 账户信息表
 * @Author stvsl
 * @Date 2022.1.24
 ****************************************************************************/

type AccountInformations struct {
	ID           string `gorm:"column:ID" json:"id"`                     // 账户ID
	TYPE         int    `gorm:"column:TYPE" json:"type"`                 // 账户类型
	PASSWD       string `gorm:"column:PASSWD" json:"passwd"`             // 账户密码
	FRAGMENT     string `gorm:"column:FRAGMENT" json:"fragment"`         // 密码残片
	USERNAME     string `gorm:"column:USER_NAME" json:"username"`        // 用户姓名
	USERID       string `gorm:"column:USER_ID" json:"userid"`            // 身份证号
	USERLOCATE   string `gorm:"column:USER_LOCATE" json:"userlocate"`    // 家庭住址
	TEL          string `gorm:"column:TEL" json:"tel"`                   // 电话号码
	ORGANIZATION string `gorm:"column:ORGANIZATION" json:"organization"` // 所属机构代码
	AES          string `gorm:"column:AES" json:"aes"`                   // AES密钥
	STATUS       int    `gorm:"column:STATUS" json:"status"`             // 账户状态
}

// 根据ID获取账户信息
func (accountInformations AccountInformations) Get(id string) (string, error) {
	db, i, err := GetDB()
	if err != nil {
		return "", err
	}
	db.Where("ID = ?", id).Find(&accountInformations)
	defer Release(i)
	// 转换为JSON
	accountJSON, err := json.Marshal(accountInformations)
	if err != nil {
		return "", err
	}
	return string(accountJSON), nil
}

// 实现AccountInfo接口的GetType方法
// 根据ID获取节点的类型
func (accountInformations AccountInformations) GetType(id string) (int, error) {
	db, i, err := GetDB()
	if err != nil {
		return -1, err
	}
	defer Release(i)
	var account AccountInformations
	db.Where("ID = ?", id).Find(&account)
	return account.TYPE, nil
}

// 实现AccountInfo接口的GetByType方法
// 根据账户类型获取符合的账户并返回
func (accountInformations AccountInformations) GetByType(accountType int) (*[]AccountInformations, error) {
	db, i, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer Release(i)
	var accounts []AccountInformations
	db.Where("TYPE = ?", accountType).Find(&accounts)
	return &accounts, nil
}

// 根据ID判断是否是最高管理员账户
func (accountInformations AccountInformations) IsHAdmin(id string) (bool, string, string, error) {
	db, i, err := GetDB()
	if err != nil {
		return false, "", "", errors.New("数据库连接失败")
	}
	defer Release(i)
	// 查询结果
	var result AccountInformations
	// 查询指定ID的节点的TYPE值是否为3
	db.Where("ID = ?", id).Find(&result)
	if result.TYPE == 3 {
		return true, result.ORGANIZATION, result.AES, nil
	}
	return false, result.ORGANIZATION, result.AES, nil
}

// 实现NodeInfo接口的Add方法
// 添加账户信息
func (accountInformations AccountInformations) Add() error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	db.Create(&accountInformations)
	return nil
}

// 实现AccountInfo接口的Delete方法
// 根据ID删除账户信息
func (accountInformations AccountInformations) Delete(id string) error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	db.Where("ID = ?", id).Delete(&accountInformations)
	return nil
}

// 实现AccountInfo接口的UpdateAES方法
// 更新指定ID的账户AES信息
func (accountInformations AccountInformations) UpdateAES(id string, aes string) error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	db.Model(&accountInformations).Where("ID = ?", id).Update("AES", aes)
	return nil
}

// 实现AccountInfo接口的Update方法
// 修改指定ID的账户信息
func (accountInformations AccountInformations) Update() error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	db.Save(&accountInformations)
	return nil
}

// 实现AccountInfo接口的GetPasswdFragment方法
// 根据ID获取账户密码特征信息
func (accountInformations AccountInformations) GetPasswdFragment(id string) (string, error) {
	db, i, err := GetDB()
	if err != nil {
		return "", err
	}
	defer Release(i)
	var account AccountInformations
	db.Where("ID = ?", id).Find(&account)
	str := account.PASSWD
	// 获取str的SHA256值
	sha256 := sha256.New()
	sha256.Write([]byte(str))
	hash := sha256.Sum(nil)
	// 转换成16进制
	hex := hex.EncodeToString(hash)
	return hex, nil
}

// 实现AccountInfo接口的GetFragment方法
// 查询指定ID的用户的密码残片
func (accountInformations AccountInformations) GetFragment(id string) (string, error) {
	db, i, err := GetDB()
	if err != nil {
		return "", err
	}
	defer Release(i)
	var account AccountInformations
	db.Where("ID = ?", id).Find(&account)
	return account.FRAGMENT, nil
}

// 实现AccountInfo接口的GetOrganization方法
// 查询指定ID的用户的所属机构信息
func (accountInformations AccountInformations) GetOrganization(id string) (string, error) {
	db, i, err := GetDB()
	if err != nil {
		return "", err
	}
	defer Release(i)
	// 查询其Organization信息
	var account AccountInformations
	db.Where("ID = ?", id).Find(&account)
	return account.ORGANIZATION, nil
}

// 实现AccountInfo接口的Exist方法
// 判断指定ID的账户是否存在
func (accountInformations AccountInformations) Exist(id string) (bool, error) {
	db, i, err := GetDB()
	if err != nil {
		return false, err
	}
	defer Release(i)
	var account AccountInformations
	db.Where("ID = ?", id).Find(&account)
	result := account.ID != ""
	// 查询status信息是否为-1的冻结状态
	if result {
		db.Where("ID = ?", id).Find(&account)
		result = account.STATUS != -1
	}
	return result, err
}
