package Sql

import (
	"errors"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局dsn
var dsn string = "root:stvsl2060@tcp(127.0.0.1:3306)/CSystemServer?charset=utf8mb4&parseTime=True&loc=Local"

// 开启的数据库连接数
var connsum int = 30

var conns []conn

// 连接优化器
var connnum int = 0

// 数据库连接存储管理器
type conn struct {
	// 数据库连接
	db *gorm.DB
	// 数据库连接锁
	lock bool
}

// 开启数据库连接池
func OpenPool() {
	// coon数组，指定长度
	conns = make([]conn, connsum)
	// 初始化数据库连接
	for i := 0; i < connsum; i++ {
		// 创建数据库连接
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		// 初始化数据库连接锁
		conns[i].lock = false
		conns[i].db = db
	}
}

func Query(sql string) ([]map[string]string, error) {
	db, i, _ := GetDB()
	defer Release(i)
	// 执行语句
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		return nil, err
	}
	// 转换为map
	columns, _ := rows.Columns()
	values := make([]string, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	result := make([]map[string]string, 0)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		entry := make(map[string]string)
		for i, col := range values {
			if col == "" {
				entry[columns[i]] = "NULL"
			} else {
				entry[columns[i]] = string(col)
			}
		}
		result = append(result, entry)
	}
	return result, nil
}

func PoolStatus() string {
	// 统计使用中的数据库连接数
	sum := 0
	for i := 0; i < connsum; i++ {
		if conns[i].lock {
			sum++
		}
	}
	return "当前连接池内连接数：" + string(rune(connsum)) + "，使用中的连接数：" + string(rune(sum))
}

func SetPoolSize(size int) error {
	if size <= connsum {
		return errors.New("连接池大小不能小于已开启的连接数")
	}
	connsum = size
	// 连接池大小变化，扩容处理
	neweconns := make([]conn, connsum)
	// 保存已有连接
	for i := 0; i < connsum; i++ {
		neweconns[i] = conns[i]
	}
	// 初始化新连接
	for i := connsum; i < connsum; i++ {
		// 创建数据库连接
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		// 初始化数据库连接锁
		neweconns[i].lock = false
		neweconns[i].db = db
	}
	// 替换连接池
	conns = neweconns
	return nil
}

// 获取数据库连接器
func GetDB() (*gorm.DB, int, error) {
	// 选择空闲的数据库连接
	for i := connnum; i < connsum; i++ {
		if !conns[i].lock {
			conns[i].lock = true
			connnum++
			return conns[i].db, i, nil
		}
	}
	return nil, -1, errors.New("数据库连接池已满")
}

// 数据库连接释放
func Release(i int) {
	// 释放数据库连接
	conns[i].lock = false
	connnum--
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
