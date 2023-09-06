package model

type Example struct {
	Id       uint   `gorm:"primarykey"`
	Example_name string `gorm:"column:Example_name"`
	Rank string `gorm:"column:Rank"`
	State string `gorm:"column:State"`
	Cpu_num uint `gorm:"column:Cpu_num"`
	Gpu_num uint `gorm:"column:Gpu_num"`
	Post_time string `gorm:"column:Post_time"`
	Dataset_url string `gorm:"column:Dataset_url"`
	Model_name string `gorm:"column:Model_name"`
	Model_type string `gorm:"column:Model_type"`
	Epoch_num string `gorm:"column:Epoch_num"`
	Loss string `gorm:"column:Loss"`
	Optimizer string `gorm:"column:Optimizer"`
	Decay string `gorm:"column:Decay"`
	Evalution string `gorm:"column:Evalution"`
	Model_url string `gorm:"column:Model_url"`
	Memory string `gorm:"column:Memory"`
	Start_time string `gorm:"column:Start_time"`
	End_time string `gorm:"column:End_time"`
	

	// "post_date": 1691418125656,
	// "post_time": 1691418125656,
	// "start_time": 1691418125656,
	// "end_time": 1691419507178,
}

func (Example) TableName() string {
	// 返回数据库中表名
	return "example"
}
