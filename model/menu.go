package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DataMenu struct {
	Name     string     `json:"name"`
	Title    string     `json:"title"`
	Icon     string     `json:"icon"`
	Children []Children `json:"children"`
}

type Children struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type DataFile struct {
	Database       []DataMenu `json:"database"`
	Modelbase      []DataMenu `json:"modelbase"`
	DefineFunction []DataMenu `json:"defineFunction"`
	Design         []DataMenu `json:"design"`
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

func writeDataFile(filename string, data *DataFile) error {
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

func AddMenu(pageKind string, filename string, parentTitle string, childTitle string) string {
	config, err := readDataFile(filename)
	if err != nil {
		return "无法读取配置文件"
	}

	// 查找要添加的页面类型
	var page *[]DataMenu
	switch pageKind {
	case "database":
		page = &config.Database
	case "modelbase":
		page = &config.Modelbase
	case "defineFunction":
		page = &config.DefineFunction
	case "design":
		page = &config.Design
	default:
		return "无效的页面类型"
	}

	// 检查父菜单是否存在
	parentExists := false
	for i, menu := range *page {
		if menu.Title == parentTitle {
			// 父菜单已存在，检查子菜单是否重复
			for _, child := range menu.Children {
				if child.Title == childTitle {
					return "子菜单已存在"
				}
			}
			// 子菜单不重复，添加子菜单
			(*page)[i].Children = append((*page)[i].Children, Children{Name: childTitle, Title: childTitle})
			parentExists = true
			break
		}
	}

	// 如果父菜单不存在，添加新的父菜单和子菜单
	if !parentExists {
		newParent := DataMenu{
			Name:  parentTitle,
			Title: parentTitle,
			Icon:  "ios-navigate",
			Children: []Children{
				{Name: childTitle, Title: childTitle},
			},
		}
		*page = append(*page, newParent)
	}

	// 将更新后的配置写入文件
	if err := writeDataFile(filename, config); err != nil {
		return "fail"
	}
	return "success"

}

func DeleteChildFromDataFile(filename string, pageKind string, parentTitle string, childTitle string) string {
	config, err := readDataFile(filename)
	if err != nil {
		return "无法读取配置文件"
	}

	// 查找要删除的页面类型
	var page *[]DataMenu
	switch pageKind {
	case "database":
		page = &config.Database
	case "modelbase":
		page = &config.Modelbase
	case "defineFunction":
		page = &config.DefineFunction
	case "design":
		page = &config.Design
	default:
		return "无效的页面类型"
	}

	// 查找父菜单的索引
	parentIndex := -1
	for i, menu := range *page {
		if menu.Title == parentTitle {
			parentIndex = i
			break
		}
	}

	// 如果父菜单不存在，返回错误
	if parentIndex == -1 {
		return "父菜单不存在"
	}

	// 如果子菜单为空字符串，删除整个父菜单以及下面的所有子菜单
	if childTitle == "" {
		*page = append((*page)[:parentIndex], (*page)[parentIndex+1:]...)
	} else {
		// 否则，查找子菜单的索引
		childIndex := -1
		for i, child := range (*page)[parentIndex].Children {
			if child.Title == childTitle {
				childIndex = i
				break
			}
		}

		// 如果子菜单不存在，返回错误
		if childIndex == -1 {
			return "子菜单不存在"
		}

		// 删除指定子菜单
		(*page)[parentIndex].Children = append((*page)[parentIndex].Children[:childIndex], (*page)[parentIndex].Children[childIndex+1:]...)
	}

	// 将更新后的配置写入文件
	if err := writeDataFile(filename, config); err != nil {
		return "无法写入配置文件"
	}

	return "success"

}
