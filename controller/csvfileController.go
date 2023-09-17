package controller

import (
	"backEnd/model"
	"net/http"
	"backEnd/common"
	"backEnd/common/response"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"fmt"
	"unicode/utf8"
	"encoding/csv"
)

// 获取csv文件属性
func GetCSVInfo(csv_path string)(int, int, []string) {
	// 打开CSV文件
	file, error := os.Open(csv_path)
	if error != nil {
		fmt.Println("无法打开CSV文件:", error)
	}
	defer file.Close()

	// 创建CSV读取器
	reader := csv.NewReader(file)

	// 读取CSV行
	lines, error := reader.ReadAll()
	if error != nil {
		fmt.Println("读取CSV文件失败:", error)
	}

	// 获取列数和行数
	numColumns := len(lines[0])
	numRows := len(lines)

	// 创建一个切片用于保存每列的数据类型
	columnTypes := make([]string, numColumns)

	// 遍历第二行以识别每列的数据类型
	for columnIndex := 0; columnIndex < numColumns; columnIndex++ {
		// 默认为字符串类型
		columnType := "字符串"

		// 尝试将第二行数据识别为rune
		if utf8.RuneCountInString(lines[1][columnIndex]) == 1 {
			columnType = "char"
		}

		// 尝试将第二行数据转换为int
		_, error := strconv.Atoi(lines[1][columnIndex])
		if error == nil {
			columnType = "int"
		}

		// 尝试将第二行数据转换为float64
		_, error = strconv.ParseFloat(lines[1][columnIndex], 64)
		if error == nil {
			columnType = "float"
		}

		// 将数据类型保存到切片中
		columnTypes[columnIndex] = columnType
	}

    // 创建一个映射来保存不重复的字符串元素
    uniqueStrings := make(map[string]bool)

    // 将元素添加到映射中
    for _, element := range columnTypes {
        uniqueStrings[element] = true
    }

	var Types []string
	for uniqueElement := range uniqueStrings {
		Types = append(Types, uniqueElement)
    }

	// // 输出列数
	// fmt.Printf("CSV文件有 %d 列\n", numColumns)

	// // 输出每列的数据类型
	// for columnIndex, columnType := range columnTypes {
	// 	fmt.Printf("  列 %d: %s\n", columnIndex+1, columnType)
	// }

	// // 输出行数
	// fmt.Printf("CSV文件有 %d 行\n", numRows)

	return numColumns, numRows, Types
}

// 获取某数据集的csv表
func GetTable(ctx *gin.Context) {
	// 参数
	Task := ctx.Query("Task")
	Dataset_name := ctx.Query("Dataset_name")
	Type := ctx.Query("Type")

	db := common.InitDB()

	datatables := []model.Datatable{}
	db.Where("Task = ? and Type = ? and Dataset_name = ?", Task, Type, Dataset_name).Find(&datatables)
	if len(datatables) == 0 {
		fmt.Println("卡片")
		response.Response(ctx, http.StatusOK, 404, nil, "No corresponding card found")
	} else {
		fmt.Println(datatables)
		response.Success(ctx, gin.H{"datatables": datatables}, "success")
	}
}

// 获取某数据集的csv表
func DeleteTable(ctx *gin.Context) {
	// 参数
	Task := ctx.Query("任务1")
	Database_name := ctx.Query("波士顿房价数据集")
	Type := ctx.Query("数值数据集")
	Table_name := ctx.Query("动作表")

	db := common.InitDB()

	datatable := []model.Datatable{}
	result := db.Where("Task = ? and Type = ? and Dataset_name = ? and Table_name = ?", Task, Type, Database_name, Table_name).Delete(&datatable)
	if result.RowsAffected == 0 {
		fmt.Println("删除失败")
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
	} else {
		fmt.Println("删除成功")
		response.Success(ctx, nil, "success")
	}
}

// 添加csv表
func CreateTable(Task string, Type string, Dataset_name string, numColumns int, numRows int, Types []string, csv_name string, csv_path string,time string)(string) {
	db := common.InitDB()
	types := ""
	for _, columnType := range Types {
		types = types + columnType
	}
	datatable := model.Datatable{
		Type:       	Type,
		Task:   		Task,
		Dataset_name:   Dataset_name,
		Table_name:     csv_name,
		Header_num: 	uint(numRows),
		Data_len:       uint(numColumns),
		Data_type:      types,
		Csv_path:		csv_path,
		Time:			time,
	}
	// 判重处理
	db.Where("Task = ? and Type = ? and Dataset_name = ? and Table_name = ?", Task, Type, Dataset_name, csv_name).First(&datatable)
	if datatable.Id != 0 {
		fmt.Println("该卡片已存在")
		// response.Response(ctx, http.StatusOK, 404, nil, "The card already exists")
		return "fail"
	} else {
		// 新增卡片
		db.Create(&datatable)
		fmt.Println("创建表")
		// response.Success(ctx, nil, "success")
		return "success"
	}
}

// 获取表信息

func GetCsvData(ctx *gin.Context) {
	Dataset_name := ctx.Query("Dataset_name")
	fmt.Println("需要查询的数据集名称：",Dataset_name)

	db := common.InitDB()

	datatables := []model.Datatable{}
	db.Where("Dataset_name = ?", Dataset_name).Find(&datatables)
	if len(datatables) == 0 {
		fmt.Println("未找到数据集下的csv文件")
		response.Response(ctx, http.StatusOK, 404, nil, "No corresponding card found")
	} else {
		fmt.Println(datatables)
		response.Success(ctx, gin.H{"datatables": datatables}, "success")
	}

}