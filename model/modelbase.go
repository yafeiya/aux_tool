package model

type Modelbase struct {
	Id           uint   `gorm:"primarykey"`
	Released     string `gorm:"column:Released"`
	Dataset_name string `gorm:"column:Dataset_name"`
	Type         string `gorm:"column:Type"`
	Rank         string `gorm:"column:Rank"`
	Lan          string `gorm:"column:Character_type"`
	Data_path    string `gorm:"column:Data_path"`
	Description  string `gorm:"column:Description"`
	Code         string `gorm:"column:Code"`
	Task         string `gorm:"column:Task"`
}

func (Modelbase) TableName() string {
	// 返回数据库中表名
	return "modelbase"
}
