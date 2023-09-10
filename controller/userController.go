package controller

import (
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册用户信息
func Register(ctx *gin.Context) {
	db := common.InitDB()
	// 与前端约定好字符串
	name := ctx.PostForm("name")
	fmt.Println(name)
	pwd := ctx.PostForm("pwd")
	fmt.Println(pwd)
	user := model.User{
		UserName: name,
		PassWord: pwd,
	}
	// 判重处理
	db.Where("UserName = ?", name).First(&user)
	if user.Id != 0 {
		fmt.Println("该用户名已存在")
		response.Response(ctx, http.StatusOK, 404, nil, "The username already exists")
	} else {
		// 新增该注册用户
		db.Create(&user)
		response.Success(ctx, nil, "register success")
	}
}

// 登录
func Login(ctx *gin.Context) {
	db := common.InitDB()
	// 与前端约定好字符串
	name := ctx.PostForm("name")
	fmt.Println(name)
	pwd := ctx.PostForm("pwd")
	fmt.Println(pwd)
	user := model.User{}
	// 执行查询语句，查询是否存在该用户
	db.Where("UserName = ?", name).First(&user)
	if user.Id == 0 {
		fmt.Println("未找到匹配的记录")
		response.Response(ctx, http.StatusOK, 404, nil, "The username does not exist")
	} else {
		fmt.Println("找到了匹配的记录")
		fmt.Println("username：" + user.UserName + "  password:" + user.PassWord)
		if user.PassWord != pwd {
			fmt.Println("密码错误")
			response.Response(ctx, http.StatusOK, 404, nil, "Password error")
		} else {
			fmt.Println("密码正确")
			response.Response(ctx, http.StatusOK, 200, gin.H{"user": user}, "login success")
		}
	}
}
