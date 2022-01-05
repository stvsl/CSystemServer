package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("pages/*")
	router.GET("/pages", func(c *gin.Context) {
		c.HTML(http.StatusOK, "map.html", nil)
	})
	router.GET("/map", func(c *gin.Context) {
		c.HTML(http.StatusOK, "map.html", nil)
	})
	router.Run(":8011")
}
