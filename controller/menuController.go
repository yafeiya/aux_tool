package controller

import (
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/model"
	"backEnd/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	Task := ctx.Query("task")
	Type := ctx.Query("type")
	fmt.Println(id)
	if pageKind == "database" {
		database := []model.Database{}
		result := db.Where("Id = ?", id).Delete(&database)
		fmt.Println("idddddddddd", Type, Task, id)
		erro := os.RemoveAll("./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + id)
		if erro != nil {
			fmt.Println("delete fail")
		}
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
		fmt.Println("idddddddddd", Type, Task, id)
		erro := os.RemoveAll("./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + id)
		if erro != nil {
			fmt.Println("delete fail")
		}
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

	//fmt.Println("1111111111111111111", pageKind)
	task := ctx.Query("task")
	//fmt.Println(task)
	Type := ctx.Query("Type")
	//fmt.Println(Type)

	if pageKind == "database" {
		database := []model.Database{}
		//fmt.Println("wwwww", database)

		db.Where("Task = ? and Type = ?", task, Type).Find(&database)
		//fmt.Println("666666666", database)
		if len(database) == 0 {
			//fmt.Println("未找到记录")
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
			//fmt.Println("未找到记录")
			response.Response(ctx, http.StatusOK, 404, nil, "fail")
		} else {
			//fmt.Println(modelbase)
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
	fmt.Println("11111111111111111111delete")
	id, _ := strconv.Atoi(ctx.Query("id"))
	fmt.Println("id:", id)

	err := model.DeleteDesign(id)
	if err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	response.Success(ctx, nil, "success")
}

// 添加Definefunction卡片
func AddDefinefunctionCard(ctx *gin.Context) {
	fmt.Println("AddDefinefunctionCard")
	released := ctx.Query("released")
	dataset_name := ctx.Query("dataset_name")
	alias := ctx.Query("alias")
	ttype := ctx.Query("type")
	rank := ctx.Query("rank")
	paramsStr := ctx.Query("params")

	fmt.Println("params", paramsStr)
	// 解析 params 字段
	params, er := parseParams(paramsStr)
	if er != nil {
		fmt.Println("Invalid params format")
		return
	}

	function_body := ctx.Query("function_body")
	lan := ctx.Query("lan")
	code := ctx.Query("code")
	description := ctx.Query("description")
	task := ctx.Query("task")

	// 添加新设计
	err := model.AddDefinefunction(released, dataset_name, alias, ttype, rank, params, function_body, lan, code, description, task)
	if err != nil {
		fmt.Println("Failed to add definefunction")
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	response.Success(ctx, nil, "success")
}

func GetDefinefunctionCards(ctx *gin.Context) {
	fmt.Println("GetDefinefunctionCards")
	ttype := ctx.Query("type")
	task := ctx.Query("task")
	definefunctions, err := model.GetDefinefunctions(task, ttype)
	if err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}

	response.Success(ctx, gin.H{"definefunctions": definefunctions}, "success")
}

func UpdateDefinefunctionCard(ctx *gin.Context) {
	fmt.Println("UpdateDefinefunctionCard")

	id, _ := strconv.Atoi(ctx.Query("id"))
	dataset_name := ctx.Query("dataset_name")
	alias := ctx.Query("alias")
	rank := ctx.Query("rank")
	paramsStr := ctx.Query("params")

	// 解析 params 字段
	params, er := parseParams(paramsStr)
	if er != nil {
		fmt.Println("Invalid params format")
		return
	}

	function_body := ctx.Query("function_body")
	lan := ctx.Query("lan")
	code := ctx.Query("code")
	description := ctx.Query("description")

	err := model.UpdateDefinefunction(id, dataset_name, alias, rank, params, function_body, lan, code, description)
	if err != nil {
		fmt.Println(err)
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}

	response.Success(ctx, nil, "success")
}

// 删除Definefunction卡片
func DeleteDefinefunctionCard(ctx *gin.Context) {
	fmt.Println("DeleteDefinefunctionCard")
	id, _ := strconv.Atoi(ctx.Query("id"))
	fmt.Println("id:", id)

	err := model.DeleteDefinefunction(id)
	if err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	response.Success(ctx, nil, "success")
}

// 解析 params 字段的函数
func parseParams(paramsStr string) ([][]float64, error) {
	// 以 "|" 分割参数组
	paramGroups := strings.Split(paramsStr, "|")

	var params [][]float64

	for _, group := range paramGroups {
		// 以 "," 分割参数
		paramParts := strings.Split(group, ",")
		if len(paramParts) != 2 {
			return nil, fmt.Errorf("parseParams error")
		}

		// 将字符串转换为 float64
		param1, err := strconv.ParseFloat(paramParts[0], 64)
		if err != nil {
			return nil, fmt.Errorf("parseParams error")
		}

		param2, err := strconv.ParseFloat(paramParts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("parseParams error")
		}

		params = append(params, []float64{param1, param2})
	}

	return params, nil
}
