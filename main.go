package main

import (

	// "reflect"
	"backEnd/common"
	// "backEnd/common/response"
	"backEnd/controller"
	// "net/http"
	// "encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	// "os"
)

func main() {
	common.InitConfig()

	engine := gin.Default()
	engine.Use(common.CORS())
	engine.POST("/upload", controller.UploadCsvFile)
	engine.GET("/getPageMenu", controller.GetPageMenu)
	engine.GET("/getCsv", controller.GetCsvData)
	engine.GET("/updataCard", controller.UpdataCard)
	engine.GET("/createCard", controller.CreateCard)
	engine.GET("/deleteCard", controller.DeleteCard)
	engine.GET("/getCard", controller.GetCard)
	engine.GET("/addPageMenuItem", controller.AddPageMenuItem)
	engine.GET("/deletePageMenuItem", controller.DeletePageMenuItem)
	engine.POST("/downloadCsvFile", controller.DownloadCsvFile)
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
