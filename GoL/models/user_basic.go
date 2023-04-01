package models

import (
	"fmt"
	"ginchat/utils"
	"gorm.io/gorm"
	"time"
)

type User_basic struct {
	gorm.Model
	Name         string
	PassWord     string
	Phone        string
	Email        string
	Identity     string
	ClientIP     string    //Ip地址
	ClienetPort  string    //端口地址
	LoginTime    time.Time //登录时间
	HearBeatTime time.Time //心跳时间
	LoginOutTime time.Time //登出时间 `gorm:"column:login_out_time;comment:'登出时间'" json:"login_out_time"`
	IsOnline     bool      //是否在线
	DeviceType   string    //设备类型
}

/*
table *UserBasic 表示一个指针类型的参数，
指向一个 UserBasic 结构体类型的变量。
在函数中，我们可以使用这个指针来访问和修改
UserBasic 类型变量的字段和方法。
使用table.xxx
*/
func (table *User_basic) TableName() string {
	return "user_basic"
}

func GetUserList() []*User_basic {
	data := make([]*User_basic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println("打印v", v)
	}
	return data
}
