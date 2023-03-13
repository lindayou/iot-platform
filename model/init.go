package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:root@tcp(192.168.179.142:3306)/iot-platform?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("[DB ERROR] :", err)
	}
	err = db.AutoMigrate(&UserBasic{}, &ProductBasic{}, &DeviceBasic{})
	if err != nil {
		log.Fatalln("[DB Error] :", err)
		return
	}
	DB = db
}
