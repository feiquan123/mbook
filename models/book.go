package models

import "time"

type Book struct {
	BookID         int       `json:"book_id" orm:"column(book_id);pk;auto"`
	BookName       string    `json:"book_name" orm:"column(book_name);size(500)"`                        // 书名
	Identify       string    `json:"identify" orm:"column(identify);size(100);unique"`                   // 唯一标识
	OrderIndex     int       `json:"order_index" orm:"column(order_index);default(0)"`                   // 排序索引
	Description    string    `json:"description" orm:"column(description);size(1024)"`                   // 图书描述
	Cover          string    `json:"cover" orm:"column(cover);size(1024)"`                               // 图书封面
	Editor         string    `json:"editor" orm:"column(editor);size(64)"`                               // 图书编辑器类型： markdown
	Status         int       `json:"status" orm:"column(status);default(0)"`                             // 状态， 0：正常 ， 1：已删除
	PrivatelyOwned int       `json:"privately_owned" orm:"column(privately_owned);default(0)"`           // 是否私有， 0：公开， 1：私有
	PrivateToken   string    `json:"private_token" orm:"column(private_token);size(1024)"`               // 私有图书访问 Token
	MemberID       int       `json:"member_id" orm:"column(member_id);default(0)"`                       // 成员ID
	CreateTime     time.Time `json:"create_time" orm:"column(create_time);type(datetime); auto_now_add"` // 创建时间
	ModifyTime     time.Time `json:"modify_time" orm:"column(modify_time);type(datetime); auto_now_add"` // 修改时间
	ReleaseTime    time.Time `json:"release_time" orm:"column(release_time);type(datetime)"`             // 发布时间
	DocCount       int       `json:"doc_count" orm:"column(doc_count)"`                                  // 章节数量
	CommentCount   int       `json:"comment_count" orm:"column(comment_count);type(int)"`                // 评论数量
	ViewCount      int       `json:"view_count" orm:"column(view_count);default(0)"`                     // 阅读数量
	Start          int       `json:"start" orm:"column(start);default(0)"`                               // 收藏次数
	Score          int       `json:"score" orm:"column(score);default(40)"`                              // 评分
	CntScore       int       `json:"cnt_score" orm:"column(cnt_score);default(0)"`                       // 评分人数
	CntComment     int       `json:"cnt_comment" orm:"column(cnt_comment);default(0)"`                   // 评论人量
	Author         string    `json:"author" orm:"column(author);size(100)"`                              // 作者
	AuthorURL      string    `json:"author_url" orm:"column(author_url);size(1024)"`                     // 作者链接
}

func (b Book) TableName() string {
	return TableNameBook
}
