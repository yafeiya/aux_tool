package controller

import (
	"encoding/json"
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/model"
	"fmt"
	"net/http"
	"backEnd/utils"
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
	pageKind := ctx.PostForm("pageKind")
	name := ctx.PostForm("name")
	title := ctx.PostForm("title")
	icon := ctx.PostForm("icon")
	childrenlist := ctx.PostForm("childrenlist")
	// 在实际应用中，你需要将 childrenList 解析为切片，这里假设它是一个 JSON 字符串
	var children []model.Children
	if err := json.Unmarshal([]byte(childrenlist), &children); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid childrenlist format"})
		return
	}
	fmt.Println("pageKind: " + pageKind)

	filename := "./data/config.json"

	// 创建新的模型实例
	newData := model.DataMenu{
		Name:  name,
		Title: title,
		Icon:  icon,
		Children: children,
	}

	er := model.AddMenu(pageKind, filename, newData)
	if er != nil {
		fmt.Println(er)
	}
}

// TODO:获取当前界面卡片列表
/* 参数：前端传来的参数pageKind、task和type。
pageKind:页面类型，字符串分别为database、modelbase
task：
type：
*/
// 返回：pageKind的卡片列表
func GetMenuList(ctx *gin.Context) {
	db := common.InitDB()
	// 与前端约定好字符串
	pageKind := ctx.PostForm("pageKind")
	fmt.Println(pageKind)
	task := ctx.PostForm("task")
	fmt.Println(task)
	Type := ctx.PostForm("type")
	fmt.Println(Type)

	if pageKind == "database" {
		databases := []model.Database{}
		db.Where("Task = ? and Type = ?", task, Type).Find(&databases)
		if len(databases) == 0 {
			fmt.Println("没找到task=" + task + "并且type=" + Type + "的数据集卡片")
			response.Response(ctx, http.StatusOK, 404, nil, "No corresponding card found")
		} else {
			fmt.Println(databases)
			response.Success(ctx, gin.H{"databases": databases}, "success")
		}
	}
	if pageKind == "modelbase" {
		modelbases := []model.Modelbase{}
		db.Where("Task = ? and Type = ?", task, Type).Find(&modelbases)
		if len(modelbases) == 0 {
			fmt.Println("没找到task=" + task + "并且type=" + Type + "的模型卡片")
			response.Response(ctx, http.StatusOK, 404, nil, "No corresponding card found")
		} else {
			fmt.Println(modelbases)
			response.Success(ctx, gin.H{"databases": modelbases}, "success")
		}
	}
}

// TODO:删除指定卡片
// 参数：前端传来的参数pageKind和id。
// 返回：卡片列表,即database的model列表databases
func DeleteCard(ctx *gin.Context) {
	db := common.InitDB()

	pageKind := ctx.PostForm("pageKind")
	fmt.Println(pageKind)
	id := ctx.PostForm("id")
	fmt.Println(id)

	if pageKind == "database" {
		database := []model.Database{}
		result := db.Where("Id = ?", id).Delete(&database)
		if result.RowsAffected == 0 {
			fmt.Println("删除失败")
			response.Response(ctx, http.StatusOK, 404, nil, "fail")
		} else {
			fmt.Println("删除成功")
			response.Success(ctx, nil, "success")
		}
	}
	if pageKind == "modelbase" {
		modelbase := []model.Modelbase{}
		result := db.Where("Id = ?", id).Delete(&modelbase)
		if result.RowsAffected == 0 {
			fmt.Println("删除失败")
			response.Response(ctx, http.StatusOK, 404, nil, "fail")
		} else {
			fmt.Println("删除成功")
			response.Success(ctx, nil, "success")
		}
	}
}

