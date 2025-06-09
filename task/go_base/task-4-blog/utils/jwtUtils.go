package utils

import (
	"MetaNode-Task/task/go_base/task-4-blog/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateJws(storedUser models.User) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.Id,
		"username": storedUser.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	return token
}
