package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name         string
	PassWord     string
	Phone        string
	Email        string
	Identity     string
	ClientIP     string //Ip地址
	ClienetPort  string //端口地址
	LoginTime    uint64 //登录时间
	HearBeatTime uint64 //心跳时间
	LogOutTime   uint64 //登出时间
	IsOnline     bool   //是否在线
	DeviceType   string //设备类型
}

/*
table *UserBasic 表示一个指针类型的参数，
指向一个 UserBasic 结构体类型的变量。
在函数中，我们可以使用这个指针来访问和修改
UserBasic 类型变量的字段和方法。
使用table.xxx
*/
func (table *UserBasic) TableName() string {
	return "user_basic"
}
