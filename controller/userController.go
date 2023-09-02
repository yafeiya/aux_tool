package controller

import (
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostRegist(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	db := common.InitDB()
	user := model.User{
		Id:       0,
		UserName: username,
		PassWord: password,
	}
	db.Create(&user)
	fmt.Println("username: " + username + " password: " + password)
	response.Success(ctx, gin.H{"user": user}, "register success")
}
