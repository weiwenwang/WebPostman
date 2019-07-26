package Nav

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/weiwenwang/WebRedis/models"
)

type Nav struct {
	gorm.Model
	ParentID int    `gorm:"column:pid"`
	Name     string `gorm:"column:name"`
	//List     []*Nav `json:"list,omitempty"`
}

func (Nav) TableName() string {
	return "nav"
}

type Nav_obj struct {
	Id   uint
	Name string
	Sub  []Nav_obj
}

func NavList() []Nav_obj {
	var ret []Nav_obj
	pid0 := Info(0)
	fmt.Println(pid0)
	for _, v := range pid0 {
		sub := doSub(v)
		for _, sub_v := range Info(v.ID) {
			sub_sub := doSub(sub_v)
			for _, sub_sub_v := range Info(sub_v.ID) {
				third := doSub(sub_sub_v)
				sub_sub.Sub = append(sub_sub.Sub, third)
			}
			sub.Sub = append(sub.Sub, sub_sub)
		}
		ret = append(ret, sub)
	}
	fmt.Println(ret)
	return ret
}

func doSub(nav Nav) Nav_obj {
	var sub Nav_obj
	sub.Id = nav.ID
	sub.Name = nav.Name
	return sub
}

func makeNav(id int32, name string, sub []map[string]interface{}) map[string]interface{} {
	mp1 := map[string]interface{}{
		"id":   id,
		"name": name,
		"sub":  sub,
	}
	return mp1
}

func Info(pid uint) ([]Nav) {
	var navs []Nav
	models.DB.Find(&navs, "pid = ?", pid)
	return navs
}
