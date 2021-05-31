package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init(){
	db, _ = gorm.Open("mysql", "root:@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
}

type User struct {
	gorm.Model
	Name string
	Age int64
	Active bool
}

func main() {
	defer db.Close()

	// 2.关系对应
	db.AutoMigrate(&User{})

	// 3.插入数据
	//u1 := User{Name: "张三",Age: 18,Active: true}
	//u2 := User{Name: "李四",Age: 20,Active: false}
	//db.Create(&u1)
	//db.Create(&u2)

	// 更新
	var user User
	db.First(&user)
	user.Name = "zs"
	db.Debug().Save(user) // db.Save() 默认会更新所有字段

	// 使用Update或Updates更新指定字段
	db.Debug().Model(&user).Update("name","xzs")

	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	//// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;


}
