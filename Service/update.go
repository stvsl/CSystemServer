package Service

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 版本信息结构体
type version struct {
	Version string `json:"version"` // 版本号
	Date    string `json:"date"`    //更新日期
	Desc    string `json:"desc"`    //更新描述
	AES     string `json:"AES"`     //更新包AES
}

var latestVersion = version{
	Version: "0.0.1",
	Date:    "2022.2.3",
	Desc:    "测试版本",
	AES:     "0000000000",
}

func updateHandler(c *gin.Context) {
	// 返回latestVersion
	c.JSON(http.StatusOK, latestVersion)
}

func updateDownloadHandler(c *gin.Context) {
	// 获取query中type参数
	typee := c.Query("type")
	version := c.Query("version")
	// 拼接文件名
	fileName := "CSystem_" + version
	if typee == "Linux" {
		fileName += "_Linux.tar.gz"
	} else if typee == "Windows" {
		fileName += "_Windows.zip"
	}
	// 构建文件路径
	filePath := "package/" + fileName
	// 文件byte[] 流
	fileByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    "CX404",
			"message": "文件不存在" + filePath,
		})
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "CSystem/appdownload")
	c.Header("Accept-Length", fmt.Sprintf("%d", len(fileByte)))
	c.Writer.Write(fileByte)
}
