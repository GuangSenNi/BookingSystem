package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Comments struct {
	Cid         int64
	Fid         int64
	PhoneNumber string
	Message     string
	Ctime       string
}

type Orders struct {
	Oid         int64
	PhoneNumber string
	Price       float64
	State       string
	StartTime   string
	MealCode    string
	Detail      string
	Windows     string
	MealTime    string
}

type FavoriteFood struct {
	PhoneNumber string
	Fid         int64
}

type UserInfos struct {
	Uid         int64 `gorm:"AUTO_INCREMENT" json:"uid"`
	UserName    string
	PhoneNumber string
	Password    string
}

type Foods struct {
	Fid       int64 `gorm:"AUTO_INCREMENT" json:"fid"`
	FoodName  string
	Price     float64
	Picture   string
	Brief     string
	FoodKinds string
	Count     int64
}

func InitDB() {
	db, err = gorm.Open("mysql", "custom:19961028@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local&readTimeout=500ms")
	if err != nil {
		fmt.Printf("InitDB fail,err=%v", err)
	} else {
		fmt.Println("database connect success")
		db.SingularTable(true)
		//db.AutoMigrate(&Demo_user{}) 自动建表
		//fmt.Println("build table success")
	}
}
