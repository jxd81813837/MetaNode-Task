package repository

import (
	"MetaNode-Task/task/go_base/task-4-blog/config"
	"MetaNode-Task/task/go_base/task-4-blog/models"
)

// 创建用户
func CreateUser(user models.User) error {
	err := config.Db.Create(&user).Error
	return err
}

// 验证用户信息
func AuthUser(user models.User) (error, models.User) {
	storedUser := models.User{}
	err := config.Db.Where("user_name = ?", user.UserName).First(&storedUser).Error
	return err, storedUser
}
