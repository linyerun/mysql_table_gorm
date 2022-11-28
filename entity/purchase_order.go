package entity

import "gorm.io/gorm"

type PurchaseOrder struct { // 进货单
	gorm.Model
	GoodsId    uint  // 商品ID
	GoodsCount uint  // 商品数量
	State      uint8 // 进货单状态 (0是刚创建, 1是完成进货)
}
