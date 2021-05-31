package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init(){
	db, _ = gorm.Open("mysql", "root:@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
}

// 1.定义模型
type User struct {
	gorm.Model
	Name string
	Age int64
}

// https://www.liwenzhou.com/posts/Go/gorm_crud/
func main() {
	defer db.Close()
	// 2.关系对应
	db.AutoMigrate(&User{})

	// 3.插入数据
	//u1 := User{Name: "张三",Age: 18}
	//u2 := User{Name: "李四",Age: 20}
	//db.Create(&u1)
	//db.Create(&u2)

	// 4.查询
	query()
}

// 一般查询
func query() {
	var user User
	db.Debug().First(&user) // 查询第一条数据
	fmt.Printf("user:%#v\n",user)

	var users []User
	db.Debug().Find(&users) // 查询所有未删除数据
	fmt.Printf("user:%#v\n",users)
}

// 条件查询
func query2() {
	var user User
	var users []User
	db.Where("name <> ?", "jinzhu").Find(&users)
	// IN
	db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&users)
	// AND
	db.Where("name = ? AND age >= ?","张三",10).Find(&users)

	// 使用 struct 或 map 进行查询
	// Struct
	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;

	// Map
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 主键的切片
	db.Where([]int64{20, 21, 22}).Find(&users)
	//// SELECT * FROM users WHERE id IN (20, 21, 22);
}