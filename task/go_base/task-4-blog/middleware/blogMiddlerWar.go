package middleware

import (
	"MetaNode-Task/task/go_base/task-4-blog/config"
	"MetaNode-Task/task/go_base/task-4-blog/models"
	"MetaNode-Task/task/go_base/task-4-blog/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 认证中间件
func AuthMiddlerWareBlog() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
			c.Abort()
			return
		}
		authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Println("中间件 authHeader:", authHeader)

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JWT_SECRET), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Println("用户ID:", claims["id"])
		fmt.Println("用户名:", claims["username"])
		fmt.Println("过期时间:", claims["exp"])
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		// 将用户信息存入上下文，方便后续处理函数使用
		c.Set("userID", claims["id"])
		c.Set("username", claims["username"])
		fmt.Println("AuthMiddlerWareBlog：中间件 after 完成")
		c.Next() // 执行后续的处理函数

	}
}

// 文章认证权限 ，认证这个文章是不是作者的
func AuthAuthorMiddlerWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthAuthorMiddlerWare 开始")
		var post models.Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// 取出 userID
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "userID not found in context"})
			c.Abort()
			return
		}
		post.AuthorId = uint(userID.(float64))
		//验证是不是 文章是不是作者的
		if !repository.CheckPostOwnerShip(post) {
			c.JSON(http.StatusAccepted, gin.H{"error": "userID not Post in id 该文章不是作者的"})
			c.Abort()
			return
		}
		c.Set("post", post)
		fmt.Println("AuthAuthorMiddlerWare 认证作者：中间件 after 完成")
		c.Next() // 执行后续的处理函数
	}
}
