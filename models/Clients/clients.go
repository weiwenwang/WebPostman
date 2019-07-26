package Clients

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/weiwenwang/WebRedis/models"
)

type Cli struct {
	gorm.Model
	Id     int
	Count  string `gorm:"count"`
	Delete string `gorm:"DeletedAt"`
}

func (Cli) TableName() string {
	return "clients"
}

func List() []*Cli {
	var clients []*Cli
	err := models.DB.Find(&clients).Error
	fmt.Println(err)
	fmt.Println(clients)
	return clients
}

func Info() (clients []*Cli, err error) {
	fmt.Println(models.DB.HasTable("clients"))
	err = models.DB.Find(&clients).Error
	return clients, err
}
