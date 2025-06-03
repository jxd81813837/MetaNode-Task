package controller

import (
	"MetaNode-Task/task-go/go_base/task-4-blog/models"
	"MetaNode-Task/task-go/go_base/task-4-blog/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建评论 只有认证的才可以创建
func CreateCommentControl(c *gin.Context) {
	var com models.Comment
	if err := c.ShouldBindJSON(&com); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	com.AuthorId = uint(c.GetFloat64("userID"))
	err := repository.CreateCommon(com)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error评论失败": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "评论成功",
		"code":    "sucess",
	})
}
