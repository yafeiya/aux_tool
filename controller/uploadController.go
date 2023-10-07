package controller

import (
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"reflect"
	// "io"
)

type Record struct {
	Name   string `json:"name"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

// 上传文件
func UploadCsvFile(ctx *gin.Context) {
	// 路径参数

	Task := ctx.PostForm("task")
	Type := ctx.PostForm("type")
	Dataset_name := ctx.PostForm("name")
	time := ctx.PostForm("time")
	file, _ := ctx.FormFile("file")
	fmt.Println(file.Filename)
	fmt.Println(Dataset_name, Task, Type, time)

	// 要创建的文件夹的路径
	folderPath := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Dataset_name
	// 使用os.Mkdir创建文件夹
	err := os.Mkdir(folderPath, 0755) // 0755是文件夹的权限设置
	if err != nil {
		fmt.Println("创建文件夹失败:", err)
	}

	dst := folderPath + "/" + file.Filename
	dst_sql := "http://49.234.4.144:5173/" + Type + "/" + Task + "/" + Dataset_name + "/" + file.Filename
	// 上传文件至指定的完整文件路径
	ctx.SaveUploadedFile(file, dst)

	// 读取csv文件信息，获取行数、列数、数据类型
	numColumns, numRows, Types := GetCSVInfo(dst)
	result := CreateTable(Task, Type, Dataset_name, numColumns, numRows, Types, file.Filename, dst_sql, time)
	if result == "success" {
		response.Success(ctx, nil, "success")
	} else {
		response.Success(ctx, nil, "fail")
	}
}

// 将参数解析为map（因为可多选）
func DeleteCsvFile(ctx *gin.Context) {
	var records []Record
	db := common.InitDB()
	// 解析JSON数据到结构体切片中
	if err := ctx.BindJSON(&records); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		fmt.Println("解析json失败")
	}

	var flag int
	flag = 1
	for _, record := range records {
		Task := reflect.ValueOf(record).FieldByName("Task").String()
		Dataset_name := reflect.ValueOf(record).FieldByName("Dataset_name").String()
		Type := reflect.ValueOf(record).FieldByName("Type").String()
		Csv_name := reflect.ValueOf(record).FieldByName("Csv_name").String()

		// 在这里执行删除记录的操作
		datatable := []model.Datatable{}
		result := db.Where("Task = ? and Dataset_name = ? and Type = ? and Csv_name = ?", Task, Dataset_name, Type, Csv_name).Delete(&datatable)
		if result.RowsAffected == 0 {
			fmt.Println("删除失败")
			flag = 0
		} else {
			fmt.Println("删除成功")
		}
	}
	// 有一条或者多条记录删除失败
	if flag == 0 {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
	}
	response.Success(ctx, nil, "success")

}

// 下载文件
func DownloadCsvFile(ctx *gin.Context) {
	db := common.InitDB()
	// 路径参数
	Task := ctx.PostForm("task")
	Dataset_name := ctx.PostForm("dataset_name")
	Type := ctx.PostForm("type")
	Table_name := ctx.PostForm("table_name")
	fmt.Println(Task)
	fmt.Println(Dataset_name)
	fmt.Println(Type)
	fmt.Println(Table_name)
	datatables := []model.Datatable{}

	db.Where("Task = ? and Type = ? and Dataset_name = ? and Table_name = ?", Task, Type, Dataset_name, Table_name).Find(&datatables)
	if len(datatables) == 0 {
		fmt.Println("没找到")
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
	} else {
		filePath := datatables[0].Csv_path
		fmt.Println(filePath)
		response.Response(ctx, http.StatusOK, 404, gin.H{"url": filePath}, filePath)
		// file, err := os.Open(filePath)
		//		// if err != nil {
		//		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open CSV file"})
		//		// 	return
		//		// }
		//		// defer file.Close()
		//
		//		// // 设置HTTP响应头，指定文件类型和文件名
		//		// ctx.Header("Content-Type", "text/csv")
		//		// ctx.Header("Content-Disposition", "attachment; filename=" + filePath)
		//
		//		// // 将文件内容拷贝到HTTP响应主体中
		//		// _, err = io.Copy(ctx.Writer, file)
		//		// if err != nil {
		//		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not copy file to response"})
		//		// 	return
		//		// }
	}

}
