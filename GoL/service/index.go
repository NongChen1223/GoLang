package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	fmt.Printf("GetIndex 路由方法")
	c.JSON(200, gin.H{
		"message": "welcome to ginchat",
	})
}
