package main

import (
	"fmt"
	_ "fmt"

	// "reflect"
	"backEnd/common"
	"backEnd/controller"
	"backEnd/model"

	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	// db.Create(&model.User{Id: 1, UserName: "123", PassWord: "1234"})
	var user model.User
	db.First(&user)
	fmt.Println(user)
	engine := gin.Default()
	engine.Use(common.CORS())
	engine.GET("/getPageMenu", controller.GetPageMenu)
	engine.POST("/addPageMenuItem", controller.AddPageMenuItem)
	err := engine.Run("192.168.0.103:2345")
	if err != nil {
		panic(err)
	}
}
