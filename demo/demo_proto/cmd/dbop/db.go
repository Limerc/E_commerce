package main

import (
	"fmt"

	"github.com/Limerc/E_commerce/gomall/demo/demo_proto/biz/dal"
	"github.com/Limerc/E_commerce/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/Limerc/E_commerce/gomall/demo/demo_proto/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err!= nil {
		panic(err)
	}
	dal.Init()
	// 数据库操作
	// 添加数据
	mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "john123"})
	
	// 修改数据
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "123456")
	
	// 读取数据: first() 获取单条数据，Find() 获取多条数据
	var row model.User
	mysql.DB.First(&model.User{}).Where("email = ?", "demo@example.com").First(&row)
	fmt.Printf("row = %+v\n", row)

	// 删除数据
	mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})

	// 强制删除:Unscoped()查询被软删的内容
	mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})
}