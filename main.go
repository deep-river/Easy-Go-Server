package main

import (
	"EasyGo-Server/pkg/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() 
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello go",
		})
	})

	r.Run(fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port))
}