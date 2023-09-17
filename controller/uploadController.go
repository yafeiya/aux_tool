package controller

import (
	"backEnd/common/response"
	"backEnd/model"
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"backEnd/common"
	// "io"
)

// 上传文件
func UploadFile(ctx *gin.Context) {
		// 路径参数
		
		Task := ctx.PostForm("task")
		Type  := ctx.PostForm("type")
		Dataset_name := ctx.PostForm("name")
		file, _ := ctx.FormFile("file")
		fmt.Println(file.Filename)
		fmt.Println(Dataset_name,Task,Type)

		// 要创建的文件夹的路径
		folderPath := "./" + Type + "/" + Task + "/" + Dataset_name 

		// 使用os.Mkdir创建文件夹
		err := os.Mkdir(folderPath, 0755) // 0755是文件夹的权限设置
		if err != nil {
			fmt.Println("创建文件夹失败:", err)
		}
		
		dst := folderPath + "/"+ file.Filename
		// 上传文件至指定的完整文件路径
		ctx.SaveUploadedFile(file, dst)

		// 读取csv文件信息，获取行数、列数、数据类型
		numColumns, numRows, Types := GetCSVInfo(dst)
		result := CreateTable(Task, Type, Dataset_name, numColumns, numRows, Types, file.Filename, dst)
		
		fmt.Println("创建文件夹失败222222222222:", result)

		if(result == "success"){
			response.Success(ctx, nil, "success")
		}else{
			response.Success(ctx, nil, "fail")
		}
}

// 下载文件
func DownloadFile(ctx *gin.Context) {
	db := common.InitDB()
	// 路径参数
	Task := ctx.PostForm("Task")
	Dataset_name := ctx.PostForm("dataset_name")
	Type := ctx.PostForm("Type")
	Table_name := ctx.PostForm("csvName")
	fmt.Println(Task)
	fmt.Println(Dataset_name)
	fmt.Println(Type)
	fmt.Println(Table_name)
	datatables := []model.Datatable{}

	db.Where("Task = ? and Type = ? and Dataset_name = ? and Table_name = ?",Task, Type, Dataset_name, Table_name).Find(&datatables)
	if len(datatables) == 0 {
		fmt.Println("没找到")
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
	} else {
		filePath := datatables[0].Csv_path
		fmt.Println(filePath)
		response.Response(ctx, http.StatusOK, 404, gin.H{"url": filePath}, filePath)
		// file, err := os.Open(filePath)
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open CSV file"})
		// 	return
		// }
		// defer file.Close()
	
		// // 设置HTTP响应头，指定文件类型和文件名
		// ctx.Header("Content-Type", "text/csv")
		// ctx.Header("Content-Disposition", "attachment; filename=" + filePath)
	
		// // 将文件内容拷贝到HTTP响应主体中
		// _, err = io.Copy(ctx.Writer, file)
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not copy file to response"})
		// 	return
		// }
	}

}
