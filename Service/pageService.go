package Service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func mapPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "map.html", nil)
	fmt.Println("IP:", c.ClientIP())
}

func faviconHandler(c *gin.Context) {
	c.File("favicon.png")
}

func redpointHandler(c *gin.Context) {
	c.File("redpoint.png")
}

func bluepointHandler(c *gin.Context) {
	c.File("favicon.png")
}

func greenpointHandler(c *gin.Context) {
	c.File("greenpoint.png")
}
