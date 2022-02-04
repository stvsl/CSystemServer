package Sql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局dsn
var dsn string = "root:stvsl2060@tcp(127.0.0.1:3306)/CSystemServer?charset=utf8mb4&parseTime=True&loc=Local"

// 获取数据库连接器
func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("数据库连接失败：", err)
		return nil, err
	}
	return db, nil
}

// 数据库优化
func Optimize() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	db.Exec("optimize table node_information")
	db.Exec("optimize table account_informations")
}
