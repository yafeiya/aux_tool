package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
)

type DefineFunction struct {
	Released      string      `json:"Released"`
	Dataset_name  string      `json:"Dataset_name"`
	Alias         string      `json:"Alias"`
	Type          string      `json:"Type"`
	Rank          string      `json:"Rank"`
	Params        [][]float64 `json:"Params"`
	DataPath      string      `json:"Data_path"`
	Function_body string      `json:"Function_body"`
	Lan           string      `json:"Lan"`
	Code          string      `json:"Code"`
	Description   string      `json:"Description"`
	Task          string      `json:"Task"`
	Id            int         `json:"Id"`
}

type Design struct {
	Released     string                 `json:"Released"`
	Dataset_name string                 `json:"Dataset_name"`
	Type         string                 `json:"Type"`
	Rank         string                 `json:"Rank"`
	Description  string                 `json:"Description"`
	Task         string                 `json:"Task"`
	Id           int                    `json:"Id"`
	Cell         map[string]interface{} `json:"Cell"`
}

type Data struct {
	Designs         []Design         `json:"design"`
	DefineFunctions []DefineFunction `json:"definefunction"`
}

// 添加新设计的函数
func AddDesign(released string, dataset_name string, ttype string, rank string, description string, task string) error {
	// 读取JSON文件
	jsonData, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}

	var data Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return err
	}

	// 检查是否已存在相同dataset_name的定义函数
	for _, existingDesign := range data.Designs {
		if existingDesign.Dataset_name == dataset_name {
			return fmt.Errorf("exist design")
		}
	}

	// 创建新的r模型
	newDesign := Design{
		Released:     released,
		Dataset_name: dataset_name,
		Type:         ttype,
		Rank:         rank,
		Description:  description,
		Task:         task,
	}

	// 找到最大的用户ID并自增
	maxID := 0
	for _, design := range data.Designs {
		if design.Id > maxID {
			maxID = design.Id
		}
	}
	newDesign.Id = maxID + 1

	// 插入新的user模型
	data.Designs = append(data.Designs, newDesign)

	// 更新JSON文件
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("./data/data.json", updatedJSON, 0644); err != nil {
		return err
	}

	return nil
}

// 读取design列表的函数
func GetDesigns(Task, Type string) ([]Design, error) {
	jsonData, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return nil, err
	}

	var data Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	// 如果Task和Type都为空，则返回所有用户
	if Task == "" && Type == "" {
		return data.Designs, nil
	}

	// 否则，筛选出符合条件的用户
	var filteredDesigns []Design
	for _, design := range data.Designs {
		if (Task == "" || design.Task == Task) &&
			(Type == "" || design.Type == Type) {
			filteredDesigns = append(filteredDesigns, design)
		}
	}

	return filteredDesigns, nil
}

// 编辑函数
func UpdateDesign(id int, dataset_name string, rank string, description string) error {
	// 读取JSON文件
	jsonData, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}

	var data Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return err
	}

	// 检查是否已存在相同dataset_name的定义函数
	for _, existingDesign := range data.Designs {
		if existingDesign.Dataset_name == dataset_name {
			return fmt.Errorf("exist design")
		}
	}

	// 遍历列表
	for i := range data.Designs {
		if data.Designs[i].Id == id {
			data.Designs[i].Dataset_name = dataset_name
			data.Designs[i].Rank = rank
			data.Designs[i].Description = description

			// 更新JSON文件
			updatedJSON, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile("./data/data.json", updatedJSON, 0644); err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("not found design")
}

func DeleteDesign(id int) error {
	// 读取JSON文件
	jsonData, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}

	var data Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return err
	}

	// 找到要删除的索引
	Index := -1
	for i, design := range data.Designs {
		if design.Id == id {
			Index = i
			break
		}
	}

	if Index == -1 {
		return nil // 不存在，不执行删除操作
	}

	// 从用户列表中删除用户
	data.Designs = append(data.Designs[:Index], data.Designs[Index+1:]...)

	// 更新JSON文件
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("./data/data.json", updatedJSON, 0644); err != nil {
		return err
	}
	return nil
}

// 添加新的自定义算子
func AddDefinefunction(released string, dataset_name string, alias string, ttype string, rank string, params [][]float64, function_body string, lan string, code string, description string, task string) error {
	// 读取JSON文件
	jsonData, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}

	var data Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return err
	}

	// 检查是否已存在相同dataset_name的定义函数
	for _, existingDefinefunction := range data.DefineFunctions {
		if existingDefinefunction.Dataset_name == dataset_name {
			return fmt.Errorf("exist design")
		}
	}

	// 创建新的r模型
	newDefinefunction := DefineFunction{
		Released:      released,
		Dataset_name:  dataset_name,
		Alias:         alias,
		Type:          ttype,
		Rank:          rank,
		Params:        params,
		Function_body: function_body,
		Lan:           lan,
		Code:          code,
		Description:   description,
		Task:          task,
	}

	// 找到最大的用户ID并自增
	maxID := 0
	for _, definefunction := range data.DefineFunctions {
		if definefunction.Id > maxID {
			maxID = definefunction.Id
		}
	}
	newDefinefunction.Id = maxID + 1

	// 插入新的模型
	data.DefineFunctions = append(data.DefineFunctions, newDefinefunction)

	// 更新JSON文件
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("./data/data.json", updatedJSON, 0644); err != nil {
		return err
	}

	return nil
}

// 读取Definefunction列表的函数
func GetDefinefunctions(Task, Type string) ([]DefineFunction, error) {
	jsonData, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return nil, err
	}

	var data Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	// 如果Task和Type都为空，则返回所有用户
	if Task == "" && Type == "" {
		return data.DefineFunctions, nil
	}

	// 否则，筛选出符合条件的
	var filteredDefineFunctions []DefineFunction
	for _, defineFunction := range data.DefineFunctions {
		if (Task == "" || defineFunction.Task == Task) &&
			(Type == "" || defineFunction.Type == Type) {
			filteredDefineFunctions = append(filteredDefineFunctions, defineFunction)
		}
	}

	return filteredDefineFunctions, nil
}

// 编辑函数
func UpdateDefinefunction(id int, dataset_name string, alias string, rank string, params [][]float64, function_body string, lan string, code string, description string) error {
	// 读取JSON文件
	jsonData, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}

	var data Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return err
	}

	// 遍历列表
	for i := range data.DefineFunctions {
		if data.DefineFunctions[i].Id == id {
			data.DefineFunctions[i].Dataset_name = dataset_name
			data.DefineFunctions[i].Alias = alias
			data.DefineFunctions[i].Rank = rank
			data.DefineFunctions[i].Params = params
			data.DefineFunctions[i].Function_body = function_body
			data.DefineFunctions[i].Lan = lan
			data.DefineFunctions[i].Code = code
			data.DefineFunctions[i].Description = description

			// 更新JSON文件
			updatedJSON, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile("./data/data.json", updatedJSON, 0644); err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("not found Definefunctions")
}

func DeleteDefinefunction(id int) error {
	// 读取JSON文件
	jsonData, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}

	var data Data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return err
	}

	// 找到要删除的索引
	Index := -1
	for i, definefunction := range data.DefineFunctions {
		if definefunction.Id == id {
			Index = i
			break
		}
	}

	if Index == -1 {
		return nil // 不存在，不执行删除操作
	}

	// 从用户列表中删除用户
	data.DefineFunctions = append(data.DefineFunctions[:Index], data.DefineFunctions[Index+1:]...)

	// 更新JSON文件
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("./data/data.json", updatedJSON, 0644); err != nil {
		return err
	}
	return nil
}
