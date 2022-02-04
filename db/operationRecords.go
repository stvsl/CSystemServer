package Sql

import (
	"encoding/json"
	"errors"
	"log"
	"time"
)

/****************************************************************
 * 操作日志表
 * @Author: stvsl
 * @Date: 2022.1.25
 ****************************************************************/
type OperationRecords struct {
	OperateTime time.Time `gorm:"column:OPERATE_TIME" json:"operatetime"` // 操作日期
	AccountID   string    `gorm:"column:ACCOUNT_ID" json:"accountid"`     // 操作账户ID
	Action      string    `gorm:"column:ACTION" json:"action"`            // 操作
	Result      string    `gorm:"column:RESULT" json:"result"`            // 操作结果
}

// 实现OperationRecordsInterface接口的Insert方法
func (operationRecords *OperationRecords) Insert() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	db.Create(&operationRecords)
	return nil
}

// 实现OperationRecordsInterface接口的GetByID方法
func (operationRecords *OperationRecords) GetByID(accountID string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 保存查询集合
	var operationRecordsList []OperationRecords
	// 查询
	db.Where("ACCOUNT_ID = ?", accountID).Find(&operationRecordsList)
	// 返回查询结果json
	operationRecordsJSON, err := json.Marshal(operationRecordsList)
	if err != nil {
		log.Panicln("json转换失败：", err)
		return "", errors.New("json转换失败")
	}
	return string(operationRecordsJSON), nil
}

// 实现OperationRecordsInterface接口的GetByTime方法
func (operationRecords *OperationRecords) GetByTime(startTime, endTime string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 保存查询集合
	var operationRecordsList []OperationRecords
	// 将string转换为time
	startTimeTime, err := time.Parse("2006-01-02 15:04:05", startTime)
	if err != nil {
		log.Panicln("时间转换失败：", err)
		return "", errors.New("时间转换失败")
	}
	endTimeTime, err := time.Parse("2006-01-02 15:04:05", endTime)
	if err != nil {
		log.Panicln("时间转换失败：", err)
		return "", errors.New("时间转换失败")
	}
	// 查询
	db.Where("OPERATE_TIME >= ? AND OPERATE_TIME <= ?", startTimeTime, endTimeTime).Find(&operationRecordsList)
	// 返回查询结果json
	operationRecordsJSON, err := json.Marshal(operationRecordsList)
	if err != nil {
		log.Panicln("json转换失败：", err)
		return "", errors.New("json转换失败")
	}
	return string(operationRecordsJSON), nil
}

// 实现OperationRecordsInterface接口的GetResult方法
func (operationRecords *OperationRecords) GetResult(accountID string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	// 保存查询集合
	var result []string
	// 查询
	db.Select("RESULT").Where("ACCOUNT_ID = ?", accountID).Find(&result)
	// 返回查询结果json
	resultJSON, err := json.Marshal(result)
	if err != nil {
		log.Panicln("json转换失败：", err)
		return "", errors.New("json转换失败")
	}
	return string(resultJSON), nil
}

// 实现OperationRecordsInterface接口的Delete方法
func (operationRecords *OperationRecords) Delete(outtime string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	// 将string转换为time
	outtimeTime, err := time.Parse("2006-01-02 15:04:05", outtime)
	if err != nil {
		log.Panicln("时间转换失败：", err)
		return errors.New("时间转换失败")
	}
	// 查询
	db.Where("OPERATE_TIME <= ?", outtimeTime).Delete(&operationRecords)
	return nil
}
