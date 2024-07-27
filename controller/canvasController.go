package controller

import (
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/model"
	"encoding/json"
	"io/ioutil"
	"strconv"

	// "backEnd/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "strconv"
	// "strings"
)

/*
前端保存按钮
*/

type CellsRequest struct {
	Id    string        `json:"id"` // 添加 name 字段
	Cells []interface{} `json:"cells"`
}

func SaveCanvas(ctx *gin.Context) {
	var request CellsRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 将 ID 字符串转换为整数
	idInt, err := strconv.Atoi(request.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// 打印接收到的 ID
	fmt.Printf("Canvas ID: %d\n", idInt)

	// 读取JSON文件
	jsonData, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	// 解析JSON数据
	var data model.Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println("解析JSON数据失败:", err)
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}

	// 遍历列表
	for i := range data.Designs {
		if data.Designs[i].Id == idInt {
			fmt.Println("要覆盖的design的id为:", idInt)
			//var newCell []interface{}
			data.Designs[i].Cell = request.Cells
			// 更新JSON文件
			updatedJSON, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				response.Response(ctx, http.StatusOK, 404, nil, "fail")
				return
			}

			if err := ioutil.WriteFile("./data/data.json", updatedJSON, 0644); err != nil {
				response.Response(ctx, http.StatusOK, 404, nil, "fail")
				return
			}
			response.Success(ctx, nil, "success")
			return
		}
	}

	fmt.Println("未找到指定的design")
	response.Response(ctx, http.StatusOK, 404, nil, "fail")

}

// /*
// 前端运行按钮
// */
func RunCanvas(ctx *gin.Context) {
	fmt.Println("RunCanvas")
	Start_time := ctx.Query("start_time")

	id, _ := strconv.Atoi(ctx.Query("id"))
	newCellData := ctx.QueryArray("cell")
	Dataset_url := ctx.Query("dataset_url")
	fmt.Println("Dataset_url", Dataset_url)
	// fmt.Println("newCellData", newCellData)
	fmt.Println("id", id)
	fmt.Println("time", Start_time)

	fmt.Println("start_time", Start_time)

	// 读取JSON文件
	jsonData, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}

	// 解析JSON数据
	var data model.Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println("解析JSON数据失败:", err)
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}

	// 遍历列表
	for i := range data.Designs {
		if data.Designs[i].Id == id {
			fmt.Println("要覆盖的design的id为:", id)
			var newCell []interface{}
			//获取对应案例名称
			Design_name := data.Designs[i].Dataset_name
			Rank := data.Designs[i].Rank
			fmt.Println("rank", Rank)
			// 解析 JSON 数组数据
			for _, jsonStr := range newCellData {
				byteSlice := []byte(jsonStr) // Convert to []byte
				var item interface{}
				if err := json.Unmarshal(byteSlice, &item); err != nil {
					fmt.Println("解析 JSON 数组数据fail:", err)
					response.Response(ctx, http.StatusOK, 404, nil, "fail")
					return
				}
				newCell = append(newCell, item)
			}
			// fmt.Println("newCell:", newCell)
			data.Designs[i].Cell = newCell

			// 更新JSON文件
			updatedJSON, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				response.Response(ctx, http.StatusOK, 404, nil, "fail")
				return
			}

			if err := ioutil.WriteFile("./data/data.json", updatedJSON, 0644); err != nil {
				response.Response(ctx, http.StatusOK, 404, nil, "fail")
				return
			}

			// example添加
			db := common.InitDB()

			example := model.Example{
				Example_name: Design_name,
				Rank:         Rank,
				State:        "运行中",
				Cpu_num:      4,
				Gpu_num:      1,
				Post_data:    0,
				Dataset_url:  Dataset_url,
				Model_name:   "",
				Model_type:   "",
				Epoch_num:    "200e",
				Loss:         "loss",
				Optimizer:    "optimizer",
				Decay:        "decay",
				Evaluation:   "evaluation",
				Model_url:    "model_url",
				Memory:       "2000M",
				Start_time:   Start_time,
				End_time:     "0",
			}
			// 判重处理
			//pageKind、task、type、dataset_name

			db.Where("Example_name = ?", Design_name).Find(&example)
			fmt.Println("Example_name", Design_name)
			fmt.Println("example_id", example.Id)
			if example.Id != 0 {
				fmt.Println("该实例已存在")
				response.Response(ctx, http.StatusOK, 404, nil, "The example already exists")
			} else {
				// 新增
				db.Create(&example)
				response.Success(ctx, nil, "success")
			}

			return
		}
	}

	fmt.Println("未找到指定的design")
	response.Response(ctx, http.StatusOK, 404, nil, "fail")
}

/*
前端保存png
*/
func SaveCanvasPNG(ctx *gin.Context) {
	fmt.Println("SaveCanvasPNG")
	Id := ctx.PostForm("id")
	Type := ctx.Query("type")
	Task := ctx.Query("task")
	image, _ := ctx.FormFile("image")
	// 要创建的文件夹的路径
	folderPath := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id
	// 使用os.Mkdir创建文件夹
	err := os.Mkdir(folderPath, 0755) // 0755是文件夹的权限设置
	if err != nil {
		fmt.Println("创建文件夹失败:", err)
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}

	dst := folderPath + "/image.png"
	// 上传文件至指定的完整文件路径
	ctx.SaveUploadedFile(image, dst)

	response.Success(ctx, nil, "success")

}

