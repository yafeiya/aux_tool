package main

import (

	// "reflect"
	"backEnd/common"
	"backEnd/controller"
	"backEnd/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	engine.GET("/getPageMenu", controller.GetPageMenu)
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
