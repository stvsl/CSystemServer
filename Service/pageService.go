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
