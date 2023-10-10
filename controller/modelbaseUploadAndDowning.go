package controller

import (
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	// "reflect"
	// "io"
)

// 上传模型文件和模型的png文件
func UploadModelFile(ctx *gin.Context) {
	// 路径参数
	fmt.Println("UploadModelFile")
	db := common.InitDB()
	Task := ctx.PostForm("task")
	Type := ctx.PostForm("type")
	Id := ctx.PostForm("id")
	file, _ := ctx.FormFile("file")
	fmt.Println(file.Filename)
	fmt.Println("ID", Id)
	fmt.Println("Type", Type)
	fmt.Println("Task", Task)

	// 要创建的文件夹的路径
	folderPath := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id
	fmt.Println("111111111111111111", folderPath)
	// 使用os.Mkdir创建文件夹
	err := os.Mkdir(folderPath, 0755) // 0755是文件夹的权限设置
	if err != nil {
		fmt.Println("创建文件夹失败:", err)
	}
	dst := folderPath + "/" + file.Filename
	dst_sql := "http://127.0.0.1:5173/" + Type + "/" + Task + "/" + Id + "/" + file.Filename
	var modelbase model.Modelbase
	// 判重处理
	//pageKind、task、type、dataset_name
	db.Where("Id = ?", Id).First(&modelbase)
	if modelbase.Id != 0 {
		fmt.Println("该卡片存在")
		_, file_name := filepath.Split(modelbase.Data_path)
		erro := os.Remove("./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id + "/" + file_name)
		if erro != nil {
			fmt.Println("delete fail")
		}
		// 上传文件至指定的完整文件路径
		ctx.SaveUploadedFile(file, dst)

		modelbase.Data_path = dst_sql
		er := db.Save(&modelbase).Error
		if er != nil {
			fmt.Println("无法保存更新后的记录")
			response.Success(ctx, nil, "fail")
			return
		}
		response.Success(ctx, nil, "success")
		//response.Response(ctx, http.StatusOK, 404, nil, "success")
	} else {
		response.Success(ctx, nil, "fail")
	}
}

func UploadModelPNGFile(ctx *gin.Context) {
	// 路径参数
	fmt.Println("UploadModelPNGFile")
	db := common.InitDB()
	Task := ctx.PostForm("task")
	Type := ctx.PostForm("type")
	Id := ctx.PostForm("id")
	file, _ := ctx.FormFile("file")
	fmt.Println(file.Filename)
	fmt.Println("ID", Id)
	fmt.Println("Type", Type)
	fmt.Println("Task", Task)

	// 要创建的文件夹的路径
	folderPath := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id
	fmt.Println("111111111111111111", folderPath)
	// 使用os.Mkdir创建文件夹
	err := os.Mkdir(folderPath, 0755) // 0755是文件夹的权限设置
	if err != nil {
		fmt.Println("创建文件夹失败:", err)
	}
	dst := folderPath + "/" + file.Filename
	dst_sql := "http://127.0.0.1:5173/" + Type + "/" + Task + "/" + Id + "/" + file.Filename
	var modelbase model.Modelbase
	// 判重处理
	//pageKind、task、type、dataset_name
	db.Where("Id = ?", Id).First(&modelbase)
	if modelbase.Id != 0 {
		fmt.Println("该卡片存在")
		_, file_name := filepath.Split(modelbase.Image_path)
		erro := os.Remove("./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id + "/" + file_name)
		if erro != nil {
			fmt.Println("delete fail")
		}
		// 上传文件至指定的完整文件路径
		ctx.SaveUploadedFile(file, dst)

		modelbase.Image_path = dst_sql
		er := db.Save(&modelbase).Error
		if er != nil {
			fmt.Println("无法保存更新后的记录")
			response.Success(ctx, nil, "fail")
			return
		}
		response.Success(ctx, nil, "success")
		//response.Response(ctx, http.StatusOK, 404, nil, "success")
	} else {
		response.Success(ctx, nil, "fail")
	}
}

// 下载文件
func DownloadModelFile(ctx *gin.Context) {
	db := common.InitDB()
	Id := ctx.PostForm("id")
	var modelbase model.Modelbase
	db.Where("Id = ?", Id).First(&modelbase)
	if modelbase.Id != 0 {
		fmt.Println("该卡片存在")
		filePath := modelbase.Data_path
		fmt.Println(filePath)
		response.Response(ctx, http.StatusOK, 404, gin.H{"url": filePath}, filePath)
	} else {
		fmt.Println("没找到")
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
	}
}

// 获取模型图片
func GetModelImage(ctx *gin.Context) {
	db := common.InitDB()
	Id := ctx.PostForm("id")
	var modelbase model.Modelbase
	db.Where("Id = ?", Id).First(&modelbase)
	if modelbase.Id != 0 {
		fmt.Println("该卡片存在")
		filePath := modelbase.Image_path
		fmt.Println(filePath)
		response.Response(ctx, http.StatusOK, 404, gin.H{"url": filePath}, filePath)
	} else {
		fmt.Println("没找到")
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
	}
}
