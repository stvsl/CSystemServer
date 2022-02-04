package Sql

/****************************************************************
 * 企业机构信息表
 * @Author stvsl
 * @Date 2022.1.25
 ****************************************************************/
type Organization struct {
	ID         string `gorm:"column:ID" json:"id"`                  // 企业ID
	NAME       string `gorm:"column:NAME" json:"nAME"`              // 企业名称
	REMARK     string `gorm:"column:REMARK" json:"rEMARK"`          // 企业备注
	REGISTTIME string `gorm:"column:REGIST_TIME" json:"rEGISTTIME"` // 注册时间
	TYPE       string `gorm:"column:TYPE" json:"tYPE"`              // 企业类型
	STANDARD   string `gorm:"column:STANDARD" json:"sTANDARD"`      // 服从标准
}

// 实现OrganizationInfo接口的Insert方法
// 新增企业信息
func (o Organization) Insert() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	// 写入数据库
	db.Create(o)
	return nil
}

// 实现OrganizationInfo接口的Update方法
// 更新企业信息
func (o Organization) Update() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	// 更新数据库
	db.Save(o)
	return nil
}

// 实现OrganizationInfo接口的Delete方法
// 删除企业信息
func (o Organization) Delete(id string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	// 删除数据库
	db.Delete(o, id)
	return nil
}

// 实现OrganizationInfo接口的GetByID方法
func (o Organization) GetByID(id string) (*Organization, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	// 查询数据库
	db.Where("id = ?", id).Find(o)
	return &o, nil
}

// 实现OrganizationInfo接口的GetByName方法
// 根据企业名称获取企业信息
func (o Organization) GetByName(name string) (*[]Organization, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	// 保存结果集合
	var result []Organization
	// 模糊查询数据库
	db.Where("name like ?", "%"+name+"%").Find(&result)
	return &result, nil
}

// 实现OrganizationInfo接口的GetByType方法
// 根据企业类型获取企业信息
func (o Organization) GetByType(typee string) (*[]Organization, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	// 保存结果集合
	var result []Organization
	// 查询数据库
	db.Where("type = ?", typee).Find(&result)
	return &result, nil
}

// 实现OrganizationInfo接口的GetByStandard方法
// 根据服从标准获取企业信息
func (o Organization) GetByStandard(standard string) (*[]Organization, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	// 保存结果集合
	var result []Organization
	// 查询数据库
	db.Where("standard = ?", standard).Find(&result)
	return &result, nil
}
