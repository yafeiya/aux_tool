package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Canvas struct {
	Name     string    `json:"name"`
	Title    string    `json:"title"`
	Children []CanMenu `json:"children"`
}
type CanMenu struct {
	Name     string     `json:"name"`
	Title    string     `json:"title"`
	Children []Children `json:"children"`
}
type Children struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}
type DataMenu struct {
	Name     string     `json:"name"`
	Title    string     `json:"title"`
	Icon     string     `json:"icon"`
	Children []Children `json:"children"`
}

type DataFile struct {
	Database       []DataMenu `json:"database"`
	Modelbase      []DataMenu `json:"modelbase"`
	DefineFunction []DataMenu `json:"defineFunction"`
	Design         []DataMenu `json:"design"`
	Canvas         []Canvas   `json:"canvas"`
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

func readJson(filename string) (DataFile, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return DataFile{}, err
	}
	var jsonData DataFile
	err = json.Unmarshal(data, &jsonData)
	return jsonData, err

}
func writeJson(filename string, jsonData DataFile) error {
	updatedJsonData, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return err
	}

	// 将更新后的 JSON 数据写回文件
	err = ioutil.WriteFile(filename, updatedJsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

func AddMenu(pageKind string, filename string, parentTitle string, childTitle string, granTitle string) string {
	config, err := readDataFile(filename)
	if err != nil {
		return "无法读取配置文件"
	}
	fmt.Println(pageKind)
	// 查找要添加的页面类型
	if pageKind != "canvas" {
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
	} else {

		// 解析 JSON 数据
		jsonData, err := readJson(filename)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return "fail"
		}
		page := jsonData.Canvas
		// 检查父菜单是否存在
		fmt.Println("// 检查父菜单是否存在")
		for i, menu := range page {
			if menu.Title == parentTitle {
				// 父菜单已存在，检查子菜单是否重复
				fmt.Println("父菜单已存在，检查子菜单是否重复", parentTitle)
				for j, child := range menu.Children {
					if child.Title == childTitle {
						// 子菜单已存在，检查子菜单下的子菜单是否重复
						fmt.Println("子菜单已存在，检查子菜单下的子菜单是否重复")
						for _, gran := range child.Children {
							if gran.Title == granTitle {
								return "所有菜单已存在"
							}
						}
						fmt.Println("孙菜单不存在，添加孙菜单", granTitle)
						newChild := Children{Name: granTitle, Title: granTitle}
						page[i].Children[j].Children = append(page[i].Children[j].Children, newChild)
						writeJson(filename, jsonData)
						return "success"
					}
				}
				// 子菜单不重复，添加子菜单
				fmt.Println("子菜单不重复，添加子菜单", childTitle)
				if granTitle == "" {
					newChild := CanMenu{Name: childTitle, Title: childTitle, Children: nil} // 或者可以用 []Children{}}
					page[i].Children = append(page[i].Children, newChild)
					writeJson(filename, jsonData)

				} else {
					newChild := CanMenu{Name: childTitle, Title: childTitle, Children: []Children{{Name: granTitle, Title: granTitle}}}
					page[i].Children = append(page[i].Children, newChild)
					writeJson(filename, jsonData)
				}
				break
			}

		}

		fmt.Println("如果父菜单不存在，添加新的父菜单和子菜单")
		// 如果父菜单不存在，添加新的父菜单和子菜单
		if childTitle == "" {
			newParent := Canvas{Name: parentTitle, Title: parentTitle, Children: nil}
			jsonData.Canvas = append(jsonData.Canvas, newParent)
			writeJson(filename, jsonData)
			return "success"
		} else if granTitle == "" {
			newParent := Canvas{Name: parentTitle, Title: parentTitle, Children: []CanMenu{{Name: childTitle, Title: childTitle, Children: nil}}}
			jsonData.Canvas = append(jsonData.Canvas, newParent)
			writeJson(filename, jsonData)
			return "success"
		} else {
			newParent := Canvas{Name: parentTitle, Title: parentTitle, Children: []CanMenu{{Name: childTitle, Title: childTitle, Children: []Children{{Name: granTitle, Title: granTitle}}}}}
			jsonData.Canvas = append(jsonData.Canvas, newParent)
			writeJson(filename, jsonData)
			return "success"
		}

	}

}

func DeleteChildFromDataFile(filename string, pageKind string, parentTitle string, childTitle string, granTitle string) string {
	config, err := readDataFile(filename)
	if err != nil {
		return "无法读取配置文件"
	}
	fmt.Println("开始删除目录", pageKind, childTitle, granTitle)
	if pageKind != "canvas" {
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
	} else {
		jsonData, err := readJson(filename)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return "fail"
		}
		fmt.Println("开始删除目录", parentTitle, childTitle, granTitle)
		for i, menu := range jsonData.Canvas {
			if menu.Title == parentTitle {
				if childTitle == "" {
					jsonData.Canvas = append(jsonData.Canvas[:i], jsonData.Canvas[i+1:]...)
					writeJson(filename, jsonData)
				} else if granTitle == "" {
					for j, child := range menu.Children {
						if child.Title == childTitle {
							jsonData.Canvas[i].Children = append(jsonData.Canvas[i].Children[:j], jsonData.Canvas[i].Children[j+1:]...)
							writeJson(filename, jsonData)
						}
					}
				} else {
					for j, child := range menu.Children {
						if child.Title == childTitle {
							for k, gran := range child.Children {
								if gran.Title == granTitle {
									jsonData.Canvas[i].Children[j].Children = append(jsonData.Canvas[i].Children[j].Children[:k], jsonData.Canvas[i].Children[j].Children[k+1:]...)
									writeJson(filename, jsonData)
								}
							}
						}
					}

				}

			}

		}

		return "success"
	}

	return "success"

}
