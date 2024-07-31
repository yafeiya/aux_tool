package controller

import (
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// TODO:获取当前界面实例列表
/* 参数：前端传来的参数pageKind
pageKind:页面类型，字符串分别为database、modelbase
*/
// 返回：pageKind的卡片列表
func GetExampleList(ctx *gin.Context) {
	db := common.InitDB()
	// 与前端约定好字符串
	pageKind := ctx.Query("pageKind")
	fmt.Println(pageKind)

	if pageKind == "example" {
		examples := []model.Example{}
		db.Find(&examples)
		if len(examples) == 0 {
			fmt.Println("没找到相关实例")
			response.Response(ctx, http.StatusOK, 404, nil, "fail")
		} else {
			fmt.Println("example:", examples)
			response.Success(ctx, gin.H{"examples": examples}, "success")
		}
	}
}

// TODO:删除指定实例
// 参数：前端传来的参数pageKind和id。
// 返回：fail或success
func DeleteExample(ctx *gin.Context) {
	db := common.InitDB()
	pageKind := ctx.Query("pagekind")
	fmt.Println(pageKind)
	Id := ctx.Query("id")
	fmt.Println(Id)
	IdList := strings.Split(Id, "/")
	fmt.Println("IdList", IdList)
	if pageKind == "example" {
		for _, value := range IdList {
			if value == "" {
				fmt.Println("记录删除结束")
			} else {
				fmt.Println("vvvvvvvvvvvvvvvvalueid", value)
				example := []model.Example{}
				result := db.Where("Id = ?", value).Delete(&example)
				fmt.Println("rrrrrrrrrrrr", result)
				if result.RowsAffected == 0 {
					fmt.Println("删除失败")
					response.Response(ctx, http.StatusOK, 404, nil, "fail")
				} else {
					fmt.Println("删除成功")
					response.Success(ctx, nil, "success")
				}
			}
		}
	}
}




/*
TODO:更新指定实例
参数：前端传来的参数
返回：fail或success
*/
func UpdateExample(ctx *gin.Context) {
	db := common.InitDB()
	Id := ctx.Query("id")
	State := ctx.Query("state")
	fmt.Println(Id)
	fmt.Println(State)
	IdList := strings.Split(Id, "/")
	fmt.Println("IdList", IdList)
	// 获取当前时间并格式化为字符串
	currentTime := strconv.FormatInt(time.Now().UnixMilli(), 10)
	example := model.Example{
		State: State,
		End_time: currentTime,

	}
	for _, value := range IdList {
		if value == "" {
			fmt.Println("记录更新结束")
		} else {
			result := db.Model(model.Example{}).Where("Id = ?", value).Updates(&example)
			if result.RowsAffected == 0 {
				fmt.Println("修改失败")
				response.Response(ctx, http.StatusOK, 404, nil, "fail")
			} else {
				// 新增卡片
				fmt.Println("修改成功")
				response.Success(ctx, nil, "success")
			}
		}
	}


}
