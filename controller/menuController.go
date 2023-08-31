package controller

import (
	"backEnd/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 每个页面刚加载时获取左侧菜单信息
func GetPageMenu(ctx *gin.Context) {
	// 与前端约定好字符串
	pageName := ctx.Query("pageKind")
	fmt.Println(pageName)
	if len(pageName) == 0 {
		ctx.String(http.StatusBadRequest, "未得到页面名称，无法返回菜单配置")
		return
	}
	menuList, err := utils.LoadJson("./data/config.json")
	
	if !err {
		fmt.Println("加载文件失败")
		return
	}
	menu := menuList[pageName]
	fmt.Println(menu)
	ctx.JSON(http.StatusOK, menu)
}

// 增加左侧菜单栏的选项
func AddPageMenuItem(ctx *gin.Context) {
	pageName := ctx.PostForm("pageKind")
	menuItemName := ctx.PostForm("menuItemName")
	menuItemList := ctx.PostForm("menuItemList")
	fmt.Println("pageName: " + pageName)
	fmt.Println("menuItemName: " + menuItemName)
	fmt.Println(menuItemList)
	// fmt.Println(tmp)
}
