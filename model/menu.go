package model

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type DataMenu struct {
	Name            string   	`json:"name"`
	Title       	string 	 	`json:"title"`
	Icon   			string 		`json:"icon"`
	Children        []Children 	`json:"children"`
}

type Children struct {
	Name            string  `json:"name"`
	Title       	string 	`json:"title"`
}

type DataFile struct {
	Database  []DataMenu `json:"database"`
	Modelbase []DataMenu `json:"modelbase"`
	DefineFunction []DataMenu `json:"defineFunction"`
	Design []DataMenu `json:"design"`
}

func readDataFile(filename string) (*DataFile, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("读文件就出错了")
		return nil, err
	}

	var data DataFile

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func AddMenu(pageKind string, filename string, newData DataMenu) error {
	data, err := readDataFile(filename)
	if err != nil {
		return err
	}

	// 根据 pageKind 添加新的数据对象
	switch pageKind {
	case "database":
		data.Database = append(data.Database, newData)
	case "modelbase":
		data.Modelbase = append(data.Modelbase, newData)
	case "defineFunction":
		data.DefineFunction = append(data.DefineFunction, newData)
	case "design":
		data.Design = append(data.Design, newData)
	default:
		return fmt.Errorf("Invalid pageKind: %s", pageKind)
	}


	// 将更新后的数据写回文件
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
