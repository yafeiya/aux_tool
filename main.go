package main

import (

	// "reflect"
	"backEnd/common"
	// "backEnd/common/response"
	"backEnd/controller"
	// "net/http"
	// "encoding/csv"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	// "os"
)

func main() {
	common.InitConfig()

	engine := gin.Default()
	engine.Use(common.CORS())
	
	// 主前端项目 (auxTool-frontEnd-main) - 根路径
	engine.Static("/assets", "./dist/assets")
	engine.StaticFile("/", "./dist/index.html")
	engine.StaticFile("/favicon.ico", "./dist/favicon.ico")
	engine.StaticFile("/aicard.png", "./dist/aicard.png")
	engine.StaticFile("/user.json", "./dist/user.json")
	engine.StaticFile("/user.txt", "./dist/user.txt")
	
	// AI工具前端项目 (AI_Tool_Dify_Front) - 子路径
	engine.Static("/ai-tool/css", "./AI_Tool_Dify_Front/dist/css")
	engine.Static("/ai-tool/js", "./AI_Tool_Dify_Front/dist/js")
	engine.StaticFile("/ai-tool/", "./AI_Tool_Dify_Front/dist/index.html")
	engine.StaticFile("/ai-tool/favicon.ico", "./AI_Tool_Dify_Front/dist/favicon.ico")
	engine.POST("/upload", controller.UploadCsvFile)
	engine.POST("/uploadReward", controller.UploadReward)
	engine.POST("/uploadActions", controller.UploadActions)
	engine.POST("/uploadLoss", controller.UploadLoss)
	engine.POST("/uploadResult", controller.UploadResult)
	engine.GET("/getprocessFile", controller.GetprocessFile)
	engine.GET("/getResultFile", controller.GetResultFile)
	engine.POST("/saveCanvas", controller.SaveCanvas)
	engine.POST("/saveCanvasPNG", controller.SaveCanvasPNG)
	engine.GET("/getDesignsById", controller.GetDesignsById)
	engine.POST("/runCanvas", controller.RunCanvas)
	engine.GET("/getExampleList", controller.GetExampleList)
	engine.GET("/updateExample", controller.UpdateExample)
	engine.GET("/deleteExample", controller.DeleteExample)
	engine.GET("/getCanvas", controller.GetCanvas)
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

			// 处理前端路由 - 根据路径分发
	engine.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/ai-tool") {
			// AI工具的路由交给AI工具前端处理
			c.File("./AI_Tool_Dify_Front/dist/index.html")
		} else {
			// 主项目的路由交给主前端处理
			c.File("./dist/index.html")
		}
	})
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
