package main

import (
	"fmt"

	"stvsljl.com/stvsl/RSA"
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

func test() {
	RSA.GenerateLocalRsaKey()
	test := "121211"
	enstr, err := RSA.Encrypt([]byte(test), []byte(RSA.RSA_PUBLIC_LOCAL))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(enstr)
	destr, err := RSA.Decrypt([]byte(enstr), []byte(RSA.RSA_PRIVATE_LOCAL))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(destr))
}
