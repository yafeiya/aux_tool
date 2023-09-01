package common

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Println("读取当前路径错误：" + err.Error())
	}
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
