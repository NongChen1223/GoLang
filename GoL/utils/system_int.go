package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败", err)
	}
	fmt.Println("读取配置文件 config app:", viper.Get("app"))
	fmt.Println("读取配置文件 config mysql:", viper.Get("mysql"))

}
func InitMySQL() {
	DB, _ := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{})
	fmt.Println("MySQL连接成功", DB)
	//user := models.User_basic{}
	//DB.Find(&user)
	//fmt.Println("打印user", user)

}
