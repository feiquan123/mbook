package models

type BookCategory struct {
	ID         int `json:"id" orm:"column(id);"`
	BookID     int `json:"book_id" orm:"column(book_id);"`         // 图书ID
	CategoryID int `json:"category_id" orm:"column(category_id);"` // 分类ID
}

func (b BookCategory) TableName() string {
	return TableNameBookCategory
}
