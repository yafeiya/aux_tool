package main

import (
	"fmt"

	// "reflect"
	"backEnd/common"
	"backEnd/controller"
	"backEnd/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	common.InitConfig()
	db := common.InitDB()
	// db.Create(&model.User{Id: 1, UserName: "123", PassWord: "1234"})
	var user model.User
	db.First(&user)
	fmt.Println(user)
	engine := gin.Default()
	engine.Use(common.CORS())
	engine.GET("/getPageMenu", controller.GetPageMenu)
	engine.POST("/addPageMenuItem", controller.AddPageMenuItem)
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
