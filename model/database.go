package model

type Database struct {
	Id             uint   `gorm:"primarykey"`
	Released       string `gorm:"column:Released"`
	Dataset_name   string `gorm:"column:Dataset_name"`
	Type           string `gorm:"column:Type"`
	Rank           string `gorm:"column:Rank"`
	Character_type string `gorm:"column:character_type"`
	Header         string `gorm:"column:header"`
	Data_path      string `gorm:"column:data_path"`
	Description    string `gorm:"column:description"`
	Task           string `gorm:"column:task"`
}

func (Database) TableName() string {
	// 返回数据库中表名
	return "database"
}
