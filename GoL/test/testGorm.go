package main

import (
	"fmt"
	"ginchat/models"
)

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

func main() {
	//迁移 schema
	DB.AutoMigrate(&models.User_basic{})
	//Create
	// db.Create(&Product{Code: "D42", Price: 100})
	user := &models.User_basic{}
	user.Name = "NongChen"
	DB.Create(user)
	//Read
	a := DB.First(user, "id = ?", "2") //查找code为D42的商品
	fmt.Printf("打印a", a)               //根据整型主键查找
	//Update - 更新user的PassWord为200
	DB.Model(user).Update("PassWord", "1234")
	//Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Update(map[string]interface{}{"Price": 200, "Code": "F42"})

	//delete - 删除product
	// db.Delete(&product, 1)
	// 输出hello

}
