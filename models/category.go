package models

import "github.com/beego/beego/orm"

type Category struct {
	ID     int    `json:"id" orm:"column(id);"`
	PID    int    `json:"pid" orm:"column(pid);"`                      // 分类ID
	Title  string `json:"title" orm:"column(title);size(100); unique"` // 分类名
	Intro  string `json:"intro" orm:"column(intro);"`                  // 分类介绍
	Icon   string `json:"icon" orm:"column(icon);"`                    // 分类图标
	Cnt    int    `json:"cnt" orm:"column(cnt)"`                       // 分类下的图书
	Sort   int    `json:"sort" orm:"column(sort);"`                    // 排序
	Status bool   `json:"status" orm:"column(status);"`                // 状态： true 显示
}

func (c Category) TableName() string {
	return TableNameCategory
}

func (c Category) GetCategories(pid int, status bool) (categories []Category, err error) {
	qs := orm.NewOrm().QueryTable(c.TableName())
	if pid > -1 {
		qs = qs.Filter("pid = ?", pid)
	}
	qs = qs.Filter("status = ?", status)
	_, err = qs.OrderBy("-status", "sort", "title").All(categories)
	return
}
