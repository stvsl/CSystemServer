package main

import (
	"fmt"

	"stvsljl.com/stvsl/Service"
	"stvsljl.com/stvsl/Sql"
	"stvsljl.com/stvsl/influxdb"
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

func test() {
	var x Sql.NodeInformation
	fmt.Println(x.GetIP("C0000000001"))
	err := influxdb.ConnectTest()
	if err != nil {
		fmt.Println(err)
	}
}
