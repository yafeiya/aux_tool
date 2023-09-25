package controller

import (
	"backEnd/common/response"
	"backEnd/model"
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"backEnd/common"
	// "reflect"
	// "io"
)

// 上传模型文件和模型的png文件
func UploadModelFile(ctx *gin.Context) {
	// db := common.InitDB()
	// 路径参数
	Task := ctx.PostForm("task")
	Dataset_name := ctx.PostForm("dataset_name")
	Type := ctx.PostForm("type")

	file, _ := ctx.FormFile("file")
	fmt.Println(file.Filename)

	// 要创建的文件夹的路径
	folderPath := "./" + Type + "/" + Task + "/" + Dataset_name + "/"

	// 使用os.Mkdir创建文件夹
	err := os.Mkdir(folderPath, 0755) // 0755是文件夹的权限设置
	if err != nil {
		fmt.Println("创建文件夹失败:", err)
	}
	
	dst := folderPath + file.Filename
	// 上传文件至指定的完整文件路径
	ctx.SaveUploadedFile(file, dst)
	response.Success(ctx, nil, "fail")
	
}

// 下载文件
func DownloadModelFile(ctx *gin.Context) {
	db := common.InitDB()
	// 路径参数
	Task := ctx.PostForm("task")
	Type := ctx.PostForm("type")
	Model_name := ctx.PostForm("model_name")
	File_type := ctx.PostForm("file_type")

	modelbase := []model.Modelbase{}

	db.Where("Task = ? and Type = ? and Dataset_name = ?",Task, Type, Model_name).Find(&modelbase)
	if len(modelbase) == 0 {
		fmt.Println("没找到")
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
	} else {
		if File_type == "png"{
			filePath := modelbase[0].Data_path + Model_name + ".png"
			fmt.Println(filePath)
			response.Response(ctx, http.StatusOK, 404, gin.H{"url": filePath}, filePath)
		}else{
			filePath := modelbase[0].Data_path + Model_name + ".pt"
			fmt.Println(filePath)
			response.Response(ctx, http.StatusOK, 404, gin.H{"url": filePath}, filePath)
		}
	}
}