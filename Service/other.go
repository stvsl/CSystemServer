package Service

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"stvsljl.com/stvsl/RSA"
	"stvsljl.com/stvsl/Sql"
)

func pingpostHandler(c *gin.Context) {
	// 解析json数据
	var j struct {
		AES1    string `json:"KEY"`
		FEATURE string `json:"FEATURE"`
		RSA     string `json:"RSA"`
	}
	c.BindJSON(&j)
	// base64解密
	deAES1, _ := base64.StdEncoding.DecodeString(j.AES1)
	defeature, _ := base64.StdEncoding.DecodeString(j.FEATURE)
	// 解密数据
	AES1byte, err := RSA.Decrypt(deAES1, []byte(RSA.RSA_PRIVATE_LOCAL))
	if err != nil {
		CX304(c)
		return
	}
	featurebyte, err := RSA.Decrypt(defeature, []byte(RSA.RSA_PRIVATE_LOCAL))
	if err != nil {
		CX304(c)
		return
	}
	// 生成随机字符串
	bytes := []byte(RSA.RSA_PUBLIC_LOCAL)
	var result []byte
	for i := 0; i < 15; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	AES2 := string(result)
	AESstr := string(AES1byte) + AES2
	// 加密并返回字符串
	retAES2, _ := RSA.Encrypt([]byte(AES2), []byte(j.RSA))
	c.JSON(200, gin.H{
		"code":    "CX200",
		"message": "success",
		"KEY":     retAES2,
	})
	// 计算AES
	AES_AES := md5.Sum([]byte(AESstr))
	// 保存到数据库
	var l Sql.Logining
	l.TIME = time.Now()
	l.FEATURE = string(featurebyte)
	l.AESKEY = fmt.Sprintf("%x", AES_AES)
	l.Write()
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":       "CX200",
		"message":    "pong",
		"RSA_PUBLIC": RSA.GetPublicKey(),
	})
}
