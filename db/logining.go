package Sql

import (
	"time"
)

// Logining 登录过程通信暂存表
type Logining struct {
	TIME    time.Time `gorm:"column:TIME"`    // 记录时间
	FEATURE string    `gorm:"column:FEATURE"` // 客户端特征
	AESKEY  string    `gorm:"column:AES_KEY"` // 通信密钥
}

// 写入数据库
func (l Logining) Write() error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	return db.Create(l).Error
}

// 读取密钥
func (l Logining) Read(feature string) (string, error) {
	db, i, err := GetDB()
	if err != nil {
		return "", err
	}
	defer Release(i)
	err = db.Where("FEATURE = ?", feature).Find(&l).Error
	if err != nil {
		return "", err
	}
	return l.AESKEY, nil
}

// 删除密钥
func (l Logining) Delete() error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	return db.Where("FEATURE = ?", l.FEATURE).Delete(l).Error
}

// 删除过期2小时的全部记录
func (l Logining) DeleteExpired() error {
	db, i, err := GetDB()
	if err != nil {
		return err
	}
	defer Release(i)
	return db.Where("TIME < ?", time.Now().Add(-2*time.Hour)).Delete(l).Error
}
