package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init(){
	db, _ = gorm.Open("mysql", "root:@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
}

// 1.定义模型
type User struct {
	ID int64
	Name string `gorm:"default:'匿名'"` // 指定默认值
	// 允许使用字符串零值
	//Name *string `gorm:"default:'匿名'"`
	//Name sql.NullString `gorm:"default:'匿名'"` //
	Age int64
}

func main() {
	defer db.Close()
	// 2.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	// 3.创建记录
	u := User{Name: "张三",Age: 18}
	println(db.NewRecord(&u)) // 判断主键是否为空
	db.Create(&u)
	println(db.NewRecord(&u))
}
