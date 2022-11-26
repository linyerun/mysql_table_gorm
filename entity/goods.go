package entity

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	GoodsName  string `gorm:"type:varchar(50)"`
	GoodsCount uint   // 商品剩余的库存数量
}
