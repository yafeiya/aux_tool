package model

type Example struct {
	Id       uint   `gorm:"primarykey"`
	Example_name string `gorm:"column:Example_name"`
	Rank string `gorm:"column:Rank"`
	State string `gorm:"column:State"`
	Cpu_num uint64 `gorm:"column:Cpu_num"`
	Gpu_num uint64 `gorm:"column:Gpu_num"`
	Post_data uint64 `gorm:"column:Post_data"`
	Dataset_url string `gorm:"column:Dataset_url"`
	Model_name string `gorm:"column:Model_name"`
	Model_type string `gorm:"column:Model_type"`
	Epoch_num string `gorm:"column:Epoch_num"`
	Loss string `gorm:"column:Loss"`
	Optimizer string `gorm:"column:Optimizer"`
	Decay string `gorm:"column:Decay"`
	Evaluation string `gorm:"column:Evaluation"`
	Model_url string `gorm:"column:Model_url"`
	Memory string `gorm:"column:Memory"`
	Start_time uint64 `gorm:"column:Start_time"`
	End_time uint64 `gorm:"column:End_time"`

}

func (Example) TableName() string {
	// 返回数据库中表名
	return "example"
}
