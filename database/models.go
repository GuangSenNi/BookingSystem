package database
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type UserInfos struct {
	Uid         int64 `gorm:"AUTO_INCREMENT" json:"uid"`
	UserName    string
	PhoneNumber string
	Password    string
}

func InitDB() {
	db, err = gorm.Open("mysql", "custom:19961028@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local&readTimeout=500ms")
	if err != nil {
		fmt.Printf("InitDB fail,err=%v", err)
	} else {
		fmt.Println("connect success")
		db.SingularTable(true)
		//db.AutoMigrate(&Demo_user{}) 自动建表
		//fmt.Println("build table success")
	}
	defer db.Close()
}
