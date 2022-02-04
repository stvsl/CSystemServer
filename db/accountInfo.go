package Sql

import "encoding/json"

/******************************************************************************
 * 账户信息表
 * @Author stvsl
 * @Date 2022.1.24
 ****************************************************************************/

type AccountInformations struct {
	ID           string `gorm:"column:ID" json:"id"`                     // 账户ID
	TYPE         int    `gorm:"column:TYPE" json:"tYPE"`                 // 账户类型
	PASSWD       string `gorm:"column:PASSWD" json:"passwd"`             // 账户密码
	FRAGMENT     string `gorm:"column:FRAGMENT" json:"fragment"`         // 密码残片
	USERNAME     string `gorm:"column:USER_NAME" json:"uSERNAME"`        // 用户姓名
	USERID       string `gorm:"column:USER_ID" json:"uSERId"`            // 身份证号
	USERLOCATE   string `gorm:"column:USER_LOCATE" json:"uSERLOCATE"`    // 家庭住址
	ORGANIZATION string `gorm:"column:ORGANIZATION" json:"oRGANIZATION"` // 所属机构代码
	RSAPRIVATE   string `gorm:"column:RSA_PRIVATE" json:"rSAPRIVATE"`    // RSA私钥
	RSAPUBLIC    string `gorm:"column:RSA_PUBLIC" json:"rSAPUBLIC"`      // PSA公钥
	STATUS       int    `gorm:"column:STATUS" json:"sTATUS"`             // 账户状态
}

// 根据ID获取账户信息
func (accountInformations AccountInformations) Get(id string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	var account AccountInformations
	db.Where("ID = ?", id).Find(&account)
	// 转换为JSON
	accountJSON, err := json.Marshal(account)
	if err != nil {
		return "", err
	}
	return string(accountJSON), nil
}

// 实现AccountInfo接口的GetType方法
// 根据ID获取节点的类型
func (accountInformations AccountInformations) GetType(id string) (int, error) {
	db, err := GetDB()
	if err != nil {
		return -1, err
	}
	var account AccountInformations
	db.Where("ID = ?", id).Find(&account)
	return account.TYPE, nil
}

// 实现AccountInfo接口的GetByType方法
// 根据账户类型获取符合的账户并返回
func (accountInformations AccountInformations) GetByType(accountType int) (*[]AccountInformations, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	var accounts []AccountInformations
	db.Where("TYPE = ?", accountType).Find(&accounts)
	return &accounts, nil
}

// 实现NodeInfo接口的Add方法
// 添加账户信息
func (accountInformations AccountInformations) Add() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	db.Create(&accountInformations)
	return nil
}

// 实现AccountInfo接口的Delete方法
// 根据ID删除账户信息
func (accountInformations AccountInformations) Delete(id string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	db.Where("ID = ?", id).Delete(&accountInformations)
	return nil
}

// 实现AccountInfo接口的Update方法
// 修改指定ID的账户信息
func (accountInformations AccountInformations) Update() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	db.Save(&accountInformations)
	return nil
}

// 实现AccountInfo接口的GetFragment方法
// 查询指定ID的用户的密码残片
func (accountInformations AccountInformations) GetFragment(id string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	var account AccountInformations
	db.Where("ID = ?", id).Find(&account)
	return account.FRAGMENT, nil
}

// 实现AccountInfo接口的GetOrganization方法
// 查询指定ID的用户的所属机构信息
func (accountInformations AccountInformations) GetOrganization(id string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 查询其Organization信息
	var account AccountInformations
	db.Where("ID = ?", id).Find(&account)
	return account.ORGANIZATION, nil
}

// 实现AccountInfo接口的Exist方法
// 判断指定ID的账户是否存在
func (accountInformations AccountInformations) Exist(id string) (bool, error) {
	db, err := GetDB()
	if err != nil {
		return false, err
	}
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
