package models

import (
	"github.com/beego/beego/orm"
)

func init() {
	orm.RegisterModel(
		new(Book),
		new(Category),
		new(BookCategory),
	)
}

const (
	TableNameCategory     = "md_category"
	TableNameBook         = "md_book"
	TableNameBookCategory = "md_book_category"
)
