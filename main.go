package main

import (

	// "reflect"
	"backEnd/common"
	"backEnd/controller"
	// "os"
	// "strconv"
	// "backEnd/model"
	"fmt"
	// "log"
	// "unicode/utf8"
	"net/http"
	// "encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	common.InitConfig()

	// filename := "./data/config.json"

	// // 创建新的模型实例
	// newData := model.DataMenu{
	// 	Name:  "newData",
	// 	Title: "新数据",
	// 	Icon:  "ios-new-icon",
	// 	Children: []model.Children{
	// 		{Name: "newChild1", Title: "新子项1"},
	// 		{Name: "newChild2", Title: "新子项2"},
	// 	},
	// }
	// objectType := "design"
	// er := model.AddMenu(objectType, filename, newData)
	// if er != nil {
	// 	fmt.Println(er)
	// }

	engine := gin.Default()
	engine.Use(common.CORS())
	engine.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		dst := "./" + file.Filename
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	engine.GET("/getPageMenu", controller.GetPageMenu)
	engine.POST("/getCard", controller.GetCard)
	engine.POST("/addPageMenuItem", controller.AddPageMenuItem)
	engine.POST("/login", controller.Login)
	engine.POST("/register", controller.Register)
	engine.POST("/download", controller.DownloadFile)
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
