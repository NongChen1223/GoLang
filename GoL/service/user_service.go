package service

import (
	"fmt"
	"ginchat/models"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	fmt.Printf("GetUserList 路由方法")
	data := make([]*models.User_basic, 10) //创建一个指向User_basic结构体的指针数组 分配10个空间 用于存储查询到的数据
	data = models.GetUserList()            //调用models层的GetUserList方法
	fmt.Println("打印data", data)
	c.JSON(200, gin.H{
		"message": data,
	})
}
