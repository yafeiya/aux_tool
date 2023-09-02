package main

import (
	"fmt"
	_ "fmt"

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

	//test
	name := "sk"
	pwd := "111"
	user := model.User{
		UserName: name,
		PassWord: pwd,
	}
	// 判重处理
	db.Where("UserName = ?", name).First(&user)
	if user.Id != 0 {
		fmt.Println("该用户已存在")
	} else {
		// 新增该注册用户
		db.Create(&user)
	}



	engine := gin.Default()
	engine.Use(common.CORS())
	engine.GET("/getPageMenu", controller.GetPageMenu)
	engine.POST("/addPageMenuItem", controller.AddPageMenuItem)
	engine.POST("/postRegist", controller.Register)
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
