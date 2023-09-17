package controller

import (
	"backEnd/common/response"
	// // "backEnd/model"
	"fmt"
	// "net/http"
	"os"
	"github.com/gin-gonic/gin"
)

// 上传文件
func UploadFile(ctx *gin.Context) {
		// 路径参数
		Task := ctx.Query("task")
		Dataset_name := ctx.Query("dataset_name")
		Type := ctx.Query("type")

		file, _ := ctx.FormFile("file")
		fmt.Println(file.Filename)

		// 要创建的文件夹的路径
		folderPath := "./" + Type + "/" + Task + "/" + Dataset_name

		// 使用os.Mkdir创建文件夹
		err := os.Mkdir(folderPath, 0755) // 0755是文件夹的权限设置
		if err != nil {
			fmt.Println("创建文件夹失败:", err)
		}
		
		dst := folderPath + file.Filename
		// 上传文件至指定的完整文件路径
		ctx.SaveUploadedFile(file, dst)

		// 读取csv文件信息，获取行数、列数、数据类型
		numColumns, numRows, Types := GetCSVInfo(folderPath + file.Filename)
		result := CreateTable(Task, Type, Dataset_name, numColumns, numRows, Types, file.Filename)
		if(result == "success"){
			response.Success(ctx, nil, "success")
		}else{
			response.Success(ctx, nil, "fail")
		}
}

// 下载文件
func DownloadFile(ctx *gin.Context) {
	// 路径参数
	// Task := ctx.Query("task")
	// Database_name := ctx.Query("database_name")
	// Type := ctx.Query("type")


}
