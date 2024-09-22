package controller

import (
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/model"
	"backEnd/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
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
	Start_time := ctx.Query("start_time")
	id, _ := strconv.Atoi(ctx.Query("id"))
	newCellData := ctx.QueryArray("cell")
	Dataset_name := ctx.Query("Dataset_name")
	Model_name := ctx.Query("Model_name")
	Train_state := ctx.Query("Train_state")
	network_num, _ := strconv.Atoi(ctx.Query("Network_num"))
	learning_rate, _ := strconv.ParseFloat(ctx.Query("Learning_rate"), 64)
	Act_function := ctx.Query("Act_function")
	radom_seed, _ := strconv.Atoi(ctx.Query("Radom_seed"))
	Optimizer := ctx.Query("Optimizer")
	batch_size, _ := strconv.Atoi(ctx.Query("Batch"))
	epoch_num, _ := strconv.Atoi(ctx.Query("Epoch_num"))
	explore_rate, _ := strconv.ParseFloat(ctx.Query("Explore_rate"), 64)
	decay_factor, _ := strconv.ParseFloat(ctx.Query("Decay_factor"), 64)

	fmt.Println("Dataset_name:", Dataset_name)
	// 创建命令
	cmd := exec.Command("python", "main.py", "12333", "131321212") // 或者使用 "python3" 根据你的环境	// 获取命令输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to run command: %v", err),
		})
		return
	}
	// 打开日志文件（如果不存在则创建）
	logFile, err := os.OpenFile("command_output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Failed to open log file: %v\n", err)
		ctx.String(http.StatusInternalServerError, "Failed to open log file")
		return
	}
	defer logFile.Close()

	// 创建日志记录器
	logger := log.New(logFile, "", log.LstdFlags)

	// 写入输出到日志文件
	logger.Printf("Command Output:\n%s\n", string(output))

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
			Task := data.Designs[i].Task
			Type := data.Designs[i].Type
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
				Example_id:    id,
				Example_name:  Design_name,
				Rank:          Rank,
				State:         "运行中",
				Task:          Task,
				Type:          Type,
				Dataset_name:  Dataset_name,
				Model_name:    Model_name,
				Train_state:   Train_state,
				Metics:        "string",
				Network_num:   network_num,
				Learning_rate: learning_rate,
				Act_function:  Act_function,
				Radom_seed:    radom_seed,
				Optimizer:     Optimizer,
				Batch_size:    batch_size,
				Epoch_num:     epoch_num,
				Explore_rate:  explore_rate,
				Decay_factor:  decay_factor,
				Start_time:    Start_time,
				End_time:      "0",
			}
			// 判重处理
			//pageKind、task、type、dataset_name
			db.Where("Example_name = ?", Design_name).Find(&example)
			if example.Id != 0 {
				fmt.Println("该实例已存在")
				// 更新现有实例的字段
				example.Example_id = id
				example.Rank = Rank
				example.State = "运行中"
				example.Task = Task
				example.Type = Type
				example.Dataset_name = Dataset_name
				example.Model_name = Model_name
				example.Train_state = Train_state
				example.Metics = "string"
				example.Network_num = network_num
				example.Learning_rate = learning_rate
				example.Act_function = Act_function
				example.Radom_seed = radom_seed
				example.Optimizer = Optimizer
				example.Batch_size = batch_size
				example.Epoch_num = epoch_num
				example.Explore_rate = explore_rate
				example.Decay_factor = decay_factor
				example.Start_time = Start_time
				example.End_time = "0"

				db.Save(&example) // 保存更新
				db.Save(&example) // 保存更新
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
	Id := ctx.PostForm("id")
	Type := ctx.PostForm("type")
	Task := ctx.PostForm("task")
	image, _ := ctx.FormFile("image")
	// 要创建的文件夹的路径
	folderPath := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id
	// 使用os.Mkdir创建文件夹
	err := os.Mkdir(folderPath, 0755) // 0755是文件夹的权限设置
	if err != nil {
		fmt.Println("创建文件夹失败:", err)
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
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

	dst := folderPath + "/" + "loss.csv"
	erro := os.Remove(dst)
	if erro != nil {
		fmt.Println("delete fail")
	}
	// 上传文件至指定的完整文件路径1
	ctx.SaveUploadedFile(loss, dst)

	response.Success(ctx, nil, "success")
}
func UploadResult(ctx *gin.Context) {
	fmt.Println("UploadResult")
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

	dst := folderPath + "/" + "result.json"
	erro := os.Remove(dst)
	if erro != nil {
		fmt.Println("delete fail")
	}
	// 上传文件至指定的完整文件路径1
	ctx.SaveUploadedFile(loss, dst)

	response.Success(ctx, nil, "success")
}
func GetprocessFile(ctx *gin.Context) {

	Id := ctx.Query("id")

	Type := ctx.Query("type")
	Task := ctx.Query("task")
	//processFile为reward.txt \ actions.json \ loss.csv
	processFile := ctx.Query("processFile")

	file_path := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id + "/" + processFile
	Info, err := GetTrainingProcessFileInfo(file_path)

	if err != nil {
		fmt.Println(err)
		response.Response(ctx, http.StatusOK, 404, nil, "fail")
		return
	}
	response.Success(ctx, gin.H{"Info": Info}, "success")
}
func GetResultFile(ctx *gin.Context) {

	Id := ctx.Query("id")
	Type := ctx.Query("type")
	Task := ctx.Query("task")
	processFile := ctx.Query("processFile")
	file_path := "./auxTool-frontEnd-main/" + Type + "/" + Task + "/" + Id + "/" + processFile
	Info, err := utils.LoadJson(file_path)

	if !err {
		fmt.Println("加载文件失败")
		return
	}
	response.Success(ctx, gin.H{"Info": Info}, "success")
}

func GetDesignsById(ctx *gin.Context) {

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
