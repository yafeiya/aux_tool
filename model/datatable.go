package model

type Datatable struct {
	Id             		uint   `gorm:"primarykey"`
	Type       			string `gorm:"column:Type"`
	Task   				string `gorm:"column:Task"`
	Dataset_name        string `gorm:"column:Dataset_name"`
	Table_name          string `gorm:"column:Table_name"`
	Header_num 			uint `gorm:"column:Header_num"`
	Data_len         	uint `gorm:"column:Data_len"`
	Data_type      		string `gorm:"column:Data_type"`
	Csv_path			string `gorm:"column:Csv_path"`
}

func (Datatable) TableName() string {
	// 返回数据库中表名
	return "datatable"
}
