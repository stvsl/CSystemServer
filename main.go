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
		for i := 0; i < 1; i++ {
			test()
		}
		testread()
	}
}

func start() {
	Service.Start()
}

func test() {
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
	sub.NodeId = "CX0000001"
	// 生成随机数
	sub.GasConcentration = rand.Float64()
	sub.Temperature = rand.Float64()
	sub.PH = rand.Float64()
	sub.Density = rand.Float64()
	sub.Conductivity = rand.Float64()
	sub.OxygenConcentration = rand.Float64()
	sub.MetalConcentration = rand.Float64()
	sub.SolidsConcentration = rand.Float64()
	sub.FloatingSolidsConcentration = rand.Float64()
	sub.TotalNitrogen = rand.Float64()
	sub.TotalPhosphorus = rand.Float64()
	sub.TotalOrganicCarbon = rand.Float64()
	sub.BiologicalOxygenDemand = rand.Float64()
	sub.ChemicalOxygenDemand = rand.Float64()
	sub.BacteriaCount = rand.Int()
	sub.StaphylococcusCount = rand.Int()
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
	// 查询一天
	str, err := influxdb.Query(stlist, time.Now().AddDate(0, 0, -1).String(), time.Now().String())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}
