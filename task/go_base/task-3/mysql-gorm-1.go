package task

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

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

func main_grom_1() {
	db := InitDbgGorm()

	//查询某个用户所有文章以及评论
	user := GetPostAndCommentByUser(db, "jjc")
	fmt.Println(user)
	//获取评论最多的文章
	fmt.Println(GetCommentMax(db))
	post := Post{}
	post.Data = "新增内容1111"
	post.Title = "新增内容标题"
	post.AuthorId = 1
	//创建文章时候更新文章数量统计字段
	//CreatePostTx(db, post)
	//删除评论 如果文章评论0 更新文章状态为"无评论"
	//DeleteComment(db, 3)

	//var posts []Post
	//db.Model(&Post{}).Find(&posts)
	//fmt.Println(posts)
}

// 获取所有的文章
func AllPostList(db *gorm.DB) ([]Post, error) {
	var posts = []Post{}
	err := db.Model(&Post{}).Find(&posts).Error
	return posts, err
}

func DeleteComment(db *gorm.DB, commentId uint) {
	var comment Comment
	db.Where("id = ?", commentId).First(&comment)

	db.Transaction(func(tx *gorm.DB) error {
		tx.Delete(&Comment{}, commentId)
		if err := AfterDelete(tx, comment.PostId); err != nil {
			return err
		}
		return nil
	})
}

// 删除文章
func DeletePost(db *gorm.DB, postId uint) error {
	err := db.Unscoped().Delete(&Post{}, postId).Error
	return err
}

// 删除评论的钩子函数 更新文章状态
func AfterDelete(tx *gorm.DB, postId uint) error {
	var count int64
	tx.Unscoped().Model(&Comment{}).Where("post_id = ?", postId).Count(&count)
	if count == 0 {
		err := tx.Model(&Post{}).Where("id = ?", postId).UpdateColumn("comment_state", "无评论").Error
		if err != nil {
			return err
		}
	}
	return nil
}

// 更新文章文章内容标题等信息
func UpdatePost(db *gorm.DB, post Post) error {
	return db.Model(&Post{}).
		Where("id = ?", post.Id).
		Select("Title", "Data").
		Updates(post).Error
}

// 判断这个文章是不是作者的
// 是作者的返回 true 不是返回false
func CheckPostOwnerShip(db *gorm.DB, post Post) bool {
	err := db.Model(&Post{}).
		Where("id =? AND author_id =?", post.Id, post.AuthorId).
		First(&post).Error
	if err != nil {
		return false
	}
	return true
}

// 创建文章 公共方法
func CreatePostTx(db *gorm.DB, post Post) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&post).Error; err != nil {
			return err
		}
		if err := AfterCreate(tx, post.AuthorId); err != nil {
			return err
		}
		return nil
	})
	return err
}

// 更新文章数量 -钩子函数
func AfterCreate(tx *gorm.DB, userId uint) error {
	// 使用 GORM 的 Update 方法对 PostCount +1
	err := tx.Model(&User{}).
		Where("id = ?", userId).
		UpdateColumn("post_count", gorm.Expr("post_count + 1")).
		Error
	if err != nil {
		return err
	}
	return nil
}

// 依据文章id获取评论数量
func GetCommentCount(db *gorm.DB, postId uint) int64 {
	var count int64
	db.Model(&Comment{}).Where("post_id = ?", postId).Count(&count)
	return count
}

// 获取某一篇文章所有评论
func GetCommentById(db *gorm.DB, postId uint) []Comment {
	var comments []Comment
	err := db.Model(&Comment{}).Where("post_id = ?", postId).Find(&comments).Error
	if err != nil {
		return nil
	}
	return comments
}

// 创建评论
func CreateCommon(db *gorm.DB, com Comment) error {
	if err := db.Create(&com).Error; err != nil {
		return err
	}
	return nil
}

// 通过用户 获取所有文章以及评论
func GetPostAndCommentByUser(db *gorm.DB, name string) User {
	user := User{}
	posts := []Post{}
	db.Where("name = ?", name).Find(&user)
	db.Where("author_id = ?", user.Id).Find(&posts)
	user.Posts = posts

	for i := range posts {
		var comments []Comment
		db.Where("post_id = ?", posts[i].Id).Find(&comments)
		posts[i].Comments = comments
	}
	return user
}

// 获取评论最多的文章
func GetCommentMax(db *gorm.DB) (Post, error) {
	var comment Comment
	var post Post
	err := db.Raw("select post_id,count(post_id) as count from ajxd_comment group by post_id order by count desc limit 1").Scan(&comment).Error
	if err != nil {
		return post, err
	}
	err = db.Where("id = ?", comment.PostId).Find(&post).Error

	return post, nil
}

func InitDbgGorm() *gorm.DB {
	dsn := "admin:test@tcp(·127.0.0.1:3306)/dsc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 获取通用数据库对象 sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get generic database object")
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(10)           // 空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(100)          // 数据库打开的最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间
	return db
}
