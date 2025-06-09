package models

import "time"

// 表名自定义
func (User) TableName() string {
	return "ajxd_user"
}
func (Post) TableName() string {
	return "ajxd_post"
}

func (Comment) TableName() string {
	return "ajxd_comment"
}

type User struct {
	Id         uint      `gorm:"primarykey; comment:主键"`
	Name       string    `gorm:"type:varchar(255); comment:姓名"`
	UserName   string    `gorm:"type:varchar(255); comment:账号"`
	Password   string    `gorm:"type:varchar(255); comment:密码"`
	CreateTime time.Time `gorm:"type:datetime; autoCreateTime; comment:创建时间"`
	UpdateTime time.Time `gorm:"type:datetime; autoUpdateTime; comment:更新时间"`
	Posts      []Post    `gorm:"foreignKey:AuthorId;references:Id;-"`
	PostCount  uint      `gorm:"comment:文章统计数量; "`
}
type Post struct {
	Id       uint `gorm:"primarykey; comment:主键"`
	AuthorId uint `gorm:"comment:作者id"`
	//Title*表示 可以传入null 进去 不加的话默认空字符串
	Title        string    `gorm:"type:varchar(255); comment:标题" `
	Data         string    `gorm:"type:text; comment:文章" `
	CommentState string    `gorm:"type:varchar(8); comment:状态"`
	CreateTime   time.Time `gorm:"type:datetime; autoCreateTime; comment:创建时间"`
	UpdateTime   time.Time `gorm:"type:datetime; autoUpdateTime; comment:更新时间"`
	Comments     []Comment `gorm:"foreignkey:PostId评论;-"`
}
type Comment struct {
	Id         uint      `gorm:"primarykey; comment:主键"`
	AuthorId   uint      `gorm:"comment:作者Id"`
	PostId     uint      `gorm:"comment:文章Id"`
	Remake     string    `gorm:"type:varchar(256); comment:评论"`
	CreateTime time.Time `gorm:"type:datetime; autoCreateTime; comment:创建时间"`
	UpdateTime time.Time `gorm:"type:datetime; autoUpdateTime; comment:更新时间"`
}
