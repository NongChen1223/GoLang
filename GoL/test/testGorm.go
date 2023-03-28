package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:xww19981223@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	//Create
	// db.Create(&Product{Code: "D42", Price: 100})
	user := &models.UserBasic{}
	user.Name = "NongChen"
	db.Create(user)
	//Read
	fmt.Println("打印", db.First(user, 1)) //根据整型主键查找
	//db.First(user, "code = ?", "D42") //查找code为D42的商品

	//Update - 更新product的price为200
	db.Model(user).Update("PassWord", "1234")
	//Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Update(map[string]interface{}{"Price": 200, "Code": "F42"})

	//delete - 删除product
	// db.Delete(&product, 1)
}
