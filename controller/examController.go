package controller

import (
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/model"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// TODO:获取当前界面实例列表
/* 参数：前端传来的参数pageKind
pageKind:页面类型，字符串分别为database、modelbase
*/
// 返回：pageKind的卡片列表
func GetExampleList(ctx *gin.Context) {
	db := common.InitDB()
	// 与前端约定好字符串
	pageKind := ctx.PostForm("pageKind")
	fmt.Println(pageKind)

	if pageKind == "example" {
		examples := []model.Example{}
		db.Find(&examples)
		if len(examples) == 0 {
			fmt.Println("没找到相关实例")
			response.Response(ctx, http.StatusOK, 404, nil, "fail")
		} else {
			fmt.Println(examples)
			response.Success(ctx, gin.H{"examples": examples}, "success")
		}
	}
}

// TODO:删除指定实例
// 参数：前端传来的参数pageKind和id。
// 返回：fail或success
func DeleteExample(ctx *gin.Context) {
	db := common.InitDB()

	pageKind := ctx.PostForm("pageKind")
	fmt.Println(pageKind)
	id := ctx.PostForm("id")
	fmt.Println(id)

	if pageKind == "example" {
		example := []model.Example{}
		result := db.Where("Id = ?", id).Delete(&example)
		if result.RowsAffected == 0 {
			fmt.Println("删除失败")
			response.Response(ctx, http.StatusOK, 404, nil, "fail")
		} else {
			fmt.Println("删除成功")
			response.Success(ctx, nil, "success")
		}
	}
}


/*TODO:添加指定实例
参数：前端传来的参数
返回：The example already exists或success
*/
func CreateExample(ctx *gin.Context) {
	db := common.InitDB()
	pageKind := ctx.PostForm("pageKind")
	fmt.Println(pageKind)
	if pageKind == "example" {
		example_name := ctx.PostForm("example_name")
		fmt.Println(example_name)
		rank := ctx.PostForm("rank")
		fmt.Println(rank)
		state := ctx.PostForm("state")
		fmt.Println(state)
		cpu_num := ctx.PostForm("cpu_num")
		cpu, _ := strconv.ParseUint(cpu_num, 10, 64)
		fmt.Println(cpu_num)
		gpu_num := ctx.PostForm("gpu_num")
		gpu, _ := strconv.ParseUint(gpu_num, 10, 64)
		fmt.Println(gpu_num)
		post_data := ctx.PostForm("post_data")
		post_time, _ := strconv.ParseUint(post_data, 10, 64)
		fmt.Println(post_data)
		dataset_url := ctx.PostForm("dataset_url")
		fmt.Println(dataset_url)
		model_name := ctx.PostForm("model_name")
		fmt.Println(model_name)
		model_type := ctx.PostForm("model_type")
		fmt.Println(model_type)
		epoch_num := ctx.PostForm("epoch_num")
		fmt.Println(epoch_num)
		loss := ctx.PostForm("loss")
		fmt.Println(loss)
		optimizer := ctx.PostForm("optimizer")
		fmt.Println(optimizer)
		decay := ctx.PostForm("decay")
		fmt.Println(decay)
		evaluation := ctx.PostForm("evaluation")
		fmt.Println(evaluation)
		model_url := ctx.PostForm("model_url")
		fmt.Println(model_url)
		memory := ctx.PostForm("memory")
		fmt.Println(memory)
		start_time := ctx.PostForm("start_time")
		s_time, _ := strconv.ParseUint(start_time, 10, 64)
		fmt.Println(start_time)
		end_time := ctx.PostForm("end_time")
		e_time, _ := strconv.ParseUint(end_time, 10, 64)
		fmt.Println(end_time)
		example := model.Example{
			Example_name:   	example_name,
			Rank:   			rank,
			State:           	state,
			Cpu_num:           	cpu,
			Gpu_num: 			gpu,
			Post_data:         	post_time,
			Dataset_url:      	dataset_url,
			Model_name:    		model_name,
			Model_type:         model_type,
			Epoch_num:          epoch_num,
			Loss:           	loss,
			Optimizer:          optimizer,
			Decay:           	decay,
			Evaluation:         evaluation,
			Model_url:          model_url,
			Memory:           	memory,
			Start_time:         s_time,
			End_time:           e_time,
		}
		// 判重处理
		//pageKind、task、type、dataset_name
		db.Where("Example_name = ?", example).First(&example)
		if example.Id != 0 {
			fmt.Println("该实例已存在")
			response.Response(ctx, http.StatusOK, 404, nil, "The example already exists")
		} else {
			// 新增
			db.Create(&example)
			response.Success(ctx, nil, "success")
		}
	}
}

/*TODO:更新指定实例
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
		s_time, _ := strconv.ParseUint(start_time, 10, 64)
		end_time := "254235234"
		e_time, _ := strconv.ParseUint(end_time, 10, 64)
		example := model.Example{
			Example_name:   	example_name,
			Rank:   			rank,
			State:           	state,
			Cpu_num:           	cpu,
			Gpu_num: 			gpu,
			Post_data:         	post_time,
			Dataset_url:      	dataset_url,
			Model_name:    		model_name,
			Model_type:         model_type,
			Epoch_num:          epoch_num,
			Loss:           	loss,
			Optimizer:          optimizer,
			Decay:           	decay,
			Evaluation:         evaluation,
			Model_url:          model_url,
			Memory:           	memory,
			Start_time:         s_time,
			End_time:           e_time,
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