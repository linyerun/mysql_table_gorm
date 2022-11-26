package entity

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	// 联系相关
	PhoneNumber string `gorm:"type:varchar(20)"`
	Email       string `gorm:"type:varchar(50)"`
	ClientName  string `gorm:"type:varchar(20)"`
}
