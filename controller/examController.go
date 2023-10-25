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
		//example := []model.Example{}
		//result := db.Where("Id = ?", Id).Delete(&example)
		//if result.RowsAffected == 0 {
		//	fmt.Println("删除失败")
		//	response.Response(ctx, http.StatusOK, 404, nil, "fail")
		//} else {
		//	fmt.Println("删除成功")
		//	response.Success(ctx, nil, "success")
		//}
	}
}

/*
TODO:更新指定实例
参数：前端传来的参数
返回：fail或success
*/
func UpdateExample(ctx *gin.Context) {
	db := common.InitDB()
	pageKind := ctx.PostForm("pageKind")
	fmt.Println(pageKind)
	if pageKind == "example" {
		example_name := "example_name"
		rank := "example_name"
		state := "example_name"
		cpu_num := "5"
		cpu, _ := strconv.ParseUint(cpu_num, 10, 64)
		gpu_num := "5"
		gpu, _ := strconv.ParseUint(gpu_num, 10, 64)
		post_data := "21342134214"
		post_time, _ := strconv.ParseUint(post_data, 10, 64)
		dataset_url := "example_name"
		model_name := "example_name"
		model_type := "example_name"
		epoch_num := "example_name"
		loss := "example_name"
		optimizer := "example_name"
		decay := "example_name"
		evaluation := "example_name"
		model_url := "example_name"
		memory := "example_name"
		start_time := "25234234234"
		end_time := "254235234"
		example := model.Example{
			Example_name: example_name,
			Rank:         rank,
			State:        state,
			Cpu_num:      cpu,
			Gpu_num:      gpu,
			Post_data:    post_time,
			Dataset_url:  dataset_url,
			Model_name:   model_name,
			Model_type:   model_type,
			Epoch_num:    epoch_num,
			Loss:         loss,
			Optimizer:    optimizer,
			Decay:        decay,
			Evaluation:   evaluation,
			Model_url:    model_url,
			Memory:       memory,
			Start_time:   start_time,
			End_time:     end_time,
		}
		// 判重处理
		//pageKind、task、type、dataset_name
		result := db.Model(model.Example{}).Where("Example_name = ?", example_name).Updates(&example)
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
