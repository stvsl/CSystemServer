package main

import (
	// "fmt"
	// "net/http"

	"fmt"
	"influxdb"
	"time"
	// "github.com/gin-gonic/gin"
)

func main() {
	//连接时序数据库
	influxdb.Write()
	// influxdb.Write2()
	influxdb.Query()
	time.Sleep(time.Second * 10)
	fmt.Println("!")
	// router := gin.Default()
	// router.LoadHTMLGlob("./pages/*")
	// router.GET("/pages", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "map.html", nil)
	// 	fmt.Println("IP:", c.ClientIP())
	// })
	// router.Run(":8011")
}
