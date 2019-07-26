package models

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:WOAImama188@/webRedis?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		DB = db
		return db, err
	}

	return nil, err
}
