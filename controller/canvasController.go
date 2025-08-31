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
	"time"
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
	id, _ := strconv.Atoi(ctx.Query("id"))
	newCellData := ctx.QueryArray("cell")

	// 添加所有必需的字段
	Design_name := ctx.Query("example_name")
	Rank := ctx.Query("Rank")
	State := ctx.Query("State")
	Task := ctx.Query("Task")
	Type := ctx.Query("Type")
	Dataset_name := ctx.Query("Dataset_name")
	Model_name := ctx.Query("Model_name")
	Train_state := ctx.Query("Train_state")
	Metics := ctx.Query("Metics")
	End_time := ctx.Query("end_time")

	// 打印接收到的参数进行调试
	fmt.Printf("接收到的参数: ID=%d, Name=%s, Rank=%s, Task=%s, Type=%s\n", id, Design_name, Rank, Task, Type)
	fmt.Printf("State=%s, Metics=%s, End_time=%s\n", State, Metics, End_time)
	network_num, _ := strconv.Atoi(ctx.Query("Network_num"))
	learning_rate, _ := strconv.ParseFloat(ctx.Query("Learning_rate"), 64)
	Act_function := ctx.Query("Act_function")
	radom_seed, _ := strconv.Atoi(ctx.Query("Radom_seed"))
	Optimizer := ctx.Query("Optimizer")
	batch_size, _ := strconv.Atoi(ctx.Query("Batch_size"))
	epoch_num, _ := strconv.Atoi(ctx.Query("Epoch_num"))
	explore_rate, _ := strconv.ParseFloat(ctx.Query("Explore_rate"), 64)
	decay_factor, _ := strconv.ParseFloat(ctx.Query("Decay_factor"), 64)

	fmt.Println("Dataset_name:", Dataset_name)
	// 创建命令
	cmd := exec.Command("python", "main.py", "12333", "131321212")
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
			//使用从URL参数获取的数据，而不是从JSON文件中读取
			// Design_name, Rank, Task, Type 已经从URL参数中获取
			fmt.Printf("使用URL参数: Design_name=%s, Rank=%s, Task=%s, Type=%s\n", Design_name, Rank, Task, Type)
			// 解析 JSON 数组数据
			fmt.Printf("接收到的cell数据数量: %d\n", len(newCellData))
			for idx, jsonStr := range newCellData {
				fmt.Printf("正在解析第%d个cell数据...\n", idx+1)
				if jsonStr == "" {
					fmt.Printf("跳过空的cell数据\n")
					continue
				}
				byteSlice := []byte(jsonStr) // Convert to []byte
				var item interface{}
				if err := json.Unmarshal(byteSlice, &item); err != nil {
					fmt.Printf("解析第%d个JSON数据失败: %v\n", idx+1, err)
					fmt.Printf("错误的JSON字符串长度: %d\n", len(jsonStr))
					// 不直接返回错误，继续处理其他数据
					continue
				}
				newCell = append(newCell, item)
			}
			fmt.Printf("成功解析的cell数据数量: %d\n", len(newCell))

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
				Start_time:    strconv.FormatInt(time.Now().UnixMilli(), 10),
				End_time:      "0",
			}
			// 判重处理 - 使用更好的判重逻辑
			var existingExample model.Example
			result := db.Where("Example_id = ?", id).First(&existingExample)

			if result.Error == nil {
				// 实例已存在，更新现有实例
				fmt.Printf("实例已存在，更新实例ID: %d\n", id)
				existingExample.Example_name = Design_name
				existingExample.Rank = Rank
				existingExample.State = "运行中"
				existingExample.Task = Task
				existingExample.Type = Type
				existingExample.Dataset_name = Dataset_name
				existingExample.Model_name = Model_name
				existingExample.Train_state = Train_state
				existingExample.Metics = "string"
				existingExample.Network_num = network_num
				existingExample.Learning_rate = learning_rate
				existingExample.Act_function = Act_function
				existingExample.Radom_seed = radom_seed
				existingExample.Optimizer = Optimizer
				existingExample.Batch_size = batch_size
				existingExample.Epoch_num = epoch_num
				existingExample.Explore_rate = explore_rate
				existingExample.Decay_factor = decay_factor
				existingExample.Start_time = strconv.FormatInt(time.Now().UnixMilli(), 10)
				existingExample.End_time = "0"

				result := db.Save(&existingExample)
				if result.Error != nil {
					fmt.Printf("更新实例失败: %v\n", result.Error)
					response.Response(ctx, http.StatusInternalServerError, 500, nil, "更新实例失败")
				} else {
					fmt.Printf("实例更新成功: %+v\n", existingExample)
					response.Success(ctx, nil, "实例已更新")
				}
			} else {
				// 新增实例
				fmt.Printf("创建新实例: %+v\n", example)
				result := db.Create(&example)
				if result.Error != nil {
					fmt.Printf("创建实例失败: %v\n", result.Error)
					response.Response(ctx, http.StatusInternalServerError, 500, nil, "创建实例失败")
				} else {
					fmt.Printf("实例创建成功: %+v\n", example)
					response.Success(ctx, nil, "实例创建成功")
				}
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

/*
接收训练配置参数
*/
type TrainingConfigRequest struct {
	SimulationPath string `json:"simulationPath"`
	ScenarioParams string `json:"scenarioParams"`
	ModelPath      string `json:"modelPath"`
	Timestamp      int64  `json:"timestamp"`
}

func ReceiveTrainingConfig(ctx *gin.Context) {
	var request TrainingConfigRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Printf("解析训练配置参数失败: %v\n", err)
		response.Response(ctx, http.StatusBadRequest, 400, nil, "参数解析失败")
		return
	}

	// 打印接收到的训练配置参数
	fmt.Println("=== 接收到训练配置参数 ===")
	fmt.Printf("仿真平台启动路径: %s\n", request.SimulationPath)
	fmt.Printf("仿真想定参数: %s\n", request.ScenarioParams)
	fmt.Printf("模型启动路径: %s\n", request.ModelPath)
	fmt.Printf("时间戳: %d\n", request.Timestamp)
	fmt.Println("========================")

	// 可以在这里添加保存到数据库或文件的逻辑
	// 例如保存到日志文件
	logFile, err := os.OpenFile("training_config.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("无法打开训练配置日志文件: %v\n", err)
	} else {
		defer logFile.Close()
		logger := log.New(logFile, "", log.LstdFlags)
		logger.Printf("训练配置参数 - 仿真平台启动路径: %s, 仿真想定参数: %s, 模型启动路径: %s, 时间戳: %d\n",
			request.SimulationPath, request.ScenarioParams, request.ModelPath, request.Timestamp)
	}

	// 返回成功响应
	response.Success(ctx, nil, "训练配置参数接收成功")
}