// TODO:获取指定卡片
// 参数：前端传来的参数pageKind、task、type、dataset_name。
// 返回：卡片列表,即database的model列表databases
func GetCard(ctx *gin.Context) {
	db := common.InitDB()

	pageKind := ctx.PostForm("pageKind")
	fmt.Println(pageKind)
	task := ctx.PostForm("task")
	fmt.Println(task)
	Type := ctx.PostForm("Type")
	fmt.Println(Type)
	dataset_name := ctx.PostForm("dataset_name")
	fmt.Println(dataset_name)

	if pageKind == "database" {
		database := model.Database{}
		db.Where("Task = ? and Type = ? and Dataset_name = ?", task, Type, dataset_name).First(&database)
		if database.Id == 0 {
			fmt.Println("未找到记录")
			response.Response(ctx, http.StatusOK, 404, nil, "fail")
		} else {
			fmt.Println(database)
			response.Success(ctx, gin.H{"database": database}, "success")
		}
	}
	if pageKind == "modelbase" {
		modelbase := model.Modelbase{}
		db.Where("Task = ? and Type = ? and Dataset_name = ?", task, Type, dataset_name).First(&modelbase)
		if modelbase.Id == 0 {
			fmt.Println("未找到记录")
			response.Response(ctx, http.StatusOK, 404, nil, "fail")
		} else {
			fmt.Println(modelbase)
			response.Success(ctx, gin.H{"modelbase": modelbase}, "success")
		}
	}
}

/*
	TODO:添加指定卡片
参数：前端传来的参数:

	release
	dataset_name
	type
	rank
	character_type
	header
	data_path
	description
	task

返回：卡片列表,即database的model列表databases
*/
func CreateCard(ctx *gin.Context) {
	db := common.InitDB()
	pageKind := ctx.PostForm("pageKind")
	fmt.Println(pageKind)
	if pageKind == "database" {
		released := ctx.PostForm("released")
		fmt.Println(released)
		dataset_name := ctx.PostForm("dataset_name")
		fmt.Println(dataset_name)
		Type := ctx.PostForm("type")
		fmt.Println(Type)
		rank := ctx.PostForm("rank")
		fmt.Println(rank)
		character_type := ctx.PostForm("character_type")
		fmt.Println(character_type)
		header := ctx.PostForm("header")
		fmt.Println(header)
		data_path := ctx.PostForm("data_path")
		fmt.Println(data_path)
		description := ctx.PostForm("description")
		fmt.Println(description)
		task := ctx.PostForm("task")
		fmt.Println(task)
		database := model.Database{
			Released:       released,
			Dataset_name:   dataset_name,
			Type:           Type,
			Rank:           rank,
			Character_type: character_type,
			Header:         header,
			Data_path:      data_path,
			Description:    description,
			Task:           task,
		}
		// 判重处理
		//pageKind、task、type、dataset_name
		db.Where("Task = ? and Type = ? and Dataset_name = ?", task, Type, dataset_name).First(&database)
		if database.Id != 0 {
			fmt.Println("该卡片已存在")
			response.Response(ctx, http.StatusOK, 404, nil, "The card already exists")
		} else {
			// 新增卡片
			db.Create(&database)
			response.Success(ctx, nil, "success")
		}
	}
	if pageKind == "modelbase" {
		released := ctx.PostForm("released")
		fmt.Println(released)
		dataset_name := ctx.PostForm("dataset_name")
		fmt.Println(dataset_name)
		Type := ctx.PostForm("type")
		fmt.Println(Type)
		rank := ctx.PostForm("rank")
		fmt.Println(rank)
		lan := ctx.PostForm("lan")
		fmt.Println(lan)
		data_path := ctx.PostForm("data_path")
		fmt.Println(data_path)
		description := ctx.PostForm("description")
		fmt.Println(description)
		code := ctx.PostForm("code")
		fmt.Println(code)
		task := ctx.PostForm("task")
		fmt.Println(task)
		modelbase := model.Modelbase{
			Released:     released,
			Dataset_name: dataset_name,
			Type:         Type,
			Rank:         rank,
			Lan:          lan,
			Data_path:    data_path,
			Description:  description,
			Code:         code,
			Task:         task,
		}
		// 判重处理
		//pageKind、task、type、dataset_name
		db.Where("Task = ? and Type = ? and Dataset_name = ?", task, Type, dataset_name).First(&modelbase)
		if modelbase.Id != 0 {
			fmt.Println("该卡片已存在")
			response.Response(ctx, http.StatusOK, 404, nil, "The card already exists")
		} else {
			// 新增卡片
			db.Create(&modelbase)
			response.Success(ctx, nil, "success")
		}
	}
}

// TODO:修改指定卡片
// 参数：前端传来的参数pageKind和id。
// 返回：卡片列表,即database的model列表databases
func UpdataCard(ctx *gin.Context) {

}
