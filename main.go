package main

import (

	// "reflect"
	"backEnd/common"
	"backEnd/common/response"
	"backEnd/controller"
	"backEnd/model"
	"fmt"
	// "net/http"
	// "encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	common.InitConfig()

	Task := "任务1"
	Dataset_name := "波士顿房价数据"
	Type := "数值数据集"
	Table_name := "动作表"

	db := common.InitDB()

	datatable := []model.Datatable{}
	result := db.Where("Task = ? and Type = ? and Dataset_name = ? and Table_name = ?", Task, Type, Dataset_name, Table_name).Delete(&datatable)
	if result.RowsAffected == 0 {
		fmt.Println("删除失败")
	} else {
		fmt.Println("删除成功")
	}

	engine := gin.Default()
	engine.Use(common.CORS())

	engine.POST("/upload", func(ctx *gin.Context) {

		// 单文件
		Task := "task"
		Database_name := "database_name"
		Type := "type"
		// fmt.Println("2222",ctx.Request.Header)
		// fmt.Println("2222",ctx.Request.Host)
		// etc
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
