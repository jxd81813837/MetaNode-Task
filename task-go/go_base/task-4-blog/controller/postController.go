package controller

import (
	"MetaNode-Task/task-go/go_base/task-4-blog/models"
	"MetaNode-Task/task-go/go_base/task-4-blog/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建文章 只有认证的才可以创建
func CreatePostControl(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//手动校验
	valFlat := checkPostpParam(post)
	if valFlat {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章内容或者 标题不能为空"})
		return
	}
	// 取出 userID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID not found in context"})
		return
	}
	post.AuthorId = uint(userID.(float64))
	//创建文章
	if err := repository.CreatePostTx(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建文章成功",
		"code":    "sucess",
	})
}

// 获取所有文章的列表
func AllPostListControl(c *gin.Context) {
	posts, err := repository.AllPostList()
	if err != nil {
		//记录错误日志。。。
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, posts)
}

// 获取文章下的所有评论
func AllCommentControl(c *gin.Context) {
	var com models.Comment
	if err := c.ShouldBindJSON(&com); err != nil {
		//记录错误日志。。。
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	comments := repository.GetCommentById(com.PostId)
	if comments == nil {
		//记录错误日志。。。
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
	}
	c.JSON(http.StatusOK, comments)
}

func checkPostpParam(post models.Post) bool {
	if post.Data == "" {
		return true
	}
	if post.Title == "" {
		return true
	}
	return false
}

// 删除文章 只有作者才行
func DeletePostControl(c *gin.Context) {
	// 从上下文中取出已经绑定的 post
	postInterface, exists := c.Get("post")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章数据未找到"})
		return
	}
	post := postInterface.(models.Post)

	err := repository.DeletePost(post.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除文章成功",
		"code":    "sucess",
	})
}

// 更新文章 只有作者才行
func UpdatePostControl(c *gin.Context) {

	fmt.Println("UpdatePostControl 开始")

	// 从上下文中取出已经绑定的 post
	postInterface, exists := c.Get("post")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章数据未找到"})
		return
	}
	post := postInterface.(models.Post)

	//更新文章
	err := repository.UpdatePost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "更新文章成功",
		"code":    "sucess",
	})
}
