package controller

import (

	"encoding/json"

	"fmt"
	"os"
	"strconv"

	"bufio"
	"encoding/csv"
	"io/ioutil"
	"strings"
)

type CSVData struct {
	Column1 []string `json:"column1"`
	Column2 []string `json:"column2"`
	Column3 []string `json:"column3"`
	Column4 []string `json:"column4"`
}

type JSONData struct {
	Key []map[string]interface{} `json:"key"`
}

// 获取训练过程文件属性txt/csv/json
func GetTrainingProcessFileInfo(filePath string) (string, error) {

	if strings.HasSuffix(filePath, ".txt") {
		// 处理txt文件
		result, err := processTxtFile(filePath)
		if err != nil {
			return "", nil
		}
		return result, nil
	} else if strings.HasSuffix(filePath, ".csv") {
		// 处理csv文件
		result, err := processCsvFile(filePath)
		if err != nil {
			return "", err
		}
		return result, nil
	} else if strings.HasSuffix(filePath, ".json") {
		// 处理json文件
		result, err := processJsonFile(filePath)
		if err != nil {
			return "", err
		}
		return result, nil
	} else {
		return " ", fmt.Errorf("不支持的文件类型")
	}


}

func processTxtFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return "", err
	}
	defer file.Close()

	// 创建一个切片来存储第三列数据
	var thirdColumn []float64

	// 创建一个Scanner来逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 拆分每一行的数据
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) >= 3 {
			// 将第三列数据解析为float64并添加到切片中
			value, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
			
				continue
			}
			thirdColumn = append(thirdColumn, value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时出错:", err)
		return "", err
	}

	// 创建一个JSON对象
	jsonData := map[string][]float64{"reward": thirdColumn}

	// 将JSON对象编码为JSON字符串
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("JSON编码出错:", err)
		return "", err
	}
	return string(jsonBytes), nil
}

func processCsvFile(filePath string) (string, error) {
	// 打开CSV文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return "", err
	}
	defer file.Close()

	// 创建一个CSV Reader
	reader := csv.NewReader(file)

	// 读取CSV文件的第一行，该行包含字段名称
	header, err := reader.Read()
	if err != nil {
		fmt.Println("读取CSV文件头部出错:", err)
		return "", err
	}

	// 创建一个map来存储字段名称和对应的数据切片
	data := make(map[string][]float64)

	for i := 0; i < len(header); i++ {
		data[header[i]] = make([]float64, 0)
	}

	for {
		// 读取下一行数据
		record, err := reader.Read()

		if err != nil {
			break // 文件读取结束
		}

		for i := 0; i < len(header); i++ {
			value, _ := strconv.ParseFloat(record[i], 64)
			data[header[i]] = append(data[header[i]], value)
		}
	}

	// 创建包含"loss"字段的JSON对象
	result := map[string]map[string][]float64{"loss": data}

	// 将JSON对象编码为JSON字符串
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println("JSON编码出错:", err)
		return "", err
	}

	return string(jsonBytes), nil
}

func processJsonFile(filePath string) (string, error) {
	// 读取JSON文件
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return "", err
	}

	// 解析JSON数据
	var data JSONData
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		fmt.Println("JSON解析出错:", err)
		return "", err
	}

	// 创建映射以将发送者和目标ID映射到整数
	senderMap := make(map[string]int)
	targetMap := make(map[string]int)
	senderCount := 0
	targetCount := 0

	// 创建结果切片
	result := [][]int{}

	// 遍历JSON数据
	for _, item := range data.Key {
		if subItem, ok := item["TgtAssignAffirm"].(map[string]interface{}); ok {
			if bIncpt, ok := subItem["bIncpt"].(bool); ok && bIncpt {
				strSenderId := subItem["strSenderId"].(string)
				strTgrId := subItem["strTgrId"].(string)

				// 检查发送者是否已存在于映射中，如果不存在，添加它
				if _, ok := senderMap[strSenderId]; !ok {
					senderCount++
					senderMap[strSenderId] = senderCount
				}

				// 检查目标ID是否已存在于映射中，如果不存在，添加它
				if _, ok := targetMap[strTgrId]; !ok {
					targetCount++
					targetMap[strTgrId] = targetCount
				}

				// 将发送者和目标ID映射为整数，并添加到结果切片
				senderID := senderMap[strSenderId]
				targetID := targetMap[strTgrId]
				result = append(result, []int{senderID, targetID})
			}
		}
	}

	// 创建包含"actions"字段的JSON对象
	jsonData := map[string][][]int{"actions": result}

	// 将JSON对象编码为JSON字符串
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("JSON编码出错:", err)
		return "", err
	}

	// 打印JSON字符串
	return string(jsonBytes), nil
}
