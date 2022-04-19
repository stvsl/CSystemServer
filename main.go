package main

import (
	"C"
	"fmt"
	"math/rand"
	"time"

	"stvsljl.com/stvsl/Service"
	"stvsljl.com/stvsl/influxdb"
)
import (
	"strconv"

	"stvsljl.com/stvsl/Sql"
)

func main() {
	fmt.Println("运行模式选择：\n 0.服务器模式（默认） 1.测试模式 ")
	status := 0
	// 从键盘输入status
	fmt.Scanf("%d", &status)
	if status == 0 {
		start()
	} else if status == 1 {
		Sql.OpenPool()
		var NI = Sql.NodeInformations{}
		NI.VirtualMake()
		var pro = Sql.Professiondata{}
		// 死循环执行test
		for i := 0; i < 20; i++ {
			for i := 1; i < 201; i++ {
				// i 转为7位字符串
				id := strconv.Itoa(i)
				//判断id是否为7位
				for j := len(id); j < 7; j++ {
					id = "0" + id
				}
				id = "CX" + id
				test(id)
				pro.VirtualMake(id)
			}
			testread()
		}
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
	sub.GasConcentration = 10 + rand.Float64()*10
	sub.Temperature = 15 + rand.Float64()*10
	sub.PH = rand.Float64()*3 + 4
	sub.Density = rand.Float64()
	sub.Conductivity = rand.Float64()
	sub.OxygenConcentration = rand.Float64() * 0.9
	sub.MetalConcentration = rand.Float64() * 2
	sub.SC = rand.Float64() * 100
	sub.FSC = rand.Float64() * 80
	sub.TN = rand.Float64() * 50
	sub.TP = rand.Float64() * 1.5
	sub.TOC = rand.Float64() * 15
	sub.BOD = rand.Float64() * 10
	sub.COD = rand.Float64() * 50
	sub.BC = (int64)(rand.Float64() * 100)
	sub.SLC = (int64)(rand.Float64() * 100)
	err := influxdb.Write(&sub)
	if err != nil {
		fmt.Println(err, "!")
	}
	// 等待上面的写入操作完成
}

func testread() {
	var stlist []string
	stlist = append(stlist, "CX0000001")
	stlist = append(stlist, "CX0000002")
	// 查询一天
	str, err := influxdb.Query(stlist, "-24h", "-0h")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}
