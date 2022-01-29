package Sql

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

// 实现AccountInfo接口的Get方法
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

// 帐号存在性校验
func (accountInformations AccountInformations) Exist(id string) (bool, error) {
	db, err := GetDB()
	if err != nil {
		return false, err
	}
	var account AccountInformations
	db.Where("ID = ?", id).Find(&account)
	return account.ID != "", err
}
