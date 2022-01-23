package main

import (
	"fmt"

	"stvsljl.com/stvsl/Service"
)

func main() {
	fmt.Println("运行模式选择：\n 0.服务器模式（默认） 1.测试模式 ")
	status := 0
	// 从键盘输入status
	fmt.Scanf("%d", &status)
	if status == 0 {
		start()
	} else if status == 1 {
		test()
	}
}

func start() {
	Service.Start()
}

// IP         string    `gorm:"column:IP" json:"ip"`                  // IP地址
// ID         string    `gorm:"column:ID" json:"id"`                  // 节点ID信息
// LOCATE     string    `gorm:"column:LOCATE" json:"locate"`          // 节点地理位置
// TYPE       int       `gorm:"column:TYPE" json:"type"`              // 节点类型
// BELONG     string    `gorm:"column:BELONG" json:"belong"`          // 所属信息
// PRINCIPAL  string    `gorm:"column:PRINCIPAL" json:"principal"`    // 负责人信息
// INSTALLER  string    `gorm:"column:INSTALLER" json:"installer"`    // 安装人信息
// MAINTAINER string    `gorm:"column:MAINTAINER" json:"maintainer"`  // 维护人信息
// DATACONFIG string    `gorm:"column:DATA_CONFIG" json:"dataconfig"` // 数据来源配置
// RSAPRIVATE string    `gorm:"column:RSA_PRIVATE" json:"rsapublic"`  // RSA私钥
// RSAPUBLIC  string    `gorm:"column:RSA_PUBLIC" json:"rsaprivate"`  // RSA公钥
// SELFINFO   string    `gorm:"column:SELF_INFO" json:"selfinfo"`     // 自检信息
// LASTUPLOAD time.Time `gorm:"column:LAST_UPLOAD" json:"lastupload"` // 上次更新日期
// SELFDATE   time.Time `gorm:"column:SELF_DATE" json:"selfdate"`     // 上次自检日期
// REMARK     string    `gorm:"column:REMARK" json:"remark"`          // 备注信息
func test() {

}
