package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

func main() {
	db, err := gorm.Open("mysql", "root:@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 创建表 自动迁移
	db.AutoMigrate(&UserInfo{})
	// 创建数据
	//u1 := UserInfo{1,"张三","男","打游戏"}
	//db.Create(u1)

	// 查询数据
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n",u)

	// 更新
	db.Model(&u).Update("hobby","学习")

	// 删除
	db.Delete(&u)
}
