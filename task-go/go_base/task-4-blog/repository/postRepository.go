package repository

import (
	"MetaNode-Task/task-go/go_base/task-4-blog/config"
	"MetaNode-Task/task-go/go_base/task-4-blog/models"
	"gorm.io/gorm"
)

// 获取所有的文章
func AllPostList() ([]models.Post, error) {
	var posts = []models.Post{}
	err := config.Db.Model(&models.Post{}).Find(&posts).Error
	return posts, err
}

// 删除文章依据id
func DeleteComment(commentId uint) {
	var comment models.Comment
	config.Db.Where("id = ?", commentId).First(&comment)

	config.Db.Transaction(func(tx *gorm.DB) error {
		tx.Delete(&models.Comment{}, commentId)
		if err := AfterDelete(tx, comment.PostId); err != nil {
			return err
		}
		return nil
	})
}

// 删除文章
func DeletePost(postId uint) error {
	err := config.Db.Unscoped().Delete(&models.Post{}, postId).Error
	return err
}

// 更新文章文章内容标题等信息
func UpdatePost(post models.Post) error {
	return config.Db.Model(&models.Post{}).
		Where("id = ?", post.Id).
		Select("Title", "Data").
		Updates(post).Error
}

// 判断这个文章是不是作者的
// 是作者的返回 true 不是返回false
func CheckPostOwnerShip(post models.Post) bool {
	err := config.Db.Model(&models.Post{}).
		Where("id =? AND author_id =?", post.Id, post.AuthorId).
		First(&post).Error
	if err != nil {
		return false
	}
	return true
}

// 创建文章 公共方法
func CreatePostTx(post models.Post) error {
	err := config.Db.Transaction(func(tx *gorm.DB) error {
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
	err := tx.Model(&models.User{}).
		Where("id = ?", userId).
		UpdateColumn("post_count", gorm.Expr("post_count + 1")).
		Error
	if err != nil {
		return err
	}
	return nil
}

// 通过用户 获取所有文章以及评论
func GetPostAndCommentByUser(name string) models.User {
	user := models.User{}
	posts := []models.Post{}
	config.Db.Where("name = ?", name).Find(&user)
	config.Db.Where("author_id = ?", user.Id).Find(&posts)
	user.Posts = posts

	for i := range posts {
		var comments []models.Comment
		config.Db.Where("post_id = ?", posts[i].Id).Find(&comments)
		posts[i].Comments = comments
	}
	return user
}
