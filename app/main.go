package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	version := "v1"

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		hostname, _ := os.Hostname()
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("This is the host %s, app version %s", hostname, version),
		})
	})
	_ = r.Run()

}
