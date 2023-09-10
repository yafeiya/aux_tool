package main

import (

	// "reflect"
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/controller"
	"backEnd/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	common.InitConfig()

	db := common.InitDB()
	// 与前端约定好字符串
	// id := 1
	pageKind := "database"
	released := "11"
	dataset_name := "11"
	Type := "11"
	rank := "11"
	character_type := "11"
	header := "11"
	data_path := "11"
	description := "11"
	task := "11"
	code := "......"
	lan := "Java"

	if pageKind == "database" {
		database := model.Database{
			Released:       released,
			Dataset_name:   dataset_name,
			Type:           Type,
			Rank:           rank,
			Character_type: character_type,
			Header:         header,
			Data_path:      data_path,
			Description:    description,
			Task:           task,
		}
		// 判重处理
		//pageKind、task、type、dataset_name
		db.Where("Task = ? and Type = ? and Dataset_name = ?", task, Type, dataset_name).First(&database)
		if database.Id != 0 {
			fmt.Println("该卡片已存在")
		} else {
			// 新增卡片
			db.Create(&database)
		}
	}
	if pageKind == "modelbase" {
		modelbase := model.Modelbase{
			Released:     released,
			Dataset_name: dataset_name,
			Type:         Type,
			Rank:         rank,
			Lan:          lan,
			Data_path:    data_path,
			Description:  description,
			Code:         code,
			Task:         task,
		}
		// 判重处理
		//pageKind、task、type、dataset_name
		db.Where("Task = ? and Type = ? and Dataset_name = ?", task, Type, dataset_name).First(&modelbase)
		if modelbase.Id != 0 {
			fmt.Println("该卡片已存在")
		} else {
			// 新增卡片
			db.Create(&modelbase)
		}
	}

	engine := gin.Default()
	engine.Use(common.CORS())

	engine.POST("/upload", func(ctx *gin.Context) {

		// 单文件
		Task := "task"
		Database_name := "database_name"
		Type := "type"

		file, _ := ctx.FormFile("file")
		fmt.Println(file.Filename)

		// 要创建的文件夹的路径
		folderPath := "./" + Type + "/" + Task + "/" + Database_name + "/"

		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			os.Mkdir(folderPath, 0755)
		}

		dst := folderPath + file.Filename
		// 上传文件至指定的完整文件路径
		ctx.SaveUploadedFile(file, dst)

		response.Success(ctx, nil, "success")

		//// 单文件
		//file, _ := c.FormFile("file")
		//log.Println(file.Filename)
		//
		//dst := "./" + file.Filename
		//// 上传文件至指定的完整文件路径
		//c.SaveUploadedFile(file, dst)
		//
		//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	engine.GET("/getPageMenu", controller.GetPageMenu)
	engine.GET("/getCard", controller.GetCard)
	engine.POST("/addPageMenuItem", controller.AddPageMenuItem)
	engine.POST("/login", controller.Login)
	engine.POST("/register", controller.Register)
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	if len(port) != 0 {
		panic(engine.Run(host + ":" + port))
	}
	err := engine.Run()
	if err != nil {
		panic(err)
	}
}
