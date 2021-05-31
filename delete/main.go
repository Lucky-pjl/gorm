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
}

func main() {
	defer db.Close()

	var u User
	u.ID = 1
	db.Debug().Delete(&u) // 默认是软删除
}
