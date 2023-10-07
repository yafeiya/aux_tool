package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
    // "log"
)

type DefineFunction struct {
	Released     string      `json:"released"`
	Dataset_name  string      `json:"dataset_name"`
	Alias        string      `json:"alias"`
	Type         string      `json:"type"`
	Rank         string      `json:"rank"`
	Params       [][]float64 `json:"params"`
	Data_Path     string      `json:"data_path"`
	Function_Body string      `json:"function_body"`
	Lan          string      `json:"lan"`
	Code         string      `json:"code"`
	Description  string      `json:"description"`
	Task         string      `json:"task"`
	Id           int         `json:"id"`
}

type Design struct {
    Released        string                 `json:"username"`
    Dataset_name    string                 `json:"password"`
    Type            string                 `json:"password"`
    Rank            string                 `json:"password"`
    Description     string                 `json:"password"`
    Task            string                 `json:"password"`
    Id              int                    `json:"id"`
    Cell            map[string]interface{} `json:"cell"`
}

type Data struct {
    Designs []Design `json:"design"`
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

    // 创建新的user模型
    newDesign := Design{
        Released:       released,
        Dataset_name:   dataset_name,
        Type:          ttype,
        Rank:           rank,
        Description:    description,
        Task:           task,
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

// 根据ID删除
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

    if userIndex == -1 {
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

// 添加新自定义算子的函数
func AddDefinefunction(released, datasetName, alias, ttype, rank, data_path, function_body, lan, code, description, task string, params [][]float64) error {
    // 读取JSON文件
    jsonData, err := ioutil.ReadFile("./data/data.json")
    if err != nil {
        return err
    }

    var data Data
    if err := json.Unmarshal(jsonData, &data); err != nil {
        return err
    }

    // 创建新的user模型
    newDefineFunction := DefineFunction{
        Released:       released,
        Dataset_name:   dataset_name,
        Alias：         alias
        Type:           ttype,
        Rank:           rank,
        Params:         params,
        Data_Path:      data_path,
        Function_Body   function_body,
        Lan:            lan,    
        Code:           code,
        Description:    description,
        Task:           task,
    }

    // 找到最大的用户ID并自增
    maxID := 0
    for _, defineFunction := range data.DefineFunctions {
        if design.Id > maxID {
            maxID = defineFunction.Id
        }
    }
    newDefineFunction.Id = maxID + 1

    // 插入新的user模型
    data.DefineFunctions = append(data.DefineFunctions, newDefineFunction)

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

// 读取自定义算子列表的函数
func GetDefinefunction(Task, Type string) ([]Design, error) {
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
func UpdateDefinefunction(id int, dataset_name string, rank string, description string) error {
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

// 根据ID删除
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
    for i, design := range data.Designs {
        if design.Id == id {
            Index = i
            break
        }
    }

    if userIndex == -1 {
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