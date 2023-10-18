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
	engine.POST("/uploadReward", controller.UploadReward)
	engine.POST("/uploadActions", controller.UploadActions)
	engine.POST("/uploadLoss", controller.UploadLoss)
	engine.POST("/getprocessFile", controller.GetprocessFile)

	engine.POST("/uploadModelFile", controller.UploadModelFile)
	engine.POST("/uploadModelPNGFile", controller.UploadModelPNGFile)
	engine.POST("/xmlinfo", controller.OutPutXml)
	engine.GET("/getPageMenu", controller.GetPageMenu)
	engine.GET("/getCsv", controller.GetCsvData)
	engine.GET("/updataCard", controller.UpdataCard)
	engine.GET("/updateDesignCard", controller.UpdateDesignCard)
	engine.GET("/createCard", controller.CreateCard)
	engine.GET("/deleteCard", controller.DeleteCard)
	engine.GET("/deleteTable", controller.DeleteTable)
	engine.GET("/deleteDesignCard", controller.DeleteDesignCard)
	engine.GET("/getCard", controller.GetCard)
	engine.GET("/getDesignCards", controller.GetDesignCards)
	engine.GET("/addPageMenuItem", controller.AddPageMenuItem)
	engine.GET("/addDesignCard", controller.AddDesignCard)
	engine.GET("/addDefinefunctionCard", controller.AddDefinefunctionCard)
	engine.GET("/getDefinefunctionCards", controller.GetDefinefunctionCards)
	engine.GET("/updateDefinefunctionCard", controller.UpdateDefinefunctionCard)
	engine.GET("/deleteDefinefunctionCard", controller.DeleteDefinefunctionCard)
	engine.GET("/deletePageMenuItem", controller.DeletePageMenuItem)
	engine.POST("/downloadCsvFile", controller.DownloadCsvFile)
	engine.POST("/downloadModelFile", controller.DownloadModelFile)
	engine.POST("/getModelImage", controller.GetModelImage)
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
