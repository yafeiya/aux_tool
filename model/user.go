package model

type User struct {
	Id       uint   `gorm:"primarykey"`
	UserName string `gorm:"column:username"`
	PassWord string `gorm:"column:password"`
}

func (User) TableName() string {
	// 返回数据库中表名
	return "user"
}
