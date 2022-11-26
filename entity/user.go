package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// 登录相关
	Username string `gorm:"type:varchar(30)"`
	Password string `gorm:"type:varchar(200)"`
	// 联系相关
	PhoneNumber string `gorm:"type:varchar(20)"`
	Email       string `gorm:"type:varchar(50)"`
	// 身份相关
	Identity     uint8
	EmployeeName string `gorm:"type:varchar(20)"`
}
