package controller

import (
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/model"
	"backEnd/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 每个页面刚加载时获取左侧菜单信息
func GetPageMenu(ctx *gin.Context) {
	// 与前端约定好字符串
	pageName := ctx.Query("pageKind")
	fmt.Println("pagename111111111111111:", pageName)
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
	fmt.Println("AddPageMenuItem")
	pageKind := ctx.Query("pageKind")
	title := ctx.Query("title")
	children_title := ctx.Query("children_title")

	if title == "" {
		response.Success(ctx, nil, "title=null")
		return
	}
	if children_title == "" {
		response.Success(ctx, nil, "children_title=null")
		return
	}
	filename := "./data/config.json"
	result := model.AddMenu(pageKind, filename, title, children_title)
	if result != "success" {
		fmt.Println(result)
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
	}
	response.Success(ctx, nil, "success")
}

func DeletePageMenuItem(ctx *gin.Context) {
	fmt.Println("DeletePageMenuItem")
	pageKind := ctx.Query("pageKind")
	title := ctx.Query("title")
	children_title := ctx.Query("children_title")
	filename := "./data/config.json"

	result := model.DeleteChildFromDataFile(filename, pageKind, title, children_title)
	if result != "success" {
		fmt.Println(result)
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
	}
	response.Success(ctx, nil, "success")
}

// TODO:删除指定卡片
// 参数：前端传来的参数pageKind和id。
// 返回：卡片列表,即database的model列表databases
func DeleteCard(ctx *gin.Context) {
	db := common.InitDB()
	pageKind := ctx.Query("pageKind")
	fmt.Println(pageKind)
	id := ctx.Query("id")
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

// 参数：前端传来的参数pageKind、task、type、dataset_name。
// 返回：卡片列表,即database的model列表databases
func GetCard(ctx *gin.Context) {

	db := common.InitDB()
	pageKind := ctx.Query("pageKind")

	fmt.Println("1111111111111111111", pageKind)
	task := ctx.Query("task")
	fmt.Println(task)
	Type := ctx.Query("Type")
	fmt.Println(Type)

	if pageKind == "database" {
		database := []model.Database{}
		//fmt.Println("wwwww", database)

		db.Where("Task = ? and Type = ?", task, Type).Find(&database)
		//fmt.Println("666666666", database)
		if len(database) == 0 {
			fmt.Println("未找到记录")
			response.Response(ctx, http.StatusOK, 404, nil, "fail")
		} else {
			//fmt.Println(database)
			//fmt.Println("wwwww")
			response.Success(ctx, gin.H{"database": database}, "success")
		}
	}
	if pageKind == "modelbase" {
		modelbase := []model.Modelbase{}
		db.Where("Task = ? and Type = ? ", task, Type).Find(&modelbase)
		if len(modelbase) == 0 {
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
	pageKind := ctx.Query("pageKind")
	fmt.Println(pageKind)
	if pageKind == "database" {
		released := ctx.Query("released")
		fmt.Println(released)
		dataset_name := ctx.Query("dataset_name")
		fmt.Println(dataset_name)
		Type := ctx.Query("type")
		fmt.Println(Type)
		rank := ctx.Query("rank")
		fmt.Println(rank)
		character_type := ctx.Query("character_type")
		fmt.Println(character_type)
		header := ctx.Query("header")
		fmt.Println(header)
		// data_path := ctx.PostForm("data_path")
		// fmt.Println(data_path)
		description := ctx.Query("description")
		fmt.Println(description)
		task := ctx.Query("task")
		fmt.Println(task)
		database := model.Database{
			Released:       released,
			Dataset_name:   dataset_name,
			Type:           Type,
			Rank:           rank,
			Character_type: character_type,
			Header:         header,
			Data_path:      "./" + task + "/" + Type + "/" + dataset_name + "/",
			Description:    description,
			Task:           task,
		}
		// 判重处理
		//pageKind、task、type、dataset_name
		db.Where("Task = ? and Type = ? and Dataset_name = ?", task, Type, dataset_name).First(&database)
		if database.Id != 0 {
			fmt.Println("该卡片已存在")
			response.Response(ctx, http.StatusOK, 404, nil, "fail")
		} else {
			// 新增卡片
			db.Create(&database)
			response.Success(ctx, nil, "success")
		}
	}
	if pageKind == "modelbase" {
		released := ctx.Query("released")
		fmt.Println(released)
		dataset_name := ctx.Query("dataset_name")
		fmt.Println(dataset_name)
		Type := ctx.Query("type")
		fmt.Println(Type)
		rank := ctx.Query("rank")
		fmt.Println(rank)
		lan := ctx.Query("lan")
		fmt.Println(lan)
		// data_path := ctx.PostForm("data_path")
		// fmt.Println(data_path)
		description := ctx.Query("description")
		fmt.Println(description)
		code := ctx.Query("code")
		fmt.Println(code)
		task := ctx.Query("task")
		fmt.Println(task)
		modelbase := model.Modelbase{
			Released:     released,
			Dataset_name: dataset_name,
			Type:         Type,
			Rank:         rank,
			Lan:          lan,
			Data_path:    "./" + task + "/" + Type + "/" + dataset_name + "/",
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

	db := common.InitDB()
	pageKind := ctx.Query("pageKind")
	fmt.Println(pageKind)
	//cardInfo := ctx.Query("cardInfo")
	//fmt.Println("111111111111111cardinfo", cardInfo, "22222222222222222222")
	if pageKind == "database" {
		//released := ctx.PostForm("released")
		//fmt.Println(released)
		Id := ctx.Query("id")
		fmt.Println(Id)
		dataset_name := ctx.Query("dataset_name")
		fmt.Println(dataset_name)
		Type := ctx.Query("type")
		fmt.Println(Type)
		rank := ctx.Query("rank")
		fmt.Println(rank)
		character_type := ctx.Query("character_type")
		fmt.Println(character_type)
		header := ctx.Query("header")
		fmt.Println(header)
		// data_path := ctx.PostForm("data_path")
		// fmt.Println(data_path)
		description := ctx.Query("description")
		fmt.Println(description)
		task := ctx.Query("task")
		fmt.Println(task)

		var database model.Database

		// 判重处理
		//pageKind、task、type、dataset_name
		db.Where("Id = ?", Id).First(&database)
		if database.Id != 0 {
			fmt.Println("该卡片存在")
			//database.Released = released
			database.Dataset_name = dataset_name
			database.Type = Type
			database.Rank = rank
			database.Data_path = "./" + task + "/" + Type + "/" + dataset_name + "/"
			database.Description = description
			database.Character_type = character_type
			database.Task = task
			database.Header = header
			er := db.Save(&database).Error
			if er != nil {
				fmt.Println("无法保存更新后的记录")
			}
			fmt.Println("uo")
			response.Success(ctx, nil, "success")
			//response.Response(ctx, http.StatusOK, 404, nil, "success")
		} else {
			fmt.Println("uo2222222222222222222222222222")
			response.Success(ctx, nil, "fail")
		}
	}
	if pageKind == "modelbase" {
		//released := ctx.PostForm("released")
		//fmt.Println(released)
		Id := ctx.Query("id")
		fmt.Println(Id)
		dataset_name := ctx.Query("dataset_name")
		fmt.Println(dataset_name)
		Type := ctx.Query("type")
		fmt.Println(Type)
		rank := ctx.Query("rank")
		fmt.Println(rank)
		lan := ctx.Query("lan")
		fmt.Println(lan)
		description := ctx.Query("description")
		fmt.Println(description)
		code := ctx.Query("code")
		fmt.Println(code)
		task := ctx.Query("task")
		fmt.Println(task)
		var modelbase model.Modelbase

		// 判重处理
		//pageKind、task、type、dataset_name
		db.Where("Id = ?", Id).First(&modelbase)
		if modelbase.Id != 0 {
			fmt.Println("该卡片存在")
			//modelbase.Released = released
			modelbase.Dataset_name = dataset_name
			modelbase.Type = Type
			modelbase.Rank = rank
			modelbase.Lan = lan
			modelbase.Data_path = "./" + task + "/" + Type + "/" + dataset_name + "/"
			modelbase.Description = description
			modelbase.Code = code
			modelbase.Task = task
			er := db.Save(&modelbase).Error
			if er != nil {
				fmt.Println("无法保存更新后的记录")
			}
			fmt.Println("uo")
			response.Success(ctx, nil, "success")
			//response.Response(ctx, http.StatusOK, 404, nil, "success")
		} else {
			response.Success(ctx, nil, "fail")
		}
	}
}

// 添加Design卡片
func AddDesignCard(ctx *gin.Context) {
	fmt.Println("AddDesignCard")
	released := ctx.Query("released")
	dataset_name := ctx.Query("dataset_name")
	ttype := ctx.Query("type")
	rank := ctx.Query("rank")
	description := ctx.Query("description")
	task := ctx.Query("task")

	// 添加新设计
	err := model.AddDesign(released, dataset_name, ttype, rank, description, task)
	if err != nil {
		fmt.Println("Failed to add design")
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	response.Success(ctx, nil, "success")
}

func GetDesignCards(ctx *gin.Context) {
	fmt.Println("GetDesignCards")
	ttype := ctx.Query("type")
	task := ctx.Query("task")
	designs, err := model.GetDesigns(task, ttype)
	if err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}

	response.Success(ctx, gin.H{"designs": designs}, "success")
}

func UpdateDesignCard(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Query("id"))
	dataset_name := ctx.Query("dataset_name")
	rank := ctx.Query("rank")
	description := ctx.Query("description")
	fmt.Println("description", description)
	err := model.UpdateDesign(id, dataset_name, rank, description)
	if err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}

	response.Success(ctx, nil, "success")
}

// 删除design卡片
func DeleteDesignCard(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	err := model.DeleteDesign(id)
	if err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}

	response.Success(ctx, nil, "success")
}
