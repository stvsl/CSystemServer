package main

import (
	"fmt"
	"math/rand"
	"time"

	"stvsljl.com/stvsl/Service"
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
		// 死循环执行test
		for i := 0; i < 100; i++ {
			test("CX0000001")
			test("CX0000002")
		}
		testread()
	}
}

func start() {
	Service.Start()
}

func test(id string) {
	// RSA.GenerateLocalRsaKey()
	// test := "121211"
	// enstr, err := RSA.Encrypt([]byte(test), []byte(RSA.RSA_PUBLIC_LOCAL))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(enstr)
	// destr, err := RSA.Decrypt([]byte(enstr), []byte(RSA.RSA_PRIVATE_LOCAL))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(destr))
	sub := influxdb.SubmitInfo{}
	// 使用时间作为随机数种子
	rand.Seed(time.Now().UnixNano())
	// 随机生成一个节点提交信息
	sub.NodeId = id
	// 生成随机数
	sub.GasConcentration = rand.Float64()
	sub.Temperature = rand.Float64()
	sub.PH = rand.Float64()
	sub.Density = rand.Float64()
	sub.Conductivity = rand.Float64()
	sub.OxygenConcentration = rand.Float64()
	sub.MetalConcentration = rand.Float64()
	sub.SC = rand.Float64()
	sub.FSC = rand.Float64()
	sub.TN = rand.Float64()
	sub.TP = rand.Float64()
	sub.TOC = rand.Float64()
	sub.BOD = rand.Float64()
	sub.COD = rand.Float64()
	sub.BC = (int64)(rand.Float64() * 100)
	sub.SLC = (int64)(rand.Float64() * 100)
	err := influxdb.Write(&sub)
	if err != nil {
		fmt.Println(err, "!")
	}
	// 等待上面的写入操作完成
	time.Sleep(time.Second * 1)
}

func testread() {
	var stlist []string
	stlist = append(stlist, "CX0000001")
	stlist = append(stlist, "CX0000002")
	// 查询一天
	str, err := influxdb.Query(stlist, time.Now().AddDate(0, 0, -1).String(), time.Now().String())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}
