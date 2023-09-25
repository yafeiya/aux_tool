package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DefineFunction struct {
	Released     string      `json:"released"`
	DatasetName  string      `json:"dataset_name"`
	Alias        string      `json:"alias"`
	Type         string      `json:"type"`
	Rank         string      `json:"rank"`
	Params       [][]float64 `json:"params"`
	DataPath     string      `json:"data_path"`
	FunctionBody string      `json:"function_body"`
	Lan          string      `json:"lan"`
	Code         string      `json:"code"`
	Description  string      `json:"description"`
	Task         string      `json:"task"`
	ID           int         `json:"id"`
}

type DefineFunctionList struct {
	DefineFunction []DefineFunction `json:"definefunction"`
}

func readDefineFunctionFile(filename string) (*DefineFunctionList, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data DefineFunctionList

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func writeDefineFunctionFile(filename string, data *DefineFunctionList) error {
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, dataJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

func AddDefineFunction(filename string, newFunction DefineFunction) error {
	data, err := readDefineFunctionFile(filename)
	if err != nil {
		return err
	}

	// 添加新的定义函数
	data.DefineFunction = append(data.DefineFunction, newFunction)

	// 写回文件
	err = writeDefineFunctionFile(filename, data)
	if err != nil {
		return err
	}

	return nil
}

func DeleteDefineFunction(filename, datasetName, alias string) error {
	data, err := readDefineFunctionFile(filename)
	if err != nil {
		return err
	}

	// 遍历数据，查找要删除的定义函数
	for i, function := range data.DefineFunction {
		if function.DatasetName == datasetName && function.Alias == alias {
			// 找到要删除的定义函数，从切片中删除它
			data.DefineFunction = append(data.DefineFunction[:i], data.DefineFunction[i+1:]...)

			// 写回文件
			err := writeDefineFunctionFile(filename, data)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("Function not found: datasetName=%s, alias=%s", datasetName, alias)
}

func GetAllDefineFunctions(filename string) ([]DefineFunction, error) {
	data, err := readDefineFunctionFile(filename)
	if err != nil {
		return nil, err
	}

	return data.DefineFunction, nil
}
