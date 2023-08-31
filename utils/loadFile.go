package utils

// 加载json文件
import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var mutex sync.Mutex

func LoadJson(filePath string) (map[string]interface{}, bool) {
	mutex.Lock()
	file, err := os.ReadFile(filePath)
	mutex.Unlock()
	if err != nil {
		fmt.Println("读取文件错误")
		return nil, false
	}
	var menuList map[string]interface{}
	err = json.Unmarshal(file, &menuList)
	if err != nil {
		fmt.Println("解码json错误")
		fmt.Println(err)
		return nil, false
	}
	// fmt.Printf("value: %+v\n", newMenu)
	// var menu map[string]interface{}
	// menu := menuList[pageName]
	// fmt.Println("menu: ")
	// fmt.Println(menu)
	return menuList, true
	// fmt.Println(reflect.TypeOf(newMenu))
}
