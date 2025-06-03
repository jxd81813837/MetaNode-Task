package router

import (
	"MetaNode-Task/task-go/go_base/task-4-blog/controller"
	"MetaNode-Task/task-go/go_base/task-4-blog/middleware"
	"github.com/gin-gonic/gin"
)

func RouterPath() {
	r := gin.Default()
	r.POST("/register", controller.RegisterController)
	r.POST("/login", controller.LoginController)

	//执行分组方法
	userGroup := r.Group("/user", middleware.AuthMiddlerWareBlog())
	{
		userGroup.POST("/createPost", controller.CreatePostControl)
		userGroup.POST("/createCommon", controller.CreateCommentControl)
		userGroup.POST("/updatePost", append([]gin.HandlerFunc{middleware.AuthAuthorMiddlerWare()}, controller.UpdatePostControl)...)
		userGroup.POST("/deletePost", append([]gin.HandlerFunc{middleware.AuthAuthorMiddlerWare()}, controller.DeletePostControl)...)
	}
	r.GET("/allPost", controller.AllPostListControl)
	r.GET("/allComment", controller.AllCommentControl)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
