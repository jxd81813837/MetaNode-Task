package repository

import (
	"MetaNode-Task/task-go/go_base/task-4-blog/config"
	"MetaNode-Task/task-go/go_base/task-4-blog/models"
	"gorm.io/gorm"
)

// 创建评论
func CreateCommon(com models.Comment) error {
	if err := config.Db.Create(&com).Error; err != nil {
		return err
	}
	return nil
}

// 获取某一篇文章所有评论
func GetCommentById(postId uint) []models.Comment {
	var comments []models.Comment
	err := config.Db.Model(&models.Comment{}).Where("post_id = ?", postId).Find(&comments).Error
	if err != nil {
		return nil
	}
	return comments
}

// 依据文章id获取评论数量
func GetCommentCount(postId uint) int64 {
	var count int64
	config.Db.Model(&models.Comment{}).Where("post_id = ?", postId).Count(&count)
	return count
}

// 删除评论的钩子函数 更新文章状态
func AfterDelete(tx *gorm.DB, postId uint) error {
	var count int64
	tx.Unscoped().Model(&models.Comment{}).Where("post_id = ?", postId).Count(&count)
	if count == 0 {
		err := tx.Model(&models.Post{}).Where("id = ?", postId).UpdateColumn("comment_state", "无评论").Error
		if err != nil {
			return err
		}
	}
	return nil
}

// 获取评论最多的文章
func GetCommentMax() (models.Post, error) {
	var comment models.Comment
	var post models.Post
	err := config.Db.Raw("select post_id,count(post_id) as count from ajxd_comment group by post_id order by count desc limit 1").Scan(&comment).Error
	if err != nil {
		return post, err
	}
	err = config.Db.Where("id = ?", comment.PostId).Find(&post).Error

	return post, nil
}