/*
前端上传reward.txt
*/
func UploadReward(ctx *gin.Context) {
	fmt.Println("SaveCanvasPNG")
	Id := ctx.PostForm("id")
	Type := ctx.PostForm("type")
	Task := ctx.PostForm("task")
	reward, _ := ctx.FormFile("file")
	folderPath := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id
	// 使用os.Mkdir创建文件夹
	err := os.Mkdir(folderPath, 0755) // 0755是文件夹的权限设置
	if err != nil {
		fmt.Println("创建文件夹失败:", err)
	}
	fmt.Println("11111111111111111111111:")

	dst := folderPath + "/" + "reward.txt"
	erro := os.Remove(dst)
	if erro != nil {
		fmt.Println("delete fail")
	}
	// 上传文件至指定的完整文件路径1
	ctx.SaveUploadedFile(reward, dst)

	response.Success(ctx, nil, "success")
}

/*
前端上传actions.json
*/
func UploadActions(ctx *gin.Context) {
	fmt.Println("UploadActions")
	Id := ctx.PostForm("id")

	Type := ctx.PostForm("type")
	Task := ctx.PostForm("task")
	actions, _ := ctx.FormFile("file")
	folderPath := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id
	// 使用os.Mkdir创建文件夹
	err := os.Mkdir(folderPath, 0755) // 0755是文件夹的权限设置
	if err != nil {
		fmt.Println("创建文件夹失败:", err)
	}
	fmt.Println("11111111111111111111111:")

	dst := folderPath + "/" + "actions.json"
	erro := os.Remove(dst)
	if erro != nil {
		fmt.Println("delete fail")
	}
	// 上传文件至指定的完整文件路径1
	ctx.SaveUploadedFile(actions, dst)

	response.Success(ctx, nil, "success")
}

/*
前端上传loss.csv
*/
func UploadLoss(ctx *gin.Context) {
	fmt.Println("UploadLoss")
	Id := ctx.PostForm("id")
	Type := ctx.PostForm("type")
	Task := ctx.PostForm("task")
	loss, _ := ctx.FormFile("file")
	// 要创建的文件夹的路径
	folderPath := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id
	// 使用os.Mkdir创建文件夹
	err := os.Mkdir(folderPath, 0755) // 0755是文件夹的权限设置
	if err != nil {
		fmt.Println("创建文件夹失败:", err)
	}
	fmt.Println("11111111111111111111111:")

	dst := folderPath + "/" + "loss.csv"
	erro := os.Remove(dst)
	if erro != nil {
		fmt.Println("delete fail")
	}
	// 上传文件至指定的完整文件路径1
	ctx.SaveUploadedFile(loss, dst)

	response.Success(ctx, nil, "success")
}

func GetprocessFile(ctx *gin.Context) {
	fmt.Println("GetprocessFile", ctx)
	Id := ctx.Query("id")
	fmt.Println("GetprocessFile_id", Id)
	Type := ctx.Query("type")
	Task := ctx.Query("task")
	//processFile为reward.txt \ actions.json \ loss.csv
	processFile := ctx.Query("processFile")
	fmt.Println("processFile", processFile)
	file_path := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id + "/" + processFile
	Info, err := GetTrainingProcessFileInfo(file_path)
	fmt.Println("file_path", file_path)
	if err != nil {
		fmt.Println(err)
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	response.Success(ctx, gin.H{"Info": Info}, "success")
}

func GetDesignsById(ctx *gin.Context) {
	fmt.Println("GetDesignsById")
	Id, _ := strconv.Atoi(ctx.Query("id"))
	jsonData, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	var data model.Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	// 否则，筛选出符合条件的用户
	for _, design := range data.Designs {
		if design.Id == Id {
			response.Success(ctx, gin.H{"design": design}, "success")
			return
		}
	}
	response.Response(ctx, http.StatusOK, 404, nil, "fail")
}

func GetPNG(ctx *gin.Context) {
	fmt.Println("GetPNG")
	Id := ctx.PostForm("id")
	Type := ctx.Query("type")
	Task := ctx.Query("task")
	//目录下的文件名（带拓展名）
	file_path := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id + "/" + "image.png"

	image, err := GetTrainingProcessFileInfo(file_path)
	if err != nil {
		fmt.Println(err)
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	response.Success(ctx, gin.H{"image": image}, "success")
}

func GetCanvas(ctx *gin.Context) {
	fmt.Println("getCanvas000000000000")
	Id, _ := strconv.Atoi(ctx.Query("id"))
	jsonData, err := ioutil.ReadFile("./data/data.json")

	if err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	var data model.Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	for i := range data.Designs {
		if data.Designs[i].Id == Id {
			response.Success(ctx, gin.H{"cell": data.Designs[i].Cell}, "success")
			return
		}
	}
}
