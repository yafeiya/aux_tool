package model

type Example struct {
	Id            uint    `gorm:"primarykey"`
	Example_id    int     `gorm:"column:Example_id"`
	Example_name  string  `gorm:"column:Example_name"`
	Rank          string  `gorm:"column:Rank"`
	State         string  `gorm:"column:State"`
	Task          string  `gorm:"column:Task"`
	Type          string  `gorm:"column:Type"`
	Dataset_name  string  `gorm:"column:Dataset_name"`
	Model_name    string  `gorm:"column:Model_name"`
	Train_state   string  `gorm:"column:Train_state"`
	Metics        string  `gorm:"column:Metics"`
	Network_num   int     `gorm:"column:Network_num"`
	Learning_rate float64 `gorm:"column:Learning_rate"`
	Act_function  string  `gorm:"column:Act_function"`
	Radom_seed    int     `gorm:"column:Radom_seed"`
	Optimizer     string  `gorm:"column:Optimizer"`
	Batch_size    int     `gorm:"column:Batch_size"`
	Epoch_num     int     `gorm:"column:Epoch_num"`
	Start_time    string  `gorm:"column:Start_time"`
	End_time      string  `gorm:"column:End_time"`
	Explore_rate  float64 `gorm:"column:Explore_rate"`
	Decay_factor  float64 `gorm:"column:Decay_factor"`
}

func (Example) TableName() string {
	// 返回数据库中表名
	return "example"
}
