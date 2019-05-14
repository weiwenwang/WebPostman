package Nav

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/weiwenwang/WebPostman/models"
)

type Nav struct {
	gorm.Model
	Id       int    `gorm:"id"`
	ParentID int    `gorm:"pid"`
	Name     string `gorm:"name"`
	//List     []*Nav `json:"list,omitempty"`
}

func (Nav) TableName() string {
	return "nav"
}

func Info(id int32) (*Nav) {
	var nav Nav
	models.DB.Where("id = ?", id).Find(&nav)
	fmt.Println(nav)
	return &nav
}

func NavList() []map[string]interface{} {
	mp5 := makeNav(5, "api", nil)
	mp6 := makeNav(6, "price", nil)
	var sub3 []map[string]interface{}
	mp3 := makeNav(3, "ggstudy", append(sub3, mp5, mp6))

	mp7 := makeNav(7, "index", nil)
	var sub4 []map[string]interface{}
	mp4 := makeNav(4, "common", append(sub4, mp7))

	var sub1 []map[string]interface{}
	mp1 := makeNav(1, "wochacha", append(sub1, mp3))

	var sub2 []map[string]interface{}
	mp2 := makeNav(2, "yxpt", append(sub2, mp4))

	var mp0 []map[string]interface{}
	mp0 = append(mp0, mp1)
	mp0 = append(mp0, mp2)
	return mp0
}

func makeNav(id int32, name string, sub []map[string]interface{}) map[string]interface{} {
	mp1 := map[string]interface{}{
		"id":   id,
		"name": name,
		"sub":  sub,
	}
	return mp1
}
