package Url

import (
	"github.com/jinzhu/gorm"
	"github.com/weiwenwang/WebRedis/models"
)

type Url struct {
	gorm.Model
	NavId      int    `gorm:"column:nav_id"`
	Url        string `gorm:"url"`
	Request    string `gorm:"request"`
	Parameters string `gorm:"parameters"`
	Response   string `gorm:"response"`
}

func (Url) TableName() string {
	return "url"
}

func Info(nav_id int) Url {
	var url Url
	models.DB.Find(&url, "nav_id = ?", nav_id)
	return url
}
