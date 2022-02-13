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
		MD51    string `json:"KEY"`
		FEATURE string `json:"FEATURE"`
		RSA     string `json:"RSA"`
	}
	c.BindJSON(&j)
	// base64解密
	demd51, _ := base64.StdEncoding.DecodeString(j.MD51)
	defeature, _ := base64.StdEncoding.DecodeString(j.FEATURE)
	// 解密数据
	md51byte, err := RSA.Decrypt(demd51, []byte(RSA.RSA_PRIVATE_LOCAL))
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
	md52 := string(result)
	md5str := string(md51byte) + md52
	// 加密并返回字符串
	retmd52, _ := RSA.Encrypt([]byte(md52), []byte(j.RSA))
	c.JSON(200, gin.H{
		"code":    "CX200",
		"message": "success",
		"KEY":     retmd52,
	})
	// 计算md5
	md5_md5 := md5.Sum([]byte(md5str))
	// 保存到数据库
	var l Sql.Logining
	l.TIME = time.Now()
	l.FEATURE = string(featurebyte)
	l.MD5KEY = fmt.Sprintf("%x", md5_md5)
	l.Write()
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":       "CX200",
		"message":    "pong",
		"RSA_PUBLIC": RSA.GetPublicKey(),
	})
}
